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

import React, { useCallback } from 'react'

import Form, { useFormContext } from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'

import OwnersSelect from '@console/containers/owners-select'

import sharedMessages from '@ttn-lw/lib/shared-messages'
import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'

import GatewayRegistrationFormSection from './gateway-registration-form-section'

const euiIdRegexp = /eui-\d{16}/

// Empty strings will be interpreted as zero value (00 fill) by the backend
// For this reason, they need to be transformed to undefined instead,
// so that the value will be properly stripped when sent to the backend.
const gatewayEuiEncoder = value => (value === '' ? undefined : value)
const gatewayEuiDecoder = value => (value === undefined ? '' : value)

const GatewayProvisioningFormSection = () => {
  const { touched, setFieldValue, values } = useFormContext()
  const idTouched = touched?.ids?.gateway_id
  const hasEuiId = euiIdRegexp.test(values.ids.gateway_id)

  // Prefill the gateway ID after the EUI is entered.
  const handleEuiBlur = useCallback(
    event => {
      const val = event.target.value
      if (val.length === 16 && (!idTouched || hasEuiId)) {
        setFieldValue('ids.gateway_id', `eui-${val.toLowerCase()}`)
      }
    },
    [hasEuiId, idTouched, setFieldValue],
  )

  return (
    <>
      <OwnersSelect name="_ownerId" required />
      <Form.Field
        title={sharedMessages.gatewayEUI}
        name="ids.eui"
        type="byte"
        min={8}
        max={8}
        placeholder={sharedMessages.gatewayEUI}
        component={Input}
        tooltipId={tooltipIds.GATEWAY_EUI}
        onBlur={handleEuiBlur}
        encode={gatewayEuiEncoder}
        decode={gatewayEuiDecoder}
        autoFocus
      />
      <GatewayRegistrationFormSection />
    </>
  )
}

export { GatewayProvisioningFormSection as default }
