/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"context"
	"net"
	"os"
	"sync"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-test/v4/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/yandex-cloud/yc-csi-driver/pkg/services"
)

func TestNodeServer(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "csi-node-test")
	defer os.RemoveAll(tmpDir)
	require.NoError(t, err)

	options := DefaultOptions()
	options.Endpoint = "unix://" + tmpDir + "/csi.sock"
	srv, err := New(options)
	require.NoError(t, err)
	require.NotNil(t, srv)

	listener, err := net.Listen(srv.Scheme, srv.Addr)
	require.NoError(t, err)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		require.NoError(t, srv.Srv.Serve(listener))
		defer wg.Done()
	}()

	defer func() {
		srv.Srv.Stop()
		wg.Wait()
	}()
	conn, err := utils.Connect(
		srv.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	require.NotNil(t, conn)
	n := csi.NewNodeClient(conn)
	require.NotNil(t, n)
	for _, c := range services.ProductionNodeCaps {
		assert.NotEqual(t, csi.NodeServiceCapability_RPC_UNKNOWN, c.GetRpc().GetType())
		assert.True(t, isNodeCapabilitySupported(t, n, c.GetRpc().GetType()))
	}
}

func isNodeCapabilitySupported(
	t *testing.T,
	n csi.NodeClient,
	capType csi.NodeServiceCapability_RPC_Type,
) bool {
	caps, err := n.NodeGetCapabilities(
		context.TODO(),
		&csi.NodeGetCapabilitiesRequest{})
	require.NoError(t, err)
	for _, c := range caps.GetCapabilities() {
		require.NotNil(t, c.GetRpc())
		if c.GetRpc().GetType() == capType {
			return true
		}
	}
	return false
}
