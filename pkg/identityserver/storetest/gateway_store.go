// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

package storetest

import (
	"fmt"
	. "testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	is "go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (st *StoreTest) TestGatewayStoreCRUD(t *T) {
	usr1 := st.population.NewUser()
	org1 := st.population.NewOrganization(usr1.GetOrganizationOrUserIdentifiers())

	s, ok := st.PrepareDB(t).(interface {
		Store
		is.GatewayStore
	})
	defer st.DestroyDB(t, true)
	if !ok {
		t.Skip("Store does not implement GatewayStore")
	}
	defer s.Close()

	start := time.Now().Truncate(time.Second)
	mask := fieldMask(ttnpb.GatewayFieldPathsTopLevel...)

	eui := &types.EUI64{1, 2, 3, 4, 5, 6, 7, 8}
	antenna := &ttnpb.GatewayAntenna{
		Gain:      6.0,
		Location:  &ttnpb.Location{Latitude: 12.34, Longitude: 56.78, Altitude: 42, Source: ttnpb.LocationSource_SOURCE_REGISTRY},
		Placement: ttnpb.GatewayAntennaPlacement_OUTDOOR,
	}
	secret := &ttnpb.Secret{
		KeyId: "some-key",
		Value: []byte("some bytes"),
	}
	claim := &ttnpb.GatewayClaimAuthenticationCode{
		ValidFrom: timestamppb.New(start),
		Secret:    secret,
	}
	var created *ttnpb.Gateway

	t.Run("CreateGateway", func(t *T) {
		a, ctx := test.New(t)
		var err error
		start := time.Now().Truncate(time.Second)

		created, err = s.CreateGateway(ctx, &ttnpb.Gateway{
			Ids:                   &ttnpb.GatewayIdentifiers{GatewayId: "foo", Eui: eui.Bytes()},
			Name:                  "Foo Name",
			Description:           "Foo Description",
			Attributes:            attributes,
			AdministrativeContact: usr1.GetOrganizationOrUserIdentifiers(),
			TechnicalContact:      org1.GetOrganizationOrUserIdentifiers(),
			VersionIds: &ttnpb.GatewayVersionIdentifiers{
				BrandId:         "some_brand_id",
				ModelId:         "some_model_id",
				HardwareVersion: "hw_v3",
				FirmwareVersion: "fw_v3",
			},
			GatewayServerAddress:           "localhost",
			AutoUpdate:                     true,
			UpdateChannel:                  "stable",
			FrequencyPlanIds:               []string{"FPLAN_XXX", "FPLAN_YYY"},
			Antennas:                       []*ttnpb.GatewayAntenna{antenna},
			StatusPublic:                   true,
			LocationPublic:                 true,
			ScheduleDownlinkLate:           true,
			EnforceDutyCycle:               true,
			DownlinkPathConstraint:         ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER,
			ScheduleAnytimeDelay:           durationpb.New(time.Second),
			UpdateLocationFromStatus:       true,
			LbsLnsSecret:                   secret,
			ClaimAuthenticationCode:        claim,
			TargetCupsUri:                  "https://cups.example.com",
			TargetCupsKey:                  secret,
			RequireAuthenticatedConnection: true,
			Lrfhss:                         &ttnpb.Gateway_LRFHSS{Supported: true},
			DisablePacketBrokerForwarding:  true,
		})

		if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) {
			a.So(created.GetIds().GetGatewayId(), should.Equal, "foo")
			a.So(created.GetIds().GetEui(), should.Resemble, eui.Bytes())
			a.So(created.Name, should.Equal, "Foo Name")
			a.So(created.Description, should.Equal, "Foo Description")
			a.So(created.Attributes, should.Resemble, attributes)
			a.So(created.AdministrativeContact, should.Resemble, usr1.GetOrganizationOrUserIdentifiers())
			a.So(created.TechnicalContact, should.Resemble, org1.GetOrganizationOrUserIdentifiers())
			a.So(created.VersionIds, should.Resemble, &ttnpb.GatewayVersionIdentifiers{
				BrandId:         "some_brand_id",
				ModelId:         "some_model_id",
				HardwareVersion: "hw_v3",
				FirmwareVersion: "fw_v3",
			})
			a.So(created.GatewayServerAddress, should.Equal, "localhost")
			a.So(created.AutoUpdate, should.BeTrue)
			a.So(created.UpdateChannel, should.Equal, "stable")
			a.So(created.FrequencyPlanIds, should.Resemble, []string{"FPLAN_XXX", "FPLAN_YYY"})
			if a.So(created.Antennas, should.HaveLength, 1) {
				a.So(created.Antennas[0], should.Resemble, antenna)
			}
			a.So(created.StatusPublic, should.BeTrue)
			a.So(created.LocationPublic, should.BeTrue)
			a.So(created.ScheduleDownlinkLate, should.BeTrue)
			a.So(created.EnforceDutyCycle, should.BeTrue)
			a.So(created.DownlinkPathConstraint, should.Equal, ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER)
			a.So(*ttnpb.StdDuration(created.ScheduleAnytimeDelay), should.Equal, time.Second)
			a.So(created.UpdateLocationFromStatus, should.BeTrue)
			a.So(created.LbsLnsSecret, should.Resemble, secret)
			a.So(created.ClaimAuthenticationCode, should.Resemble, claim)
			a.So(created.TargetCupsUri, should.Equal, "https://cups.example.com")
			a.So(created.TargetCupsKey, should.Resemble, secret)
			a.So(created.RequireAuthenticatedConnection, should.BeTrue)
			a.So(created.Lrfhss.Supported, should.BeTrue)
			a.So(created.DisablePacketBrokerForwarding, should.BeTrue)
			a.So(*ttnpb.StdTime(created.CreatedAt), should.HappenWithin, 5*time.Second, start)
			a.So(*ttnpb.StdTime(created.UpdatedAt), should.HappenWithin, 5*time.Second, start)
		}
	})

	t.Run("CreateGateway_AfterCreate", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.CreateGateway(ctx, &ttnpb.Gateway{
			Ids: &ttnpb.GatewayIdentifiers{GatewayId: "foo"},
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsAlreadyExists(err), should.BeTrue)
		}

		_, err = s.CreateGateway(ctx, &ttnpb.Gateway{
			Ids: &ttnpb.GatewayIdentifiers{GatewayId: "other", Eui: eui.Bytes()},
		})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsAlreadyExists(err), should.BeTrue)
		}
	})

	t.Run("GetGateway", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.Resemble, created)
		}
	})

	t.Run("GetGateway_ByEUI", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{Eui: eui.Bytes()}, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.Resemble, created)
		}
	})

	t.Run("GetGateway_Other", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "other"}, mask)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// _, err = s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: ""}, mask)
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("CountGateways", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.CountGateways(ctx)
		if a.So(err, should.BeNil) {
			a.So(got, should.Equal, 1)
		}
	})

	t.Run("FindGateways", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.FindGateways(ctx, nil, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
			a.So(got[0], should.Resemble, created)
		}
	})

	updatedEUI := &types.EUI64{8, 7, 6, 5, 4, 3, 2, 1}
	extraAntenna := &ttnpb.GatewayAntenna{
		Gain:      3.0,
		Location:  &ttnpb.Location{Latitude: 12.34, Longitude: 56.78, Altitude: 30, Source: ttnpb.LocationSource_SOURCE_REGISTRY},
		Placement: ttnpb.GatewayAntennaPlacement_OUTDOOR,
	}
	updatedSecret := &ttnpb.Secret{
		KeyId: "other-key",
		Value: []byte("other bytes"),
	}
	updatedClaim := &ttnpb.GatewayClaimAuthenticationCode{
		ValidFrom: timestamppb.New(start),
		ValidTo:   timestamppb.New(start.Add(time.Hour)),
		Secret:    secret,
	}
	var updated *ttnpb.Gateway

	t.Run("UpdateGateway", func(t *T) {
		a, ctx := test.New(t)
		var err error
		start := time.Now().Truncate(time.Second)

		updated, err = s.UpdateGateway(ctx, &ttnpb.Gateway{
			Ids:                   &ttnpb.GatewayIdentifiers{GatewayId: "foo", Eui: updatedEUI.Bytes()},
			Name:                  "New Foo Name",
			Description:           "New Foo Description",
			Attributes:            updatedAttributes,
			AdministrativeContact: org1.GetOrganizationOrUserIdentifiers(),
			TechnicalContact:      usr1.GetOrganizationOrUserIdentifiers(),
			VersionIds: &ttnpb.GatewayVersionIdentifiers{
				BrandId:         "other_brand_id",
				ModelId:         "other_model_id",
				HardwareVersion: "hw_v3.1",
				FirmwareVersion: "fw_v3.1",
			},
			GatewayServerAddress:           "example.com",
			AutoUpdate:                     false,
			UpdateChannel:                  "develop",
			FrequencyPlanIds:               []string{"FPLAN_XXX"},
			Antennas:                       []*ttnpb.GatewayAntenna{extraAntenna, antenna},
			StatusPublic:                   false,
			LocationPublic:                 false,
			ScheduleDownlinkLate:           false,
			EnforceDutyCycle:               false,
			DownlinkPathConstraint:         ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NONE,
			ScheduleAnytimeDelay:           durationpb.New(time.Second / 2),
			UpdateLocationFromStatus:       false,
			LbsLnsSecret:                   updatedSecret,
			ClaimAuthenticationCode:        updatedClaim,
			TargetCupsUri:                  "https://cups.example.com",
			TargetCupsKey:                  updatedSecret,
			RequireAuthenticatedConnection: false,
			Lrfhss:                         &ttnpb.Gateway_LRFHSS{Supported: false},
			DisablePacketBrokerForwarding:  false,
		}, append(mask, "ids.eui"))
		if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) {
			a.So(updated.GetIds().GetGatewayId(), should.Equal, "foo")
			a.So(updated.GetIds().GetEui(), should.Resemble, updatedEUI.Bytes())
			a.So(updated.Name, should.Equal, "New Foo Name")
			a.So(updated.Description, should.Equal, "New Foo Description")
			a.So(updated.Attributes, should.Resemble, updatedAttributes)
			a.So(updated.AdministrativeContact, should.Resemble, org1.GetOrganizationOrUserIdentifiers())
			a.So(updated.TechnicalContact, should.Resemble, usr1.GetOrganizationOrUserIdentifiers())
			a.So(updated.VersionIds, should.Resemble, &ttnpb.GatewayVersionIdentifiers{
				BrandId:         "other_brand_id",
				ModelId:         "other_model_id",
				HardwareVersion: "hw_v3.1",
				FirmwareVersion: "fw_v3.1",
			})
			a.So(updated.GatewayServerAddress, should.Equal, "example.com")
			a.So(updated.AutoUpdate, should.BeFalse)
			a.So(updated.UpdateChannel, should.Equal, "develop")
			a.So(updated.FrequencyPlanIds, should.Resemble, []string{"FPLAN_XXX"})
			if a.So(updated.Antennas, should.HaveLength, 2) {
				a.So(updated.Antennas[0], should.Resemble, extraAntenna)
				a.So(updated.Antennas[1], should.Resemble, antenna)
			}
			a.So(updated.StatusPublic, should.BeFalse)
			a.So(updated.LocationPublic, should.BeFalse)
			a.So(updated.ScheduleDownlinkLate, should.BeFalse)
			a.So(updated.EnforceDutyCycle, should.BeFalse)
			a.So(updated.DownlinkPathConstraint, should.Equal, ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NONE)
			a.So(*ttnpb.StdDuration(updated.ScheduleAnytimeDelay), should.Equal, time.Second/2)
			a.So(updated.UpdateLocationFromStatus, should.BeFalse)
			a.So(updated.LbsLnsSecret, should.Resemble, updatedSecret)
			a.So(updated.ClaimAuthenticationCode, should.Resemble, updatedClaim)
			a.So(updated.TargetCupsUri, should.Equal, "https://cups.example.com")
			a.So(updated.TargetCupsKey, should.Resemble, updatedSecret)
			a.So(updated.RequireAuthenticatedConnection, should.BeFalse)
			a.So(updated.Lrfhss.GetSupported(), should.BeFalse)
			a.So(updated.DisablePacketBrokerForwarding, should.BeFalse)
			a.So(*ttnpb.StdTime(updated.CreatedAt), should.Equal, *ttnpb.StdTime(created.CreatedAt))
			a.So(*ttnpb.StdTime(updated.UpdatedAt), should.HappenWithin, 5*time.Second, start)
		}
	})

	t.Run("UpdateGateway_Other", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.UpdateGateway(ctx, &ttnpb.Gateway{
			Ids: &ttnpb.GatewayIdentifiers{GatewayId: "other"},
		}, mask)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// _, err = s.UpdateGateway(ctx, &ttnpb.Gateway{
		// 	Ids: &ttnpb.GatewayIdentifiers{GatewayId: ""},
		// }, mask)
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("GetGateway_AfterUpdate", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.Resemble, updated)
		}
	})

	t.Run("DeleteGateway", func(t *T) {
		a, ctx := test.New(t)
		err := s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
		a.So(err, should.BeNil)
	})

	t.Run("DeleteGateway_Other", func(t *T) {
		a, ctx := test.New(t)
		err := s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "other"})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// err = s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: ""})
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("GetGateway_AfterDelete", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
	})

	t.Run("FindGateways_AfterDelete", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.FindGateways(ctx, nil, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.BeEmpty)
		}
	})

	t.Run("GetDeletedGateway", func(t *T) {
		a, ctx := test.New(t)
		ctx = store.WithSoftDeleted(ctx, true)
		got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			if a.So(got.DeletedAt, should.NotBeNil) {
				got.DeletedAt = nil // Unset DeletedAt for the should.Resemble below.
			}
			a.So(got, should.Resemble, updated)
		}
	})

	t.Run("FindDeletedGateways", func(t *T) {
		a, ctx := test.New(t)
		ctx = store.WithSoftDeleted(ctx, true)
		got, err := s.FindGateways(ctx, nil, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
			if a.So(got[0].DeletedAt, should.NotBeNil) {
				got[0].DeletedAt = nil // Unset DeletedAt for the should.Resemble below.
			}
			a.So(got[0], should.Resemble, updated)
		}
	})

	t.Run("RestoreGateway", func(t *T) {
		a, ctx := test.New(t)
		err := s.RestoreGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
		a.So(err, should.BeNil)
	})

	t.Run("RestoreGateway_Other", func(t *T) {
		a, ctx := test.New(t)
		err := s.RestoreGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "other"})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// err = s.RestoreGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: ""})
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("GetGateway_AfterRestore", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.Resemble, updated)
		}
	})

	t.Run("PurgeGateway", func(t *T) {
		a, ctx := test.New(t)
		err := s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
		a.So(err, should.BeNil)
	})

	t.Run("PurgeGateway_Other", func(t *T) {
		a, ctx := test.New(t)
		err := s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "other"})
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// err = s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: ""})
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("CreateAfterPurge", func(t *T) {
		for _, itr := range []int{1, 2} {
			t.Run(fmt.Sprintf("Iteration %d", itr), func(t *T) {
				a, ctx := test.New(t)
				var err error
				_, err = s.CreateGateway(ctx, &ttnpb.Gateway{
					Ids: &ttnpb.GatewayIdentifiers{GatewayId: "foo"},
				})
				a.So(err, should.BeNil)

				err = s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
				a.So(err, should.BeNil)

				err = s.RestoreGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
				a.So(err, should.BeNil)

				got, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
				a.So(err, should.BeNil)
				a.So(got, should.NotBeNil)

				err = s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
				a.So(err, should.BeNil)

				err = s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
				a.So(err, should.BeNil)
			})
		}
	})

	t.Run("Gateway EUI behaviour test set", func(t *T) {
		a, ctx := test.New(t)

		// Creates the first gateway.
		gtwCreated, err := s.CreateGateway(ctx, &ttnpb.Gateway{
			Ids: &ttnpb.GatewayIdentifiers{GatewayId: "foo", Eui: eui.Bytes()},
		})
		a.So(err, should.BeNil)

		// Deletes the gateway.
		err = s.DeleteGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
		a.So(err, should.BeNil)

		deletedGtw, err := s.GetGateway(
			store.WithSoftDeleted(ctx, true), &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask,
		)
		a.So(err, should.BeNil)
		// Validates if the deleted gateway still contains the initial EUI value.
		a.So(deletedGtw.Ids, should.Resemble, gtwCreated.Ids)

		// Creates a new gateway with the same EUI value.
		newGtw, err := s.CreateGateway(ctx, &ttnpb.Gateway{
			Ids: &ttnpb.GatewayIdentifiers{GatewayId: "bar", Eui: eui.Bytes()},
		})
		a.So(err, should.BeNil)
		a.So(newGtw.Ids.Eui, should.Resemble, deletedGtw.Ids.Eui)

		err = s.RestoreGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"})
		a.So(err, should.BeNil)

		// EUI of the restored gateway should be empty.
		oldGtw, err := s.GetGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}, mask)
		a.So(err, should.BeNil)
		a.So(oldGtw.Ids.Eui, should.BeEmpty)

		a.So(s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "foo"}), should.BeNil)
		a.So(s.PurgeGateway(ctx, &ttnpb.GatewayIdentifiers{GatewayId: "bar"}), should.BeNil)
	})
}

func (st *StoreTest) TestGatewayStorePagination(t *T) {
	usr1 := st.population.NewUser()

	var all []*ttnpb.Gateway
	for i := 0; i < 7; i++ {
		all = append(all, st.population.NewGateway(usr1.GetOrganizationOrUserIdentifiers()))
	}

	s, ok := st.PrepareDB(t).(interface {
		Store
		is.GatewayStore
	})
	defer st.DestroyDB(t, false)
	if !ok {
		t.Skip("Store does not implement GatewayStore")
	}
	defer s.Close()

	t.Run("FindGateways_Paginated", func(t *T) {
		a, ctx := test.New(t)

		var total uint64
		for _, page := range []uint32{1, 2, 3, 4} {
			paginateCtx := store.WithPagination(ctx, 2, page, &total)

			got, err := s.FindGateways(paginateCtx, nil, fieldMask(ttnpb.GatewayFieldPathsTopLevel...))
			if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
				if page == 4 {
					a.So(got, should.HaveLength, 1)
				} else {
					a.So(got, should.HaveLength, 2)
				}
				for i, e := range got {
					a.So(e, should.Resemble, all[i+2*int(page-1)])
				}
			}

			a.So(total, should.Equal, 7)
		}
	})
}
