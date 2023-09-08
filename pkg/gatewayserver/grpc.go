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

package gatewayserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"golang.org/x/sync/errgroup"
)

// GetGatewayConnectionStats returns statistics about a gateway connection.
func (gs *GatewayServer) GetGatewayConnectionStats(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayConnectionStats, error) {
	if err := gs.entityRegistry.AssertGatewayRights(ctx, ids, ttnpb.Right_RIGHT_GATEWAY_STATUS_READ); err != nil {
		return nil, err
	}

	uid := unique.ID(ctx, ids)
	if gs.statsRegistry != nil {
		stats, err := gs.statsRegistry.Get(ctx, ids)
		if err != nil || stats == nil {
			if errors.IsNotFound(err) {
				return nil, errNotConnected.WithAttributes("gateway_uid", uid).WithCause(err)
			}
			return nil, err
		}

		return stats, nil
	}

	val, ok := gs.connections.Load(uid)
	if !ok {
		return nil, errNotConnected.WithAttributes("gateway_uid", uid)
	}
	stats, _ := val.(connectionEntry).Stats()
	return stats, nil
}

func applyGatewayConnectionStatsFieldMask(
	dst, src *ttnpb.GatewayConnectionStats,
	paths ...string,
) (*ttnpb.GatewayConnectionStats, error) {
	if dst == nil {
		dst = &ttnpb.GatewayConnectionStats{}
	}
	return dst, dst.SetFields(src, paths...)
}

// BatchGetGatewayConnectionStats gets statistics about gateway connections to the Gateway Server
// of a batch of gateways.
// This RPC skips unconnected gateways.
// FieldMask paths can be used directly since they are sanitized by the middleware.
func (gs *GatewayServer) BatchGetGatewayConnectionStats(
	ctx context.Context,
	req *ttnpb.BatchGetGatewayConnectionStatsRequest,
) (*ttnpb.BatchGetGatewayConnectionStatsResponse, error) {
	wg, wgCtx := errgroup.WithContext(ctx)
	for _, ids := range req.GatewayIds {
		ids := ids
		wg.Go(func() error {
			return gs.entityRegistry.AssertGatewayRights(wgCtx, ids, ttnpb.Right_RIGHT_GATEWAY_STATUS_READ)
		})
	}
	if err := wg.Wait(); err != nil {
		return nil, err
	}

	if gs.statsRegistry != nil {
		entries, err := gs.statsRegistry.BatchGet(ctx, req.GatewayIds, req.FieldMask.GetPaths()...)
		if err != nil {
			return nil, err
		}
		return &ttnpb.BatchGetGatewayConnectionStatsResponse{
			Entries: entries,
		}, nil
	}

	// If there isn't a registry, load the (ephemeral) values stored in the Gateway Server connections.
	entries := make(map[string]*ttnpb.GatewayConnectionStats, len(req.GatewayIds))
	for _, id := range req.GatewayIds {
		uid := unique.ID(ctx, id)
		val, ok := gs.connections.Load(uid)
		if !ok {
			continue
		}
		st, _ := val.(connectionEntry).Stats()
		if len(req.FieldMask.GetPaths()) > 0 {
			selected, err := applyGatewayConnectionStatsFieldMask(nil, st, req.FieldMask.GetPaths()...)
			if err != nil {
				return nil, err
			}
			st = selected
		}
		entries[id.GatewayId] = st
	}
	return &ttnpb.BatchGetGatewayConnectionStatsResponse{
		Entries: entries,
	}, nil
}
