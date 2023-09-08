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

import React from 'react'
import { defineMessages } from 'react-intl'

import Link from '@ttn-lw/components/link'

import Message from '@ttn-lw/lib/components/message'

import PropTypes from '@ttn-lw/lib/prop-types'

const m = defineMessages({
  hintMessage:
    'Cannot find your exact end device? <SupportLink>Get help here</SupportLink> and try <b>enter end device specifics manually</b> option above.',
  hintNoSupportMessage:
    'Cannot find your exact end device? Try <b>enter end device specifics manually</b> option above.',
})

const ProgressHint = React.memo(props => {
  const { supportLink } = props

  return (
    <Message
      className="mb-ls-xs"
      component="div"
      content={Boolean(supportLink) ? m.hintMessage : m.hintNoSupportMessage}
      values={{
        SupportLink: msg => (
          <Link.Anchor secondary key="support-link" href={supportLink} target="_blank">
            {msg}
          </Link.Anchor>
        ),
        b: msg => <b>{msg}</b>,
      }}
    />
  )
})

ProgressHint.propTypes = {
  supportLink: PropTypes.string,
}

ProgressHint.defaultProps = {
  supportLink: undefined,
}

export default ProgressHint
