//go:build integration
// +build integration

package uatest

import (
	"context"
	"testing"

	"github.com/pascaldekloe/goe/verify"

	"github.com/zheleziv/opcua"
	"github.com/zheleziv/opcua/ua"
)

// TestRead performs an integration test to read values
// from an OPC/UA server.
func TestRead(t *testing.T) {
	tests := []struct {
		id *ua.NodeID
		v  interface{}
	}{
		{ua.NewStringNodeID(2, "ro_bool"), true},
		{ua.NewStringNodeID(2, "rw_bool"), true},
		{ua.NewStringNodeID(2, "ro_int32"), int32(5)},
		{ua.NewStringNodeID(2, "rw_int32"), int32(5)},
		{ua.NewStringNodeID(2, "array_int32"), []int32{1, 2, 3}},
		{ua.NewStringNodeID(2, "2d_array_int32"), [][]int32{{1}, {2}, {3}}},
	}

	ctx := context.Background()

	srv := NewServer("rw_server.py")
	defer srv.Close()

	c, err := opcua.NewClient(srv.Endpoint, srv.Opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Connect(ctx); err != nil {
		t.Fatal(err)
	}
	defer c.Close(ctx)

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			t.Run("Read", func(t *testing.T) {
				testRead(t, ctx, c, tt.v, tt.id)
			})
			t.Run("RegisteredRead", func(t *testing.T) {
				testRegisteredRead(t, ctx, c, tt.v, tt.id)
			})
		})
	}
}

func testRead(t *testing.T, ctx context.Context, c *opcua.Client, v interface{}, id *ua.NodeID) {
	t.Helper()

	resp, err := c.Read(ctx, &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{
			&ua.ReadValueID{NodeID: id},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	})
	if err != nil {
		t.Fatalf("Read failed: %s", err)
	}
	if resp.Results[0].Status != ua.StatusOK {
		t.Fatalf("Status not OK: %v", resp.Results[0].Status)
	}
	if got, want := resp.Results[0].Value.Value(), v; !verify.Values(t, "", got, want) {
		t.Fail()
	}
}

func testRegisteredRead(t *testing.T, ctx context.Context, c *opcua.Client, v interface{}, id *ua.NodeID) {
	t.Helper()

	resp, err := c.RegisterNodes(ctx, &ua.RegisterNodesRequest{
		NodesToRegister: []*ua.NodeID{id},
	})
	if err != nil {
		t.Fatalf("RegisterNodes failed: %s", err)
	}

	testRead(t, ctx, c, v, resp.RegisteredNodeIDs[0])
	testRead(t, ctx, c, v, resp.RegisteredNodeIDs[0])
	testRead(t, ctx, c, v, resp.RegisteredNodeIDs[0])
	testRead(t, ctx, c, v, resp.RegisteredNodeIDs[0])
	testRead(t, ctx, c, v, resp.RegisteredNodeIDs[0])

	_, err = c.UnregisterNodes(ctx, &ua.UnregisterNodesRequest{
		NodesToUnregister: []*ua.NodeID{id},
	})
	if err != nil {
		t.Fatalf("UnregisterNodes failed: %s", err)
	}
}
