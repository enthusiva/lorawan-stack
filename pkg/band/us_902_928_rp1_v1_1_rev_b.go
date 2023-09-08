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

package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// US_902_928_RP1_V1_1_Rev_B is the band definition for US902-928 in the RP1 v1.1 rev. B specification.
var US_902_928_RP1_V1_1_Rev_B = Band{
	ID: US_902_928,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 72,
	UplinkChannels:    us902928UplinkChannels(0),

	MaxDownlinkChannels: 8,
	DownlinkChannels:    us902928DownlinkChannels,

	// As per FCC Rules for Unlicensed Wireless Equipment operating in the ISM bands
	SubBands: []SubBandParameters{
		{
			MinFrequency: 902300000,
			MaxFrequency: 914900000,
			DutyCycle:    1,
			MaxEIRP:      21.0 + eirpDelta,
		},
		{
			MinFrequency: 923300000,
			MaxFrequency: 927500000,
			DutyCycle:    1,
			MaxEIRP:      26.0 + eirpDelta,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(10, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(19)),
		ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(9, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(61)),
		ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(8, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(133)),
		ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(7, 125000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(8, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),

		ttnpb.DataRateIndex_DATA_RATE_8:  makeLoRaDataRate(12, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(61)),
		ttnpb.DataRateIndex_DATA_RATE_9:  makeLoRaDataRate(11, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(137)),
		ttnpb.DataRateIndex_DATA_RATE_10: makeLoRaDataRate(10, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_11: makeLoRaDataRate(9, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_12: makeLoRaDataRate(8, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
		ttnpb.DataRateIndex_DATA_RATE_13: makeLoRaDataRate(7, 500000, Cr4_5, makeConstMaxMACPayloadSizeFunc(250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
	StrictCodingRate:    true,

	DefaultMaxEIRP: 30,
	TxOffset: []float32{
		0,
		-2,
		-4,
		-6,
		-8,
		-10,
		-12,
		-14,
		-16,
		-18,
		-20,
		-22,
		-24,
		-26,
		-28,
	},

	Rx1Channel: channelIndexModulo(8),
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		if idx > ttnpb.DataRateIndex_DATA_RATE_4 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 4)
		}
		if offset > 3 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 3)
		}
		return us902928DownlinkDRTable[idx][offset], nil
	},

	GenerateChMasks: makeGenerateChMask72(true, true),
	ParseChMask:     parseChMask72,

	FreqMultiplier:   100,
	ImplementsCFList: true,
	CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_8, 923300000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_8,
		CodingRate:    Cr4_5,
		Frequencies:   usAuBeaconFrequencies,
	},
	PingSlotFrequencies: usAuBeaconFrequencies,

	BootDwellTime: DwellTime{
		Uplinks:   boolPtr(true),
		Downlinks: boolPtr(false),
	},

	SharedParameters: universalSharedParameters,
}
