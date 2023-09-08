// Copyright © 2023 The Things Network Foundation, The Things Industries B.V.
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

package io

// MessageStream is a message stream.
type MessageStream uint32

const (
	UplinkStream   MessageStream = iota // UplinkStream is the uplink message stream.
	DownlinkStream                      // DownlinkStream is the downlink message stream.
	TxAckStream                         // TxAckStream is the transmission acknowledgment stream.
	StatusStream                        // StatusStream is the status message stream.
	RTTStream                           // RTTStream is the round-trip times stream.
)

func alwaysOnStreamState(MessageStream) bool { return true }
