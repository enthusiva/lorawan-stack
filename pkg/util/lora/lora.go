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

// Package lora contains LoRa modulation utilities.
package lora

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AdjustedRSSI returns the LoRa RSSI: the channel RSSI adjusted for SNR.
// Below -5 dB, the SNR is added to the channel RSSI.
// Between -5 dB and 10 dB, the SNR is scaled to 0 and added to the channel RSSI.
func AdjustedRSSI(channelRSSI, snr float32) float32 {
	rssi := channelRSSI
	if snr <= -5.0 {
		rssi += snr
	} else if snr < 10.0 {
		rssi += snr/3.0 - 10.0/3.0
	}
	return rssi
}

// UplinkMessage describes a message that holds RxMetadata and timestamp.
type UplinkMessage interface {
	GetRxMetadata() []*ttnpb.RxMetadata
	GetReceivedAt() *timestamppb.Timestamp
}

// GetAdjustedReceivedAt tries to improve the ReceivedAt timestamp using
// the message's GpsTime and metadatas' ReceivedAt.
func GetAdjustedReceivedAt(up UplinkMessage) *timestamppb.Timestamp {
	var ts *timestamppb.Timestamp
	for _, md := range up.GetRxMetadata() {
		if t := md.GpsTime; t != nil {
			return t
		}
		if ts == nil && md.ReceivedAt != nil {
			ts = md.ReceivedAt
		}
	}
	if ts != nil {
		return ts
	}
	return up.GetReceivedAt()
}
