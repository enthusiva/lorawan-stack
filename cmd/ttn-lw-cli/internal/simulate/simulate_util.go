// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

// Package simulate implements the simulation off device communication.
package simulate

import (
	"go.thethings.network/lorawan-stack/v3/pkg/band"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	errDataRate             = errors.DefineInvalidArgument("data_rate", "data rate is invalid")
	errFrequency            = errors.DefineInvalidArgument("frequency", "frequency is invalid")
	errInvalidDataRateIndex = errors.DefineInvalidArgument("data_rate_index", "Data rate index is invalid")
)

// SetDefaults sets the defaults for the struct where relevant.
//
//nolint:gocyclo
func SetDefaults(m *ttnpb.SimulateMetadataParams) error {
	if m.Time == nil || (m.Time.Nanos == 0 && m.Time.Seconds == 0) {
		m.Time = timestamppb.Now()
	}

	timestamp := m.Time.AsTime()
	if m.Timestamp == 0 {
		m.Timestamp = uint32(timestamp.UnixNano() / 1000)
	}
	if m.BandId == "" {
		m.BandId = band.EU_863_870
	}
	if m.LorawanPhyVersion == ttnpb.PHYVersion_PHY_UNKNOWN {
		m.LorawanPhyVersion = ttnpb.PHYVersion_RP001_V1_0_2_REV_B
	}
	phy, err := band.Get(m.BandId, m.LorawanPhyVersion)
	if err != nil {
		return err
	}
	if m.Frequency == 0 {
		m.Frequency = phy.UplinkChannels[int(m.ChannelIndex)].Frequency
	} else if m.ChannelIndex == 0 {
		chIdx, err := func() (uint32, error) {
			for i, ch := range phy.UplinkChannels {
				if ch.Frequency == m.Frequency {
					return uint32(i), nil
				}
			}
			return 0, errFrequency.New()
		}()
		if err != nil {
			return err
		}
		m.ChannelIndex = chIdx
	}
	if m.Bandwidth == 0 || m.SpreadingFactor == 0 {
		drIdx := ttnpb.DataRateIndex(m.DataRateIndex)
		if drIdx < phy.UplinkChannels[m.ChannelIndex].MinDataRate ||
			drIdx > phy.UplinkChannels[m.ChannelIndex].MaxDataRate {
			drIdx = phy.UplinkChannels[m.ChannelIndex].MaxDataRate
		}
		dr, ok := phy.DataRates[drIdx]
		if !ok {
			return errInvalidDataRateIndex.New()
		}
		lora := dr.Rate.GetLora()
		m.SpreadingFactor, m.Bandwidth = lora.SpreadingFactor, lora.Bandwidth
	} else if m.DataRateIndex == 0 {
		drIdx, err := func() (uint32, error) {
			for i, dr := range phy.DataRates {
				if lora := dr.Rate.GetLora(); lora != nil &&
					lora.SpreadingFactor == m.SpreadingFactor &&
					lora.Bandwidth == m.Bandwidth {
					return uint32(i), nil
				}
			}
			return 0, errDataRate.New()
		}()
		if err != nil {
			return err
		}
		m.DataRateIndex = drIdx
	}
	return nil
}
