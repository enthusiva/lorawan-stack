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

import {
  createPaginationIdsSelectorByEntity,
  createPaginationTotalCountSelectorByEntity,
} from '@ttn-lw/lib/store/selectors/pagination'

const ENTITY = 'apiKeys'

// Api key.
export const selectApiKeysStore = state => state.apiKeys || {}
export const selectApiKeysEntitiesStore = state => selectApiKeysStore(state).entities
export const selectApiKeyById = (state, id) => selectApiKeysEntitiesStore(state)[id]
export const selectSelectedApiKeyId = state => selectApiKeysStore(state).selectedApiKey
export const selectSelectedApiKey = state => selectApiKeyById(state, selectSelectedApiKeyId(state))

// Api keys.
const createSelectApiKeysIdsSelector = createPaginationIdsSelectorByEntity(ENTITY)
const createSelectApiKeysTotalCountSelector = createPaginationTotalCountSelectorByEntity(ENTITY)

export const selectApiKeys = state =>
  createSelectApiKeysIdsSelector(state).map(id => selectApiKeyById(state, id))
export const selectApiKeysTotalCount = state => createSelectApiKeysTotalCountSelector(state)
