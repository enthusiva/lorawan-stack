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

import React, { useState, useCallback } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import classnames from 'classnames'

import Input from '@ttn-lw/components/input'
import Form, { useFormContext } from '@ttn-lw/components/form'

import Message from '@ttn-lw/lib/components/message'

import tooltipIds from '@ttn-lw/lib/constants/tooltip-ids'
import env from '@ttn-lw/lib/env'
import sharedMessages from '@ttn-lw/lib/shared-messages'
import PropTypes from '@ttn-lw/lib/prop-types'
import attachPromise from '@ttn-lw/lib/store/actions/attach-promise'
import { getBackendErrorName } from '@ttn-lw/lib/errors/utils'

import { getApplicationDevEUICount, issueDevEUI } from '@console/store/actions/applications'

import {
  selectApplicationDevEUICount,
  selectSelectedApplicationId,
} from '@console/store/selectors/applications'

import style from './dev-eui.styl'

const DevEUIComponent = props => {
  const { name, required, disabled, autoFocus } = props
  const { values, setFieldValue, touched } = useFormContext()

  const dispatch = useDispatch()
  const appId = useSelector(selectSelectedApplicationId)
  const promisifiedIssueDevEUI = attachPromise(issueDevEUI)
  const fetchDevEUICounter = attachPromise(getApplicationDevEUICount)
  const [devEUIGenerated, setDevEUIGenerated] = useState(false)
  const [errorMessage, setErrorMessage] = useState(undefined)
  const applicationDevEUICounter = useSelector(selectApplicationDevEUICount)
  const idTouched = touched?.ids?.device_id || touched?.target_device_id
  const hasEuiId =
    /eui-\d{16}/.test(values?.target_device_id) || /eui-\d{16}/.test(values?.ids?.device_id)

  const indicatorContent = errorMessage || {
    ...sharedMessages.used,
    values: {
      currentValue: applicationDevEUICounter,
      maxValue: env.devEUIConfig.applicationLimit,
    },
  }

  const indicatorCls = classnames(style.indicator, {
    'tc-error':
      applicationDevEUICounter === env.devEUIConfig.applicationLimit || Boolean(errorMessage),
  })

  const handleDevEUIRequest = useCallback(async () => {
    const result = await dispatch(promisifiedIssueDevEUI(appId))
    await dispatch(fetchDevEUICounter(appId))
    return result.dev_eui
  }, [appId, dispatch, fetchDevEUICounter, promisifiedIssueDevEUI])

  const handleIdPrefill = useCallback(
    eventOrEui => {
      const value = typeof eventOrEui === 'string' ? eventOrEui : eventOrEui.target.value
      if (value.length === 16 && (!idTouched || hasEuiId)) {
        const generatedId = `eui-${value.toLowerCase()}`
        setFieldValue('target_device_id', generatedId)
        setFieldValue('ids.device_id', generatedId)
      }
    },
    [hasEuiId, idTouched, setFieldValue],
  )

  const handleGenerate = useCallback(async () => {
    try {
      const result = await handleDevEUIRequest()
      setDevEUIGenerated(true)
      setErrorMessage(undefined)
      handleIdPrefill(result)
      return result
    } catch (error) {
      if (getBackendErrorName(error) === 'global_eui_limit_reached') {
        setErrorMessage(sharedMessages.devEUIBlockLimitReached)
      } else setErrorMessage(sharedMessages.unknownError)
      setDevEUIGenerated(true)
    }
  }, [handleDevEUIRequest, handleIdPrefill])

  const devEUIGenerateDisabled =
    applicationDevEUICounter === env.devEUIConfig.applicationLimit ||
    !env.devEUIConfig.devEUIIssuingEnabled ||
    devEUIGenerated

  return env.devEUIConfig.devEUIIssuingEnabled ? (
    <Form.Field
      title={sharedMessages.devEUI}
      name={name}
      type="byte"
      min={8}
      max={8}
      component={Input.Generate}
      tooltipId={tooltipIds.DEV_EUI}
      onBlur={handleIdPrefill}
      onGenerateValue={handleGenerate}
      actionDisable={devEUIGenerateDisabled}
      required={required}
      disabled={disabled}
      autoFocus={autoFocus}
    >
      <Message className={indicatorCls} component="label" content={indicatorContent} />
    </Form.Field>
  ) : (
    <Form.Field
      title={sharedMessages.devEUI}
      name={name}
      type="byte"
      min={8}
      max={8}
      required={required}
      component={Input}
      tooltipId={tooltipIds.DEV_EUI}
      onBlur={handleIdPrefill}
      disabled={disabled}
      autoFocus={autoFocus}
    />
  )
}

DevEUIComponent.propTypes = {
  autoFocus: PropTypes.bool,
  disabled: PropTypes.bool,
  name: PropTypes.string.isRequired,
  required: PropTypes.bool,
}

DevEUIComponent.defaultProps = {
  autoFocus: false,
  disabled: false,
  required: false,
}

export default DevEUIComponent
