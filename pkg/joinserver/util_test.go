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

package joinserver_test

import (
	"strconv"

	"go.thethings.network/lorawan-stack/v3/pkg/provisioning"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/protobuf/types/known/structpb"
)

var Timeout = (1 << 8) * test.Delay

type byteToSerialNumber struct{}

func (p *byteToSerialNumber) UniqueID(entry *structpb.Struct) (string, error) {
	return strconv.Itoa(int(entry.Fields["serial_number"].GetNumberValue())), nil
}

func init() {
	provisioning.Register("mock", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-1", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-2", &byteToSerialNumber{})
	provisioning.Register("mock-provisioner-3", &byteToSerialNumber{})
}
