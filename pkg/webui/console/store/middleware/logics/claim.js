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

import tts from '@console/api/tts'

import createRequestLogic from '@ttn-lw/lib/store/logics/create-request-logic'

import * as claim from '@console/store/actions/claim'

const claimDeviceLogic = createRequestLogic({
  type: claim.CLAIM_DEVICE,
  process: async ({ action }) => {
    const { appId, qr_code, authenticatedIdentifiers } = action.payload

    return await tts.DeviceClaim.claim(appId, qr_code, authenticatedIdentifiers)
  },
})

const unclaimDeviceLogic = createRequestLogic({
  type: claim.UNCLAIM_DEVICE,
  process: async ({ action }) => {
    const { applicationId, deviceId, devEui, joinEui } = action.payload

    return await tts.DeviceClaim.unclaim(applicationId, deviceId, devEui, joinEui)
  },
})

const getInfoByJoinEUILogic = createRequestLogic({
  type: claim.GET_INFO_BY_JOIN_EUI,
  process: async ({ action }) => {
    const { joinEUI } = action.payload

    return await tts.DeviceClaim.GetInfoByJoinEUI(joinEUI)
  },
})

export default [claimDeviceLogic, unclaimDeviceLogic, getInfoByJoinEUILogic]
