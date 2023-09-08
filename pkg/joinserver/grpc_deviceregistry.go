// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/crypto/cryptoservices"
	"go.thethings.network/lorawan-stack/v3/pkg/crypto/cryptoutil"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/warning"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	evtCreateEndDevice = events.Define(
		"js.end_device.create", "create end device",
		events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
		events.WithPropagateToParent(),
	)
	evtUpdateEndDevice = events.Define(
		"js.end_device.update", "update end device",
		events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ),
		events.WithUpdatedFieldsDataType(),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
		events.WithPropagateToParent(),
	)
	evtDeleteEndDevice = events.Define(
		"js.end_device.delete", "delete end device",
		events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
		events.WithPropagateToParent(),
	)
	evtBatchDeleteEndDevices = events.Define(
		"js.end_device.batch.delete", "batch delete end devices",
		events.WithVisibility(ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ),
		events.WithDataType(&ttnpb.EndDeviceIdentifiersList{}),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
		events.WithPropagateToParent(),
	)
)

type jsEndDeviceRegistryServer struct {
	ttnpb.UnimplementedJsEndDeviceRegistryServer

	JS       *JoinServer
	kekLabel string
}

// Get implements ttnpb.JsEndDeviceRegistryServer.
func (srv jsEndDeviceRegistryServer) Get(ctx context.Context, req *ttnpb.GetEndDeviceRequest) (*ttnpb.EndDevice, error) {
	if err := rights.RequireApplication(ctx, req.EndDeviceIds.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ); err != nil {
		return nil, err
	}
	gets := req.FieldMask.GetPaths()
	if ttnpb.HasAnyField(req.FieldMask.GetPaths(),
		"root_keys.app_key.key",
		"root_keys.nwk_key.key",
	) {
		if err := rights.RequireApplication(ctx, req.EndDeviceIds.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ_KEYS); err != nil {
			return nil, err
		}
		gets = ttnpb.AddFields(gets,
			"provisioner_id",
			"provisioning_data",
		)
		if ttnpb.HasAnyField(req.FieldMask.GetPaths(),
			"root_keys.app_key.key",
		) {
			gets = ttnpb.AddFields(gets,
				"root_keys.app_key.encrypted_key",
				"root_keys.app_key.kek_label",
			)
		}
		if ttnpb.HasAnyField(req.FieldMask.GetPaths(),
			"root_keys.nwk_key.key",
		) {
			gets = ttnpb.AddFields(gets,
				"root_keys.nwk_key.encrypted_key",
				"root_keys.nwk_key.kek_label",
			)
		}
	}

	// TODO: Reject writing Claim Authentication Code (https://github.com/TheThingsNetwork/lorawan-stack/issues/5631).
	if ttnpb.HasAnyField(
		req.FieldMask.GetPaths(),
		"claim_authentication_code.value",
		"claim_authentication_code.valid_from",
		"claim_authentication_code.value_to",
	) {
		warning.Add(
			ctx,
			"Storage of claim authentication code in the Join Server registry is deprecated. Use the Identity Server registry instead", //nolint:lll
		)
	}

	logger := log.FromContext(ctx)
	dev, err := srv.JS.devices.GetByID(ctx, req.EndDeviceIds.ApplicationIds, req.EndDeviceIds.DeviceId, gets)
	if errors.IsNotFound(err) {
		return nil, errDeviceNotFound.New()
	}
	if err != nil {
		return nil, err
	}
	if !proto.Equal(dev.Ids.ApplicationIds, req.EndDeviceIds.ApplicationIds) {
		return nil, errDeviceNotFound.New()
	}
	if ttnpb.HasAnyField(req.FieldMask.GetPaths(),
		"root_keys.app_key.key",
		"root_keys.nwk_key.key",
	) {
		rootKeysEnc := dev.RootKeys
		dev.RootKeys = &ttnpb.RootKeys{
			RootKeyId: rootKeysEnc.GetRootKeyId(),
		}
		cc, err := srv.JS.GetPeerConn(ctx, ttnpb.ClusterRole_CRYPTO_SERVER, nil)
		if err != nil {
			if !errors.IsNotFound(err) {
				logger.WithError(err).Debug("Crypto Server connection is not available")
			}
			cc = nil
		}

		if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "root_keys.nwk_key.key") {
			switch {
			case !types.MustAES128Key(rootKeysEnc.GetNwkKey().GetKey()).OrZero().IsZero():
				dev.RootKeys.NwkKey = &ttnpb.KeyEnvelope{
					Key: rootKeysEnc.NwkKey.Key,
				}
			case len(rootKeysEnc.GetNwkKey().GetEncryptedKey()) > 0:
				nwkKey, err := cryptoutil.UnwrapAES128Key(ctx, rootKeysEnc.NwkKey, srv.JS.KeyService())
				if err != nil {
					return nil, err
				}
				dev.RootKeys.NwkKey = &ttnpb.KeyEnvelope{
					Key: nwkKey.Bytes(),
				}
			case cc != nil && dev.ProvisionerId != "":
				nwkKey, err := cryptoservices.NewNetworkRPCClient(cc, srv.JS.KeyService(), srv.JS.WithClusterAuth()).GetNwkKey(ctx, dev)
				if err != nil {
					return nil, err
				}
				if nwkKey != nil {
					dev.RootKeys.NwkKey = &ttnpb.KeyEnvelope{
						Key: nwkKey.Bytes(),
					}
				}
			}
		}

		if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "root_keys.app_key.key") {
			switch {
			case !types.MustAES128Key(rootKeysEnc.GetAppKey().GetKey()).OrZero().IsZero():
				dev.RootKeys.AppKey = &ttnpb.KeyEnvelope{
					Key: rootKeysEnc.GetAppKey().GetKey(),
				}
			case len(rootKeysEnc.GetAppKey().GetEncryptedKey()) > 0:
				appKey, err := cryptoutil.UnwrapAES128Key(ctx, rootKeysEnc.AppKey, srv.JS.KeyService())
				if err != nil {
					return nil, err
				}
				dev.RootKeys.AppKey = &ttnpb.KeyEnvelope{
					Key: appKey.Bytes(),
				}
			case cc != nil && dev.ProvisionerId != "":
				appKey, err := cryptoservices.NewApplicationRPCClient(cc, srv.JS.KeyService(), srv.JS.WithClusterAuth()).GetAppKey(ctx, dev)
				if err != nil {
					return nil, err
				}
				if appKey != nil {
					dev.RootKeys.AppKey = &ttnpb.KeyEnvelope{
						Key: appKey.Bytes(),
					}
				}
			}
		}
	}
	return ttnpb.FilterGetEndDevice(dev, req.FieldMask.GetPaths()...)
}

var (
	errInvalidFieldMask  = errors.DefineInvalidArgument("field_mask", "invalid field mask")
	errInvalidFieldValue = errors.DefineInvalidArgument("field_value", "invalid value of field `{field}`")
)

// Set implements ttnpb.JsEndDeviceRegistryServer.
func (srv jsEndDeviceRegistryServer) Set(ctx context.Context, req *ttnpb.SetEndDeviceRequest) (dev *ttnpb.EndDevice, err error) {
	if req.EndDevice.Ids == nil || req.EndDevice.Ids.JoinEui == nil {
		return nil, errNoJoinEUI.New()
	}
	if types.MustEUI64(req.EndDevice.Ids.DevEui).OrZero().IsZero() {
		return nil, errNoDevEUI.New()
	}

	// TODO: Reject writing Claim Authentication Code (https://github.com/TheThingsNetwork/lorawan-stack/issues/5631).
	if ttnpb.HasAnyField(
		req.FieldMask.GetPaths(),
		"claim_authentication_code.value",
		"claim_authentication_code.valid_from",
		"claim_authentication_code.value_to",
	) {
		warning.Add(
			ctx,
			"Storage of claim authentication code in the Join Server registry is deprecated. Use the Identity Server registry instead", //nolint:lll
		)
	}

	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "root_keys.app_key.key") &&
		types.MustAES128Key(req.EndDevice.GetRootKeys().GetAppKey().GetKey()).OrZero().IsZero() {
		return nil, errInvalidFieldValue.WithAttributes("field", "root_keys.app_key.key")
	}

	if err = rights.RequireApplication(ctx, req.EndDevice.Ids.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE); err != nil {
		return nil, err
	}
	if ttnpb.HasAnyField(req.FieldMask.GetPaths(),
		"root_keys.app_key.key",
		"root_keys.nwk_key.key",
		"root_keys.root_key_id",
	) {
		if err := rights.RequireApplication(ctx, req.EndDevice.Ids.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE_KEYS); err != nil {
			return nil, err
		}
	}

	sets := append(req.FieldMask.GetPaths()[:0:0], req.FieldMask.GetPaths()...)
	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "root_keys.app_key.key") {
		appKey, err := cryptoutil.WrapAES128Key(
			ctx, *types.MustAES128Key(req.EndDevice.RootKeys.AppKey.Key), srv.kekLabel, srv.JS.KeyService(),
		)
		if err != nil {
			return nil, err
		}
		defer func(ke *ttnpb.KeyEnvelope) {
			if dev != nil {
				dev.RootKeys.AppKey = ke
			}
		}(req.EndDevice.RootKeys.AppKey)
		req.EndDevice.RootKeys.AppKey = appKey
		sets = ttnpb.AddFields(sets,
			"root_keys.app_key.encrypted_key",
			"root_keys.app_key.kek_label",
		)
	}
	if ttnpb.HasAnyField(req.FieldMask.GetPaths(), "root_keys.nwk_key.key") {
		if !types.MustAES128Key(req.EndDevice.GetRootKeys().GetNwkKey().GetKey()).OrZero().IsZero() {
			nwkKey, err := cryptoutil.WrapAES128Key(
				ctx, *types.MustAES128Key(req.EndDevice.RootKeys.NwkKey.Key), srv.kekLabel, srv.JS.KeyService(),
			)
			if err != nil {
				return nil, err
			}
			defer func(ke *ttnpb.KeyEnvelope) {
				if dev != nil {
					dev.RootKeys.NwkKey = ke
				}
			}(req.EndDevice.RootKeys.NwkKey)
			req.EndDevice.RootKeys.NwkKey = nwkKey
		} else if req.EndDevice.RootKeys != nil {
			req.EndDevice.RootKeys.NwkKey = nil
		}
		sets = ttnpb.AddFields(sets,
			"root_keys.nwk_key.encrypted_key",
			"root_keys.nwk_key.kek_label",
		)
	}

	var evt events.Event
	dev, err = srv.JS.devices.SetByID(ctx, req.EndDevice.Ids.ApplicationIds, req.EndDevice.Ids.DeviceId, req.FieldMask.GetPaths(), func(dev *ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error) {
		if dev != nil {
			evt = evtUpdateEndDevice.NewWithIdentifiersAndData(ctx, req.EndDevice.Ids, req.FieldMask.GetPaths())
			if err := ttnpb.ProhibitFields(sets,
				"ids.dev_addr",
			); err != nil {
				return nil, nil, errInvalidFieldMask.WithCause(err)
			}
			return req.EndDevice, sets, nil
		}

		evt = evtCreateEndDevice.NewWithIdentifiersAndData(ctx, req.EndDevice.Ids, nil)
		if req.EndDevice.Ids != nil && !types.MustDevAddr(req.EndDevice.Ids.DevAddr).OrZero().IsZero() {
			return nil, nil, errInvalidFieldValue.WithAttributes("field", "ids.dev_addr")
		}
		return req.EndDevice, ttnpb.AddFields(sets,
			"ids.application_ids",
			"ids.dev_eui",
			"ids.device_id",
			"ids.join_eui",
		), nil
	})
	if err != nil {
		return nil, err
	}
	if evt != nil {
		events.Publish(evt)
	}
	return ttnpb.FilterGetEndDevice(dev, req.FieldMask.GetPaths()...)
}

// Provision is deprecated.
// TODO: Remove (https://github.com/TheThingsNetwork/lorawan-stack/issues/999)
func (srv jsEndDeviceRegistryServer) Provision(req *ttnpb.ProvisionEndDevicesRequest, stream ttnpb.JsEndDeviceRegistry_ProvisionServer) error {
	if err := rights.RequireApplication(stream.Context(), req.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE_KEYS); err != nil {
		return err
	}
	return errProvisionerNotFound.WithAttributes("id", req.ProvisionerId)
}

// Delete implements ttnpb.JsEndDeviceRegistryServer.
func (srv jsEndDeviceRegistryServer) Delete(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) (*emptypb.Empty, error) {
	if err := rights.RequireApplication(ctx, ids.ApplicationIds, ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE); err != nil {
		return nil, err
	}
	var (
		evt             events.Event
		joinEUI, devEUI *types.EUI64
	)
	_, err := srv.JS.devices.SetByID(ctx, ids.ApplicationIds, ids.DeviceId, nil, func(dev *ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error) {
		if dev == nil {
			return nil, nil, errDeviceNotFound.New()
		}
		evt = evtDeleteEndDevice.NewWithIdentifiersAndData(ctx, ids, nil)
		joinEUI, devEUI = types.MustEUI64(dev.Ids.JoinEui), types.MustEUI64(dev.Ids.DevEui)
		return nil, nil, nil
	})
	if err != nil {
		return nil, err
	}
	// Try deleting the session keys. This is best effort as it's housekeeping really.
	if joinEUI != nil && devEUI != nil {
		if err := srv.JS.keys.Delete(ctx, *joinEUI, *devEUI); err != nil {
			log.FromContext(ctx).WithError(err).Warn("Failed to delete session keys")
		}
	}
	if evt != nil {
		events.Publish(evt)
	}
	return ttnpb.Empty, err
}

type jsEndDeviceBatchRegistryServer struct {
	ttnpb.UnimplementedJsEndDeviceBatchRegistryServer

	JS *JoinServer
}

// Delete implements ttipb.JsEndDeviceBatchRegistryServer.
func (srv jsEndDeviceBatchRegistryServer) Delete(
	ctx context.Context,
	req *ttnpb.BatchDeleteEndDevicesRequest,
) (*emptypb.Empty, error) {
	// Check if the user has rights on the application.
	if err := rights.RequireApplication(
		ctx,
		req.ApplicationIds,
		ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE,
	); err != nil {
		return nil, err
	}
	deleted, err := srv.JS.devices.BatchDelete(ctx, req.ApplicationIds, req.DeviceIds)
	if err != nil {
		return nil, err
	}
	if len(deleted) != 0 {
		events.Publish(
			evtBatchDeleteEndDevices.NewWithIdentifiersAndData(
				ctx, req.ApplicationIds, &ttnpb.EndDeviceIdentifiersList{
					EndDeviceIds: deleted,
				},
			),
		)
	}

	// Try deleting the session keys in a batch.
	devices := []*ttnpb.EndDeviceIdentifiers{}
	for _, devID := range deleted {
		if devID.DevEui == nil && devID.JoinEui == nil || types.MustEUI64(devID.DevEui).IsZero() {
			continue
		}
		devices = append(devices, devID)
	}
	if err := srv.JS.keys.BatchDelete(ctx, devices); err != nil {
		// We don't return an error since this is an internal cleanup.
		log.FromContext(ctx).WithError(err).Warn("Failed to delete session keys")
	}

	return ttnpb.Empty, nil
}
