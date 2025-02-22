// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"log"

	"github.com/zheleziv/opcua"
	"github.com/zheleziv/opcua/debug"
	"github.com/zheleziv/opcua/ua"
)

func main() {
	var (
		endpoint = flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
		nodeID   = flag.String("node", "", "NodeID to read")
	)
	flag.BoolVar(&debug.Enable, "debug", false, "enable debug logging")
	flag.Parse()
	log.SetFlags(0)

	ctx := context.Background()

	c, err := opcua.NewClient(*endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close(ctx)

	id, err := ua.ParseNodeID(*nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %v", err)
	}

	regResp, err := c.RegisterNodes(ctx, &ua.RegisterNodesRequest{
		NodesToRegister: []*ua.NodeID{id},
	})
	if err != nil {
		log.Fatalf("RegisterNodes failed: %v", err)
	}

	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			{NodeID: regResp.RegisteredNodeIDs[0]},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	resp, err := c.Read(ctx, req)
	if err != nil {
		log.Fatalf("Read failed: %s", err)
	}
	if resp.Results[0].Status != ua.StatusOK {
		log.Fatalf("Status not OK: %v", resp.Results[0].Status)
	}
	log.Print(resp.Results[0].Value.Value())

	_, err = c.UnregisterNodes(ctx, &ua.UnregisterNodesRequest{
		NodesToUnregister: []*ua.NodeID{id},
	})
	if err != nil {
		log.Fatalf("UnregisterNodes failed: %v", err)
	}
}
