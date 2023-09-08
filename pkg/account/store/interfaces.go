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

package store

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
)

// Interface is the store used by the account app.
type Interface interface {
	// UserStore, LoginTokenStore and UserSessionStore are needed for user login/logout.
	store.UserStore
	store.LoginTokenStore
	store.UserSessionStore
}

// TransactionalStore is Interface, but with a method that uses a transaction.
type TransactionalInterface interface {
	Interface

	// Transact runs a transaction using the store.
	Transact(context.Context, func(context.Context, Interface) error) error
}

// WithSoftDeleted returns a context that tells the store to include (only) deleted entities.
var WithSoftDeleted = store.WithSoftDeleted
