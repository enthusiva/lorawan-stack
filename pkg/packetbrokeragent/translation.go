// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

package packetbrokeragent

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"

	packetbroker "go.packetbroker.org/api/v3"
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/encoding/lorawan"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gopkg.in/square/go-jose.v2"
)

var toPBRegion = map[string]packetbroker.Region{
	band.EU_863_870:      packetbroker.Region_EU_863_870,
	band.US_902_928:      packetbroker.Region_US_902_928,
	band.CN_779_787:      packetbroker.Region_CN_779_787,
	band.EU_433:          packetbroker.Region_EU_433,
	band.AU_915_928:      packetbroker.Region_AU_915_928,
	band.CN_470_510:      packetbroker.Region_CN_470_510,
	band.CN_470_510_20_A: packetbroker.Region_CN_470_510_20_A,
	band.CN_470_510_20_B: packetbroker.Region_CN_470_510_20_B,
	band.CN_470_510_26_A: packetbroker.Region_CN_470_510_26_A,
	band.CN_470_510_26_B: packetbroker.Region_CN_470_510_26_B,
	band.AS_923:          packetbroker.Region_AS_923,
	band.AS_923_2:        packetbroker.Region_AS_923_2,
	band.AS_923_3:        packetbroker.Region_AS_923_3,
	band.AS_923_4:        packetbroker.Region_AS_923_4,
	band.KR_920_923:      packetbroker.Region_KR_920_923,
	band.IN_865_867:      packetbroker.Region_IN_865_867,
	band.RU_864_870:      packetbroker.Region_RU_864_870,
	band.ISM_2400:        packetbroker.Region_WW_2G4,
}

func fromPBDataRate(dataRate *packetbroker.DataRate) (dr *ttnpb.DataRate, ok bool) {
	switch mod := dataRate.GetModulation().(type) {
	case *packetbroker.DataRate_Lora:
		return &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lora{
				Lora: &ttnpb.LoRaDataRate{
					SpreadingFactor: mod.Lora.SpreadingFactor,
					Bandwidth:       mod.Lora.Bandwidth,
					CodingRate:      mod.Lora.CodingRate,
				},
			},
		}, true
	case *packetbroker.DataRate_Fsk:
		return &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Fsk{
				Fsk: &ttnpb.FSKDataRate{
					BitRate: mod.Fsk.BitsPerSecond,
				},
			},
		}, true
	case *packetbroker.DataRate_Lrfhss:
		return &ttnpb.DataRate{
			Modulation: &ttnpb.DataRate_Lrfhss{
				Lrfhss: &ttnpb.LRFHSSDataRate{
					ModulationType:        mod.Lrfhss.ModulationType,
					OperatingChannelWidth: mod.Lrfhss.OperatingChannelWidth,
					CodingRate:            mod.Lrfhss.CodingRate,
				},
			},
		}, true
	default:
		return nil, false
	}
}

func toPBDataRate(dataRate *ttnpb.DataRate) (*packetbroker.DataRate, bool) {
	if dataRate == nil {
		return nil, false
	}
	switch mod := dataRate.GetModulation().(type) {
	case *ttnpb.DataRate_Lora:
		return &packetbroker.DataRate{
			Modulation: &packetbroker.DataRate_Lora{
				Lora: &packetbroker.LoRaDataRate{
					SpreadingFactor: mod.Lora.SpreadingFactor,
					Bandwidth:       mod.Lora.Bandwidth,
					CodingRate:      mod.Lora.CodingRate,
				},
			},
		}, true
	case *ttnpb.DataRate_Fsk:
		return &packetbroker.DataRate{
			Modulation: &packetbroker.DataRate_Fsk{
				Fsk: &packetbroker.FSKDataRate{
					BitsPerSecond: mod.Fsk.BitRate,
				},
			},
		}, true
	case *ttnpb.DataRate_Lrfhss:
		return &packetbroker.DataRate{
			Modulation: &packetbroker.DataRate_Lrfhss{
				Lrfhss: &packetbroker.LRFHSSDataRate{
					ModulationType:        mod.Lrfhss.ModulationType,
					OperatingChannelWidth: mod.Lrfhss.OperatingChannelWidth,
					CodingRate:            mod.Lrfhss.CodingRate,
				},
			},
		}, true
	default:
		return nil, false
	}
}

func fromPBLocation(loc *packetbroker.Location) *ttnpb.Location {
	if loc == nil {
		return nil
	}
	return &ttnpb.Location{
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
		Altitude:  int32(loc.Altitude),
		Accuracy:  int32(loc.Accuracy),
	}
}

func toPBLocation(loc *ttnpb.Location) *packetbroker.Location {
	if loc == nil {
		return nil
	}
	return &packetbroker.Location{
		Longitude: loc.Longitude,
		Latitude:  loc.Latitude,
		Altitude:  float64(loc.Altitude),
		Accuracy:  float32(loc.Accuracy),
	}
}

func toPBTerrestrialAntennaPlacement(p ttnpb.GatewayAntennaPlacement) packetbroker.TerrestrialAntennaPlacement {
	return packetbroker.TerrestrialAntennaPlacement(p)
}

type agentUplinkToken struct {
	ForwarderNetID     types.NetID `json:"fnid"`
	ForwarderTenantID  string      `json:"ftid,omitempty"`
	ForwarderClusterID string      `json:"fcid,omitempty"`
}

type compoundUplinkToken struct {
	Gateway   []byte            `json:"g,omitempty"`
	Forwarder []byte            `json:"f,omitempty"`
	Agent     *agentUplinkToken `json:"a,omitempty"`
}

func wrapUplinkTokens(gateway, forwarder []byte, agent *agentUplinkToken) ([]byte, error) {
	return json.Marshal(compoundUplinkToken{gateway, forwarder, agent})
}

func unwrapUplinkTokens(token []byte) (gateway, forwarder []byte, agent *agentUplinkToken, err error) {
	var t compoundUplinkToken
	if err := json.Unmarshal(token, &t); err != nil {
		return nil, nil, nil, err
	}
	return t.Gateway, t.Forwarder, t.Agent, nil
}

type gatewayUplinkToken struct {
	GatewayUID string `json:"uid"`
	Token      []byte `json:"t"`
}

func wrapGatewayUplinkToken(ctx context.Context, ids *ttnpb.GatewayIdentifiers, ulToken []byte, encrypter jose.Encrypter) ([]byte, error) {
	plaintext, err := json.Marshal(gatewayUplinkToken{
		GatewayUID: unique.ID(ctx, ids),
		Token:      ulToken,
	})
	if err != nil {
		return nil, err
	}
	obj, err := encrypter.Encrypt(plaintext)
	if err != nil {
		return nil, err
	}
	s, err := obj.CompactSerialize()
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func unwrapGatewayUplinkToken(token, key []byte) (string, []byte, error) {
	obj, err := jose.ParseEncrypted(string(token))
	if err != nil {
		return "", nil, err
	}
	plaintext, err := obj.Decrypt(key)
	if err != nil {
		return "", nil, err
	}
	var t gatewayUplinkToken
	if err := json.Unmarshal(plaintext, &t); err != nil {
		return "", nil, err
	}
	return t.GatewayUID, t.Token, nil
}

type gatewayIdentifier interface {
	GetGatewayId() string
	GetEui() []byte
}

func toPBGatewayIdentifier(ids gatewayIdentifier, config ForwarderConfig) *packetbroker.GatewayIdentifier {
	var res *packetbroker.GatewayIdentifier
	if config.IncludeGatewayEUI && ids.GetEui() != nil {
		eui := types.MustEUI64(ids.GetEui())
		res = &packetbroker.GatewayIdentifier{
			Eui: &wrapperspb.UInt64Value{
				Value: eui.MarshalNumber(),
			},
		}
	}
	if config.IncludeGatewayID {
		if res == nil {
			res = &packetbroker.GatewayIdentifier{}
		}
		if config.HashGatewayID {
			hash := sha256.Sum256([]byte(ids.GetGatewayId()))
			res.Id = &packetbroker.GatewayIdentifier_Hash{
				Hash: hash[:],
			}
		} else {
			res.Id = &packetbroker.GatewayIdentifier_Plain{
				Plain: ids.GetGatewayId(),
			}
		}
	}
	return res
}

var (
	errDecodePayload             = errors.DefineInvalidArgument("decode_payload", "decode LoRaWAN payload")
	errUnsupportedLoRaWANVersion = errors.DefineAborted("unsupported_lorawan_version", "unsupported LoRaWAN version `{version}`")
	errUnknownBand               = errors.DefineFailedPrecondition("unknown_band", "unknown band `{band_id}`")
	errUnknownDataRate           = errors.DefineFailedPrecondition("unknown_data_rate", "unknown data rate")
	errUnsupportedMType          = errors.DefineAborted("unsupported_m_type", "unsupported LoRaWAN MType `{m_type}`")
	errWrapGatewayUplinkToken    = errors.DefineAborted("wrap_gateway_uplink_token", "wrap gateway uplink token")
)

func toPBUplink(ctx context.Context, msg *ttnpb.GatewayUplinkMessage, config ForwarderConfig) (*packetbroker.UplinkMessage, error) {
	msg.Message.Payload = &ttnpb.Message{}
	if err := lorawan.UnmarshalMessage(msg.Message.RawPayload, msg.Message.Payload); err != nil {
		return nil, errDecodePayload.WithCause(err)
	}
	if msg.Message.Payload.MHdr.Major != ttnpb.Major_LORAWAN_R1 {
		return nil, errUnsupportedLoRaWANVersion.WithAttributes(
			"version", msg.Message.Payload.MHdr.Major,
		)
	}

	hash := sha256.Sum256(msg.Message.RawPayload[:len(msg.Message.RawPayload)-4]) // The hash is without MIC to detect retransmissions.
	up := &packetbroker.UplinkMessage{
		PhyPayload: &packetbroker.UplinkMessage_PHYPayload{
			Teaser: &packetbroker.PHYPayloadTeaser{
				Hash:   hash[:],
				Length: uint32(len(msg.Message.RawPayload)),
			},
			Value: &packetbroker.UplinkMessage_PHYPayload_Plain{
				Plain: msg.Message.RawPayload,
			},
		},
		Frequency: msg.Message.Settings.Frequency,
	}

	var ok bool
	if up.GatewayRegion, ok = toPBRegion[msg.BandId]; !ok {
		return nil, errUnknownBand.WithAttributes("band_id", msg.BandId)
	}
	if up.DataRate, ok = toPBDataRate(msg.Message.Settings.DataRate); !ok {
		return nil, errUnknownDataRate.New()
	}

	switch pld := msg.Message.Payload.Payload.(type) {
	case *ttnpb.Message_JoinRequestPayload:
		joinEUI := types.MustEUI64(pld.JoinRequestPayload.JoinEui).OrZero()
		devEUI := types.MustEUI64(pld.JoinRequestPayload.DevEui).OrZero()
		devNonce := types.MustDevNonce(pld.JoinRequestPayload.DevNonce).OrZero()
		up.PhyPayload.Teaser.Payload = &packetbroker.PHYPayloadTeaser_JoinRequest{
			JoinRequest: &packetbroker.PHYPayloadTeaser_JoinRequestTeaser{
				JoinEui:  joinEUI.MarshalNumber(),
				DevEui:   devEUI.MarshalNumber(),
				DevNonce: uint32(devNonce.MarshalNumber()),
			},
		}
	case *ttnpb.Message_MacPayload:
		up.PhyPayload.Teaser.Payload = &packetbroker.PHYPayloadTeaser_Mac{
			Mac: &packetbroker.PHYPayloadTeaser_MACPayloadTeaser{
				Confirmed:        pld.MacPayload.FHdr.FCtrl.Ack,
				DevAddr:          types.MustDevAddr(pld.MacPayload.FHdr.DevAddr).OrZero().MarshalNumber(),
				FOpts:            len(pld.MacPayload.FHdr.FOpts) > 0,
				FCnt:             pld.MacPayload.FHdr.FCnt,
				FPort:            pld.MacPayload.FPort,
				FrmPayloadLength: uint32(len(pld.MacPayload.FrmPayload)),
			},
		}
	default:
		return nil, errUnsupportedMType.WithAttributes("m_type", msg.Message.Payload.MHdr.MType)
	}

	var gatewayReceiveTime *time.Time
	var gatewayUplinkToken []byte
	if len(msg.Message.RxMetadata) > 0 && msg.Message.RxMetadata[0].GatewayIds != nil {
		md := msg.Message.RxMetadata[0]
		up.GatewayId = toPBGatewayIdentifier(md.GatewayIds, config)

		var hasGPSTime bool
		var teaser packetbroker.GatewayMetadataTeaser_Terrestrial
		var signalQuality packetbroker.GatewayMetadataSignalQuality_Terrestrial
		var localization *packetbroker.GatewayMetadataLocalization_Terrestrial
		for _, md := range msg.Message.RxMetadata {
			var rssiStandardDeviation *wrapperspb.FloatValue
			if md.RssiStandardDeviation > 0 {
				rssiStandardDeviation = &wrapperspb.FloatValue{
					Value: md.RssiStandardDeviation,
				}
			}

			sqAnt := &packetbroker.GatewayMetadataSignalQuality_Terrestrial_Antenna{
				Index: md.AntennaIndex,
				Value: &packetbroker.TerrestrialGatewayAntennaSignalQuality{
					ChannelRssi:           md.ChannelRssi,
					SignalRssi:            md.SignalRssi,
					RssiStandardDeviation: rssiStandardDeviation,
					Snr:                   md.Snr,
					FrequencyOffset:       md.FrequencyOffset,
				},
			}
			signalQuality.Antennas = append(signalQuality.Antennas, sqAnt)

			if md.Location != nil {
				if localization == nil {
					localization = &packetbroker.GatewayMetadataLocalization_Terrestrial{}
				}
				locAnt := &packetbroker.GatewayMetadataLocalization_Terrestrial_Antenna{
					Index:         md.AntennaIndex,
					Location:      toPBLocation(md.Location),
					SignalQuality: sqAnt.Value,
				}
				if md.FineTimestamp > 0 {
					teaser.FineTimestamp = true
					locAnt.FineTimestamp = &wrapperspb.UInt64Value{
						Value: md.FineTimestamp,
					}
				}
				localization.Antennas = append(localization.Antennas, locAnt)
			}

			earlierGatewayReceiveTime := func(t *timestamppb.Timestamp) bool {
				g := gatewayReceiveTime
				return g == nil || t != nil && ttnpb.StdTime(t).Before(*g)
			}
			switch {
			case hasGPSTime:
			case md.GpsTime != nil:
				gatewayReceiveTime = ttnpb.StdTime(md.GpsTime)
				hasGPSTime = true
			case earlierGatewayReceiveTime(md.Time):
				gatewayReceiveTime = ttnpb.StdTime(md.Time)
			case earlierGatewayReceiveTime(md.ReceivedAt):
				gatewayReceiveTime = ttnpb.StdTime(md.ReceivedAt)
			}

			if md.DownlinkPathConstraint == ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NEVER {
				continue
			}

			if len(md.UplinkToken) == 0 {
				log.FromContext(ctx).WithField("downlink_path_constraint", md.DownlinkPathConstraint).Error("Empty uplink token with favorable downlink path constraint")
				continue
			}

			if len(gatewayUplinkToken) == 0 {
				var err error
				gatewayUplinkToken, err = wrapGatewayUplinkToken(ctx, md.GatewayIds, md.UplinkToken, config.TokenEncrypter)
				if err != nil {
					return nil, errWrapGatewayUplinkToken.WithCause(err)
				}
			}
		}

		up.GatewayMetadata = &packetbroker.UplinkMessage_GatewayMetadata{
			Teaser: &packetbroker.GatewayMetadataTeaser{
				Value: &packetbroker.GatewayMetadataTeaser_Terrestrial_{
					Terrestrial: &teaser,
				},
			},
			SignalQuality: &packetbroker.UplinkMessage_GatewayMetadata_PlainSignalQuality{
				PlainSignalQuality: &packetbroker.GatewayMetadataSignalQuality{
					Value: &packetbroker.GatewayMetadataSignalQuality_Terrestrial_{
						Terrestrial: &signalQuality,
					},
				},
			},
		}
		if localization != nil {
			up.GatewayMetadata.Localization = &packetbroker.UplinkMessage_GatewayMetadata_PlainLocalization{
				PlainLocalization: &packetbroker.GatewayMetadataLocalization{
					Value: &packetbroker.GatewayMetadataLocalization_Terrestrial_{
						Terrestrial: localization,
					},
				},
			}
		}
	}

	up.ForwarderReceiveTime = msg.Message.ReceivedAt
	up.GatewayReceiveTime = ttnpb.ProtoTime(gatewayReceiveTime)
	up.GatewayUplinkToken = gatewayUplinkToken

	return up, nil
}

var errWrapUplinkTokens = errors.DefineAborted("wrap_uplink_tokens", "wrap uplink tokens")

func fromPBUplink(ctx context.Context, msg *packetbroker.RoutedUplinkMessage, receivedAt time.Time, includeHops bool) (*ttnpb.UplinkMessage, error) {
	dataRate, ok := fromPBDataRate(msg.Message.DataRate)
	if !ok {
		return nil, errUnknownDataRate.New()
	}

	var forwarderNetID, homeNetworkNetID types.NetID
	if err := forwarderNetID.UnmarshalNumber(msg.ForwarderNetId); err != nil {
		return nil, errNetID.WithCause(err).WithAttributes("net_id", msg.ForwarderNetId)
	}
	if err := homeNetworkNetID.UnmarshalNumber(msg.HomeNetworkNetId); err != nil {
		return nil, errNetID.WithCause(err).WithAttributes("net_id", msg.HomeNetworkNetId)
	}
	var (
		downlinkPathConstraint = ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NEVER
		uplinkToken            []byte
	)
	if len(msg.Message.GatewayUplinkToken) > 0 || len(msg.Message.ForwarderUplinkToken) > 0 {
		downlinkPathConstraint = ttnpb.DownlinkPathConstraint_DOWNLINK_PATH_CONSTRAINT_NONE
		token := &agentUplinkToken{
			ForwarderNetID:     forwarderNetID,
			ForwarderTenantID:  msg.ForwarderTenantId,
			ForwarderClusterID: msg.ForwarderClusterId,
		}
		var err error
		uplinkToken, err = wrapUplinkTokens(msg.Message.GatewayUplinkToken, msg.Message.ForwarderUplinkToken, token)
		if err != nil {
			return nil, errWrapUplinkTokens.WithCause(err)
		}
	}

	up := &ttnpb.UplinkMessage{
		RawPayload: msg.Message.PhyPayload.GetPlain(),
		Settings: &ttnpb.TxSettings{
			DataRate:  dataRate,
			Frequency: msg.Message.Frequency,
		},
		ReceivedAt:     timestamppb.New(receivedAt),
		CorrelationIds: events.CorrelationIDsFromContext(ctx),
	}

	receiveTime := msg.Message.GatewayReceiveTime
	if gtwMd := msg.Message.GatewayMetadata; gtwMd != nil {
		pbMD := &ttnpb.PacketBrokerMetadata{
			MessageId:            msg.Id,
			ForwarderNetId:       forwarderNetID.Bytes(),
			ForwarderTenantId:    msg.ForwarderTenantId,
			ForwarderClusterId:   msg.ForwarderClusterId,
			HomeNetworkNetId:     homeNetworkNetID.Bytes(),
			HomeNetworkTenantId:  msg.HomeNetworkTenantId,
			HomeNetworkClusterId: msg.HomeNetworkClusterId,
		}
		if id := msg.GetMessage().GetGatewayId(); id != nil {
			if eui := id.Eui; eui != nil {
				euiVal := types.EUI64{}
				euiVal.UnmarshalNumber(eui.Value)
				pbMD.ForwarderGatewayEui = euiVal.Bytes()
			}
			switch s := id.Id.(type) {
			case *packetbroker.GatewayIdentifier_Hash:
				pbMD.ForwarderGatewayId = &wrapperspb.StringValue{
					Value: base64.StdEncoding.EncodeToString(s.Hash),
				}
			case *packetbroker.GatewayIdentifier_Plain:
				pbMD.ForwarderGatewayId = &wrapperspb.StringValue{
					Value: s.Plain,
				}
			}
		}
		if includeHops {
			pbMD.Hops = make([]*ttnpb.PacketBrokerRouteHop, 0, len(msg.Hops))
			for _, h := range msg.Hops {
				pbMD.Hops = append(pbMD.Hops, &ttnpb.PacketBrokerRouteHop{
					ReceivedAt:    h.ReceivedAt,
					SenderName:    h.SenderName,
					SenderAddress: h.SenderAddress,
					ReceiverName:  h.ReceiverName,
					ReceiverAgent: h.ReceiverAgent,
				})
			}
		}
		if md := gtwMd.GetPlainLocalization().GetTerrestrial(); md != nil {
			for _, ant := range md.Antennas {
				up.RxMetadata = append(up.RxMetadata, &ttnpb.RxMetadata{
					GatewayIds:             cluster.PacketBrokerGatewayID,
					PacketBroker:           pbMD,
					AntennaIndex:           ant.Index,
					Time:                   receiveTime,
					FineTimestamp:          ant.FineTimestamp.GetValue(),
					Rssi:                   ant.SignalQuality.GetChannelRssi(),
					ChannelRssi:            ant.SignalQuality.GetChannelRssi(),
					SignalRssi:             ant.SignalQuality.GetSignalRssi(),
					RssiStandardDeviation:  ant.SignalQuality.GetRssiStandardDeviation().GetValue(),
					Snr:                    ant.SignalQuality.GetSnr(),
					FrequencyOffset:        ant.SignalQuality.GetFrequencyOffset(),
					Location:               fromPBLocation(ant.Location),
					DownlinkPathConstraint: downlinkPathConstraint,
					UplinkToken:            uplinkToken,
					ReceivedAt:             msg.Message.ForwarderReceiveTime,
				})
			}
		}
		if md := gtwMd.GetPlainSignalQuality().GetTerrestrial(); md != nil {
			for _, ant := range md.Antennas {
				var md *ttnpb.RxMetadata
				for _, locMd := range up.RxMetadata {
					if locMd.AntennaIndex == ant.Index {
						md = locMd
						break
					}
				}
				if md == nil {
					md = &ttnpb.RxMetadata{
						GatewayIds:             cluster.PacketBrokerGatewayID,
						PacketBroker:           pbMD,
						AntennaIndex:           ant.Index,
						Time:                   receiveTime,
						DownlinkPathConstraint: downlinkPathConstraint,
						UplinkToken:            uplinkToken,
						ReceivedAt:             msg.Message.ForwarderReceiveTime,
					}
					up.RxMetadata = append(up.RxMetadata, md)
				}
				md.Rssi = ant.Value.GetChannelRssi()
				md.ChannelRssi = ant.Value.GetChannelRssi()
				md.SignalRssi = ant.Value.GetSignalRssi()
				md.RssiStandardDeviation = ant.Value.GetRssiStandardDeviation().GetValue()
				md.Snr = ant.Value.GetSnr()
				md.FrequencyOffset = ant.Value.GetFrequencyOffset()
			}
		}
	}

	return up, nil
}

var (
	fromPBClass = map[packetbroker.DownlinkMessageClass]ttnpb.Class{
		packetbroker.DownlinkMessageClass_CLASS_A: ttnpb.Class_CLASS_A,
		packetbroker.DownlinkMessageClass_CLASS_B: ttnpb.Class_CLASS_B,
		packetbroker.DownlinkMessageClass_CLASS_C: ttnpb.Class_CLASS_C,
	}
	toPBClass = map[ttnpb.Class]packetbroker.DownlinkMessageClass{
		ttnpb.Class_CLASS_A: packetbroker.DownlinkMessageClass_CLASS_A,
		ttnpb.Class_CLASS_B: packetbroker.DownlinkMessageClass_CLASS_B,
		ttnpb.Class_CLASS_C: packetbroker.DownlinkMessageClass_CLASS_C,
	}
	fromPBPriority = map[packetbroker.DownlinkMessagePriority]ttnpb.TxSchedulePriority{
		packetbroker.DownlinkMessagePriority_LOWEST:  ttnpb.TxSchedulePriority_LOWEST,
		packetbroker.DownlinkMessagePriority_LOW:     ttnpb.TxSchedulePriority_LOW,
		packetbroker.DownlinkMessagePriority_NORMAL:  ttnpb.TxSchedulePriority_NORMAL,
		packetbroker.DownlinkMessagePriority_HIGH:    ttnpb.TxSchedulePriority_HIGH,
		packetbroker.DownlinkMessagePriority_HIGHEST: ttnpb.TxSchedulePriority_HIGHEST,
	}
	toPBPriority = map[ttnpb.TxSchedulePriority]packetbroker.DownlinkMessagePriority{
		ttnpb.TxSchedulePriority_LOWEST:       packetbroker.DownlinkMessagePriority_LOWEST,
		ttnpb.TxSchedulePriority_LOW:          packetbroker.DownlinkMessagePriority_LOW,
		ttnpb.TxSchedulePriority_BELOW_NORMAL: packetbroker.DownlinkMessagePriority_LOW,
		ttnpb.TxSchedulePriority_NORMAL:       packetbroker.DownlinkMessagePriority_NORMAL,
		ttnpb.TxSchedulePriority_ABOVE_NORMAL: packetbroker.DownlinkMessagePriority_HIGH,
		ttnpb.TxSchedulePriority_HIGH:         packetbroker.DownlinkMessagePriority_HIGH,
		ttnpb.TxSchedulePriority_HIGHEST:      packetbroker.DownlinkMessagePriority_HIGHEST,
	}
)

var (
	errNoRequest                  = errors.DefineFailedPrecondition("no_request", "downlink message is not a transmission request")
	errUnknownClass               = errors.DefineInvalidArgument("unknown_class", "unknown class `{class}`")
	errUnknownPriority            = errors.DefineInvalidArgument("unknown_priority", "unknown priority `{priority}`")
	errNoDownlinkPaths            = errors.DefineFailedPrecondition("no_downlink_paths", "no downlink paths")
	errInvalidDownlinkPath        = errors.DefineFailedPrecondition("downlink_path", "invalid uplink token downlink path")
	errFrequencyPlanNotConfigured = errors.DefineInvalidArgument("frequency_plan_not_configured", "frequency plan `{id}` is not configured")
	errIncompatibleDataRate       = errors.DefineInvalidArgument("incompatible_data_rate", "incompatible data rate in Rx{rx_window}")
)

func toPBDownlink(ctx context.Context, msg *ttnpb.DownlinkMessage, fps frequencyPlansStore) (*packetbroker.DownlinkMessage, *agentUplinkToken, error) {
	req := msg.GetRequest()
	if req == nil {
		return nil, nil, errNoRequest.New()
	}

	fp, err := fps.GetByID(req.FrequencyPlanId)
	if err != nil {
		return nil, nil, errFrequencyPlanNotConfigured.WithAttributes("id", req.FrequencyPlanId)
	}

	down := &packetbroker.DownlinkMessage{
		PhyPayload: msg.RawPayload,
		Rx1Delay:   durationpb.New(req.Rx1Delay.Duration()),
	}
	var ok bool
	if down.Region, ok = toPBRegion[fp.BandID]; !ok {
		return nil, nil, errUnknownBand.WithAttributes("band_id", fp.BandID)
	}
	for i, rx := range []struct {
		dataRate  *ttnpb.DataRate
		frequency uint64
		dst       **packetbroker.DownlinkMessage_RXSettings
	}{
		{req.Rx1DataRate, req.Rx1Frequency, &down.Rx1},
		{req.Rx2DataRate, req.Rx2Frequency, &down.Rx2},
	} {
		if rx.frequency == 0 || rx.dataRate == nil {
			continue
		}
		pbDR, ok := toPBDataRate(rx.dataRate)
		if !ok {
			return nil, nil, errIncompatibleDataRate.WithAttributes("rx_window", i+1)
		}
		*rx.dst = &packetbroker.DownlinkMessage_RXSettings{
			DataRate:  pbDR,
			Frequency: rx.frequency,
		}
	}
	if down.Class, ok = toPBClass[req.Class]; !ok {
		return nil, nil, errUnknownClass.WithAttributes("class", req.Class)
	}
	if down.Priority, ok = toPBPriority[req.Priority]; !ok {
		return nil, nil, errUnknownPriority.WithAttributes("priority", req.Priority)
	}
	if len(req.DownlinkPaths) == 0 {
		return nil, nil, errNoDownlinkPaths.New()
	}
	uplinkToken := req.DownlinkPaths[0].GetUplinkToken()
	if len(uplinkToken) == 0 {
		return nil, nil, errInvalidDownlinkPath.New()
	}
	var token *agentUplinkToken
	down.GatewayUplinkToken, down.ForwarderUplinkToken, token, err = unwrapUplinkTokens(uplinkToken)
	if err != nil {
		return nil, nil, errInvalidDownlinkPath.WithCause(err)
	}

	return down, token, nil
}

var (
	errUnwrapGatewayUplinkToken = errors.DefineAborted("unwrap_gateway_uplink_token", "unwrap gateway uplink token")
	errInvalidRx1Delay          = errors.DefineInvalidArgument("invalid_rx1_delay", "invalid Rx1 delay")
)

func fromPBDownlink(ctx context.Context, msg *packetbroker.DownlinkMessage, receivedAt time.Time, conf ForwarderConfig) (uid string, res *ttnpb.DownlinkMessage, err error) {
	uid, token, err := unwrapGatewayUplinkToken(msg.GatewayUplinkToken, conf.TokenKey)
	if err != nil {
		return "", nil, errUnwrapGatewayUplinkToken.WithCause(err)
	}

	req := &ttnpb.TxRequest{
		DownlinkPaths: []*ttnpb.DownlinkPath{
			{
				Path: &ttnpb.DownlinkPath_UplinkToken{
					UplinkToken: token,
				},
			},
		},
	}
	// NOTE: The Things Stack expects the frequency plan ID; not the band ID. Since the frequency plan ID cannot be
	// inferred from the downlink message from Packet Broker, it is intentionally left blank. This makes the Gateway
	// Server fallback to a single frequency plan configured for the gateway. This does not work if there are multiple
	// frequency plans. (https://github.com/TheThingsNetwork/lorawan-stack/issues/1394)

	var ok bool
	if req.Class, ok = fromPBClass[msg.Class]; !ok {
		return "", nil, errUnknownClass.WithAttributes("class", msg.Class)
	}
	if req.Priority, ok = fromPBPriority[msg.Priority]; !ok {
		return "", nil, errUnknownPriority.WithAttributes("priority", msg.Priority)
	}
	rx1Delay := msg.Rx1Delay.AsDuration()
	req.Rx1Delay = ttnpb.RxDelay(rx1Delay / time.Second)
	for i, rx := range []struct {
		settings  *packetbroker.DownlinkMessage_RXSettings
		dataRate  **ttnpb.DataRate
		frequency *uint64
	}{
		{msg.Rx1, &req.Rx1DataRate, &req.Rx1Frequency},
		{msg.Rx2, &req.Rx2DataRate, &req.Rx2Frequency},
	} {
		if rx.settings == nil {
			continue
		}
		dr, ok := fromPBDataRate(rx.settings.DataRate)
		if !ok {
			return "", nil, errIncompatibleDataRate.WithAttributes("rx_window", i+1)
		}
		*rx.dataRate = dr
		*rx.frequency = rx.settings.Frequency
	}

	down := &ttnpb.DownlinkMessage{
		RawPayload:     msg.PhyPayload,
		CorrelationIds: events.CorrelationIDsFromContext(ctx),
		Settings: &ttnpb.DownlinkMessage_Request{
			Request: req,
		},
	}
	return uid, down, nil
}

func fromPBDevAddrBlocks(blocks []*packetbroker.DevAddrBlock) []*ttnpb.PacketBrokerDevAddrBlock {
	res := make([]*ttnpb.PacketBrokerDevAddrBlock, len(blocks))
	for i, b := range blocks {
		res[i] = &ttnpb.PacketBrokerDevAddrBlock{
			DevAddrPrefix: &ttnpb.DevAddrPrefix{
				Length: b.GetPrefix().GetLength(),
			},
			HomeNetworkClusterId: b.GetHomeNetworkClusterId(),
		}
		var devAddr types.DevAddr
		devAddr.UnmarshalNumber(b.GetPrefix().GetValue())
		res[i].DevAddrPrefix.DevAddr = devAddr.Bytes()
	}
	return res
}

func toPBDevAddrBlocks(blocks []*ttnpb.PacketBrokerDevAddrBlock) []*packetbroker.DevAddrBlock {
	res := make([]*packetbroker.DevAddrBlock, len(blocks))
	for i, b := range blocks {
		devAddr := types.MustDevAddr(b.GetDevAddrPrefix().DevAddr)
		res[i] = &packetbroker.DevAddrBlock{
			Prefix: &packetbroker.DevAddrPrefix{
				Value:  devAddr.MarshalNumber(),
				Length: b.GetDevAddrPrefix().GetLength(),
			},
			HomeNetworkClusterId: b.GetHomeNetworkClusterId(),
		}
	}
	return res
}

func fromPBContactInfo(admin, technical *packetbroker.ContactInfo) []*ttnpb.ContactInfo {
	res := make([]*ttnpb.ContactInfo, 0, 2)
	if email := admin.GetEmail(); email != "" {
		res = append(res, &ttnpb.ContactInfo{
			ContactType:   ttnpb.ContactType_CONTACT_TYPE_OTHER,
			ContactMethod: ttnpb.ContactMethod_CONTACT_METHOD_EMAIL,
			Value:         email,
		})
	}
	if email := technical.GetEmail(); email != "" {
		res = append(res, &ttnpb.ContactInfo{
			ContactType:   ttnpb.ContactType_CONTACT_TYPE_TECHNICAL,
			ContactMethod: ttnpb.ContactMethod_CONTACT_METHOD_EMAIL,
			Value:         email,
		})
	}
	return res
}

func toPBContactInfo(info []*ttnpb.ContactInfo) (admin, technical *packetbroker.ContactInfo) {
	for _, c := range info {
		if c.GetContactMethod() != ttnpb.ContactMethod_CONTACT_METHOD_EMAIL || c.GetValue() == "" {
			continue
		}
		switch c.GetContactType() {
		case ttnpb.ContactType_CONTACT_TYPE_OTHER:
			admin = &packetbroker.ContactInfo{
				Email: c.GetValue(),
			}
		case ttnpb.ContactType_CONTACT_TYPE_TECHNICAL:
			technical = &packetbroker.ContactInfo{
				Email: c.GetValue(),
			}
		}
	}
	return admin, technical
}

func fromPBUplinkRoutingPolicy(policy *packetbroker.RoutingPolicy_Uplink) *ttnpb.PacketBrokerRoutingPolicyUplink {
	return &ttnpb.PacketBrokerRoutingPolicyUplink{
		JoinRequest:     policy.GetJoinRequest(),
		MacData:         policy.GetMacData(),
		ApplicationData: policy.GetApplicationData(),
		SignalQuality:   policy.GetSignalQuality(),
		Localization:    policy.GetLocalization(),
	}
}

func fromPBDownlinkRoutingPolicy(policy *packetbroker.RoutingPolicy_Downlink) *ttnpb.PacketBrokerRoutingPolicyDownlink {
	return &ttnpb.PacketBrokerRoutingPolicyDownlink{
		JoinAccept:      policy.GetJoinAccept(),
		MacData:         policy.GetMacData(),
		ApplicationData: policy.GetApplicationData(),
	}
}

func fromPBDefaultRoutingPolicy(policy *packetbroker.RoutingPolicy) *ttnpb.PacketBrokerDefaultRoutingPolicy {
	return &ttnpb.PacketBrokerDefaultRoutingPolicy{
		UpdatedAt: policy.GetUpdatedAt(),
		Uplink:    fromPBUplinkRoutingPolicy(policy.GetUplink()),
		Downlink:  fromPBDownlinkRoutingPolicy(policy.GetDownlink()),
	}
}

func fromPBRoutingPolicy(policy *packetbroker.RoutingPolicy) *ttnpb.PacketBrokerRoutingPolicy {
	var homeNetworkID *ttnpb.PacketBrokerNetworkIdentifier
	if policy.HomeNetworkNetId != 0 || policy.HomeNetworkTenantId != "" {
		homeNetworkID = &ttnpb.PacketBrokerNetworkIdentifier{
			NetId:    policy.GetHomeNetworkNetId(),
			TenantId: policy.GetHomeNetworkTenantId(),
		}
	}
	return &ttnpb.PacketBrokerRoutingPolicy{
		ForwarderId: &ttnpb.PacketBrokerNetworkIdentifier{
			NetId:    policy.GetForwarderNetId(),
			TenantId: policy.GetForwarderTenantId(),
		},
		HomeNetworkId: homeNetworkID,
		UpdatedAt:     policy.GetUpdatedAt(),
		Uplink:        fromPBUplinkRoutingPolicy(policy.GetUplink()),
		Downlink:      fromPBDownlinkRoutingPolicy(policy.GetDownlink()),
	}
}

func toPBUplinkRoutingPolicy(policy *ttnpb.PacketBrokerRoutingPolicyUplink) *packetbroker.RoutingPolicy_Uplink {
	return &packetbroker.RoutingPolicy_Uplink{
		JoinRequest:     policy.GetJoinRequest(),
		MacData:         policy.GetMacData(),
		ApplicationData: policy.GetApplicationData(),
		SignalQuality:   policy.GetSignalQuality(),
		Localization:    policy.GetLocalization(),
	}
}

func toPBDownlinkRoutingPolicy(policy *ttnpb.PacketBrokerRoutingPolicyDownlink) *packetbroker.RoutingPolicy_Downlink {
	return &packetbroker.RoutingPolicy_Downlink{
		JoinAccept:      policy.GetJoinAccept(),
		MacData:         policy.GetMacData(),
		ApplicationData: policy.GetApplicationData(),
	}
}

func fromPBDefaultGatewayVisibility(visibility *packetbroker.GatewayVisibility) *ttnpb.PacketBrokerDefaultGatewayVisibility {
	return &ttnpb.PacketBrokerDefaultGatewayVisibility{
		UpdatedAt: visibility.GetUpdatedAt(),
		Visibility: &ttnpb.PacketBrokerGatewayVisibility{
			Location:         visibility.GetLocation(),
			AntennaPlacement: visibility.GetAntennaPlacement(),
			AntennaCount:     visibility.GetAntennaCount(),
			FineTimestamps:   visibility.GetFineTimestamps(),
			ContactInfo:      visibility.GetContactInfo(),
			Status:           visibility.GetStatus(),
			FrequencyPlan:    visibility.GetFrequencyPlan(),
			PacketRates:      visibility.GetPacketRates(),
		},
	}
}

var errInconsistentBands = errors.DefineInvalidArgument("inconsistent_bands", "inconsistent bands")

func toPBFrequencyPlan(fps ...*frequencyplans.FrequencyPlan) (*packetbroker.GatewayFrequencyPlan, error) {
	if len(fps) == 0 {
		return nil, nil
	}
	phy, err := band.GetLatest(fps[0].BandID)
	if err != nil {
		return nil, err
	}
	res := &packetbroker.GatewayFrequencyPlan{
		Region: toPBRegion[phy.ID],
	}

	type singleSFChannel struct {
		frequency uint64
		sf, bw    uint32
	}
	singleSFChs := make(map[singleSFChannel]struct{})
	multiSFChs := make(map[uint64]struct{})

	for _, fp := range fps {
		if fp.BandID != phy.ID {
			return nil, errInconsistentBands.New()
		}
		for _, ch := range fp.UplinkChannels {
			if idx := ch.MinDataRate; idx == ch.MaxDataRate {
				dr, ok := phy.DataRates[ttnpb.DataRateIndex(idx)]
				if !ok {
					continue
				}
				switch mod := dr.Rate.Modulation.(type) {
				case *ttnpb.DataRate_Fsk:
					res.FskChannel = &packetbroker.GatewayFrequencyPlan_FSKChannel{
						Frequency: ch.Frequency,
					}
				case *ttnpb.DataRate_Lora:
					chKey := singleSFChannel{ch.Frequency, mod.Lora.SpreadingFactor, mod.Lora.Bandwidth}
					if _, ok := singleSFChs[chKey]; ok {
						continue
					}
					res.LoraSingleSfChannels = append(res.LoraSingleSfChannels, &packetbroker.GatewayFrequencyPlan_LoRaSingleSFChannel{
						Frequency:       ch.Frequency,
						SpreadingFactor: mod.Lora.SpreadingFactor,
						Bandwidth:       mod.Lora.Bandwidth,
					})
					singleSFChs[chKey] = struct{}{}
				}
			} else {
				if _, ok := multiSFChs[ch.Frequency]; ok {
					continue
				}
				res.LoraMultiSfChannels = append(res.LoraMultiSfChannels, &packetbroker.GatewayFrequencyPlan_LoRaMultiSFChannel{
					Frequency: ch.Frequency,
				})
				multiSFChs[ch.Frequency] = struct{}{}
			}
		}
	}
	return res, nil
}
