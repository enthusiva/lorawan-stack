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

package cluster_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmiddleware/rpclog"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

var ctx context.Context

func init() {
	logHandler, err := log.NewZap("console")
	if err != nil {
		panic(err)
	}
	logger := log.NewLogger(
		logHandler,
		log.WithLevel(log.DebugLevel),
	)
	ctx = log.NewContext(test.Context(), logger.WithField("namespace", "cluster"))
	rpclog.ReplaceGrpcLogger(logger.WithField("namespace", "grpc"))
}

func TestCluster(t *testing.T) {
	a := assertions.New(t)

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	go grpc.NewServer().Serve(lis)

	config := Config{
		Address:                    lis.Addr().String(),
		IdentityServer:             lis.Addr().String(),
		GatewayServer:              lis.Addr().String(),
		NetworkServer:              lis.Addr().String(),
		ApplicationServer:          lis.Addr().String(),
		JoinServer:                 lis.Addr().String(),
		GatewayConfigurationServer: lis.Addr().String(),
		Join:                       []string{lis.Addr().String()},
	}

	ctx := test.Context()

	c, err := New(ctx, &config)
	a.So(err, should.BeNil)

	a.So(c.Join(), should.BeNil)

	grpc.Dial(lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())

	// The Identity Server playing the ACCESS role should be there within reasonable time.
	var ac Peer
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond) // Wait for peers to join cluster.
		ac, err = c.GetPeer(ctx, ttnpb.ClusterRole_ACCESS, nil)
		if err == nil {
			break
		}
	}
	if !a.So(ac, should.NotBeNil) {
		t.FailNow()
	}

	er, err := c.GetPeer(ctx, ttnpb.ClusterRole_ENTITY_REGISTRY, nil)
	a.So(er, should.NotBeNil)
	a.So(err, should.BeNil)
	gs, err := c.GetPeer(ctx, ttnpb.ClusterRole_GATEWAY_SERVER, nil)
	a.So(gs, should.NotBeNil)
	a.So(err, should.BeNil)
	ns, err := c.GetPeer(ctx, ttnpb.ClusterRole_NETWORK_SERVER, nil)
	a.So(ns, should.NotBeNil)
	a.So(err, should.BeNil)
	as, err := c.GetPeer(ctx, ttnpb.ClusterRole_APPLICATION_SERVER, nil)
	a.So(as, should.NotBeNil)
	a.So(err, should.BeNil)
	js, err := c.GetPeer(ctx, ttnpb.ClusterRole_JOIN_SERVER, nil)
	a.So(js, should.NotBeNil)
	a.So(err, should.BeNil)
	gcs, err := c.GetPeer(ctx, ttnpb.ClusterRole_GATEWAY_CONFIGURATION_SERVER, nil)
	a.So(gcs, should.NotBeNil)
	a.So(err, should.BeNil)

	// Test Packet Broker Agent override; Packet Broker Agent is not in the cluster.
	pba, err := c.GetPeer(ctx, ttnpb.ClusterRole_GATEWAY_SERVER, PacketBrokerGatewayID)
	a.So(pba, should.BeNil)
	a.So(err, should.NotBeNil)

	a.So(c.Leave(), should.BeNil)

	for _, peer := range []Peer{
		ac, er, gs, ns, as, js, gcs,
	} {
		cc, err := peer.Conn()
		a.So(cc, should.NotBeNil)
		a.So(err, should.BeNil)
		a.So(cc.GetState(), should.Equal, connectivity.Shutdown)
	}
}
