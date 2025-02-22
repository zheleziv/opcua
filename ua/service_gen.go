// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Code generated by cmd/service. DO NOT EDIT!

package ua

import "github.com/zheleziv/opcua/id"

func init() {
	RegisterService(id.ServiceFault_Encoding_DefaultBinary, new(ServiceFault))
	RegisterService(id.FindServersRequest_Encoding_DefaultBinary, new(FindServersRequest))
	RegisterService(id.FindServersResponse_Encoding_DefaultBinary, new(FindServersResponse))
	RegisterService(id.FindServersOnNetworkRequest_Encoding_DefaultBinary, new(FindServersOnNetworkRequest))
	RegisterService(id.FindServersOnNetworkResponse_Encoding_DefaultBinary, new(FindServersOnNetworkResponse))
	RegisterService(id.GetEndpointsRequest_Encoding_DefaultBinary, new(GetEndpointsRequest))
	RegisterService(id.GetEndpointsResponse_Encoding_DefaultBinary, new(GetEndpointsResponse))
	RegisterService(id.RegisterServerRequest_Encoding_DefaultBinary, new(RegisterServerRequest))
	RegisterService(id.RegisterServerResponse_Encoding_DefaultBinary, new(RegisterServerResponse))
	RegisterService(id.RegisterServer2Request_Encoding_DefaultBinary, new(RegisterServer2Request))
	RegisterService(id.RegisterServer2Response_Encoding_DefaultBinary, new(RegisterServer2Response))
	RegisterService(id.OpenSecureChannelRequest_Encoding_DefaultBinary, new(OpenSecureChannelRequest))
	RegisterService(id.OpenSecureChannelResponse_Encoding_DefaultBinary, new(OpenSecureChannelResponse))
	RegisterService(id.CloseSecureChannelRequest_Encoding_DefaultBinary, new(CloseSecureChannelRequest))
	RegisterService(id.CloseSecureChannelResponse_Encoding_DefaultBinary, new(CloseSecureChannelResponse))
	RegisterService(id.CreateSessionRequest_Encoding_DefaultBinary, new(CreateSessionRequest))
	RegisterService(id.CreateSessionResponse_Encoding_DefaultBinary, new(CreateSessionResponse))
	RegisterService(id.ActivateSessionRequest_Encoding_DefaultBinary, new(ActivateSessionRequest))
	RegisterService(id.ActivateSessionResponse_Encoding_DefaultBinary, new(ActivateSessionResponse))
	RegisterService(id.CloseSessionRequest_Encoding_DefaultBinary, new(CloseSessionRequest))
	RegisterService(id.CloseSessionResponse_Encoding_DefaultBinary, new(CloseSessionResponse))
	RegisterService(id.CancelRequest_Encoding_DefaultBinary, new(CancelRequest))
	RegisterService(id.CancelResponse_Encoding_DefaultBinary, new(CancelResponse))
	RegisterService(id.AddNodesRequest_Encoding_DefaultBinary, new(AddNodesRequest))
	RegisterService(id.AddNodesResponse_Encoding_DefaultBinary, new(AddNodesResponse))
	RegisterService(id.AddReferencesRequest_Encoding_DefaultBinary, new(AddReferencesRequest))
	RegisterService(id.AddReferencesResponse_Encoding_DefaultBinary, new(AddReferencesResponse))
	RegisterService(id.DeleteNodesRequest_Encoding_DefaultBinary, new(DeleteNodesRequest))
	RegisterService(id.DeleteNodesResponse_Encoding_DefaultBinary, new(DeleteNodesResponse))
	RegisterService(id.DeleteReferencesRequest_Encoding_DefaultBinary, new(DeleteReferencesRequest))
	RegisterService(id.DeleteReferencesResponse_Encoding_DefaultBinary, new(DeleteReferencesResponse))
	RegisterService(id.BrowseRequest_Encoding_DefaultBinary, new(BrowseRequest))
	RegisterService(id.BrowseResponse_Encoding_DefaultBinary, new(BrowseResponse))
	RegisterService(id.BrowseNextRequest_Encoding_DefaultBinary, new(BrowseNextRequest))
	RegisterService(id.BrowseNextResponse_Encoding_DefaultBinary, new(BrowseNextResponse))
	RegisterService(id.TranslateBrowsePathsToNodeIDsRequest_Encoding_DefaultBinary, new(TranslateBrowsePathsToNodeIDsRequest))
	RegisterService(id.TranslateBrowsePathsToNodeIDsResponse_Encoding_DefaultBinary, new(TranslateBrowsePathsToNodeIDsResponse))
	RegisterService(id.RegisterNodesRequest_Encoding_DefaultBinary, new(RegisterNodesRequest))
	RegisterService(id.RegisterNodesResponse_Encoding_DefaultBinary, new(RegisterNodesResponse))
	RegisterService(id.UnregisterNodesRequest_Encoding_DefaultBinary, new(UnregisterNodesRequest))
	RegisterService(id.UnregisterNodesResponse_Encoding_DefaultBinary, new(UnregisterNodesResponse))
	RegisterService(id.QueryFirstRequest_Encoding_DefaultBinary, new(QueryFirstRequest))
	RegisterService(id.QueryFirstResponse_Encoding_DefaultBinary, new(QueryFirstResponse))
	RegisterService(id.QueryNextRequest_Encoding_DefaultBinary, new(QueryNextRequest))
	RegisterService(id.QueryNextResponse_Encoding_DefaultBinary, new(QueryNextResponse))
	RegisterService(id.ReadRequest_Encoding_DefaultBinary, new(ReadRequest))
	RegisterService(id.ReadResponse_Encoding_DefaultBinary, new(ReadResponse))
	RegisterService(id.HistoryReadRequest_Encoding_DefaultBinary, new(HistoryReadRequest))
	RegisterService(id.HistoryReadResponse_Encoding_DefaultBinary, new(HistoryReadResponse))
	RegisterService(id.WriteRequest_Encoding_DefaultBinary, new(WriteRequest))
	RegisterService(id.WriteResponse_Encoding_DefaultBinary, new(WriteResponse))
	RegisterService(id.HistoryUpdateRequest_Encoding_DefaultBinary, new(HistoryUpdateRequest))
	RegisterService(id.HistoryUpdateResponse_Encoding_DefaultBinary, new(HistoryUpdateResponse))
	RegisterService(id.CallMethodRequest_Encoding_DefaultBinary, new(CallMethodRequest))
	RegisterService(id.CallRequest_Encoding_DefaultBinary, new(CallRequest))
	RegisterService(id.CallResponse_Encoding_DefaultBinary, new(CallResponse))
	RegisterService(id.MonitoredItemCreateRequest_Encoding_DefaultBinary, new(MonitoredItemCreateRequest))
	RegisterService(id.CreateMonitoredItemsRequest_Encoding_DefaultBinary, new(CreateMonitoredItemsRequest))
	RegisterService(id.CreateMonitoredItemsResponse_Encoding_DefaultBinary, new(CreateMonitoredItemsResponse))
	RegisterService(id.MonitoredItemModifyRequest_Encoding_DefaultBinary, new(MonitoredItemModifyRequest))
	RegisterService(id.ModifyMonitoredItemsRequest_Encoding_DefaultBinary, new(ModifyMonitoredItemsRequest))
	RegisterService(id.ModifyMonitoredItemsResponse_Encoding_DefaultBinary, new(ModifyMonitoredItemsResponse))
	RegisterService(id.SetMonitoringModeRequest_Encoding_DefaultBinary, new(SetMonitoringModeRequest))
	RegisterService(id.SetMonitoringModeResponse_Encoding_DefaultBinary, new(SetMonitoringModeResponse))
	RegisterService(id.SetTriggeringRequest_Encoding_DefaultBinary, new(SetTriggeringRequest))
	RegisterService(id.SetTriggeringResponse_Encoding_DefaultBinary, new(SetTriggeringResponse))
	RegisterService(id.DeleteMonitoredItemsRequest_Encoding_DefaultBinary, new(DeleteMonitoredItemsRequest))
	RegisterService(id.DeleteMonitoredItemsResponse_Encoding_DefaultBinary, new(DeleteMonitoredItemsResponse))
	RegisterService(id.CreateSubscriptionRequest_Encoding_DefaultBinary, new(CreateSubscriptionRequest))
	RegisterService(id.CreateSubscriptionResponse_Encoding_DefaultBinary, new(CreateSubscriptionResponse))
	RegisterService(id.ModifySubscriptionRequest_Encoding_DefaultBinary, new(ModifySubscriptionRequest))
	RegisterService(id.ModifySubscriptionResponse_Encoding_DefaultBinary, new(ModifySubscriptionResponse))
	RegisterService(id.SetPublishingModeRequest_Encoding_DefaultBinary, new(SetPublishingModeRequest))
	RegisterService(id.SetPublishingModeResponse_Encoding_DefaultBinary, new(SetPublishingModeResponse))
	RegisterService(id.PublishRequest_Encoding_DefaultBinary, new(PublishRequest))
	RegisterService(id.PublishResponse_Encoding_DefaultBinary, new(PublishResponse))
	RegisterService(id.RepublishRequest_Encoding_DefaultBinary, new(RepublishRequest))
	RegisterService(id.RepublishResponse_Encoding_DefaultBinary, new(RepublishResponse))
	RegisterService(id.TransferSubscriptionsRequest_Encoding_DefaultBinary, new(TransferSubscriptionsRequest))
	RegisterService(id.TransferSubscriptionsResponse_Encoding_DefaultBinary, new(TransferSubscriptionsResponse))
	RegisterService(id.DeleteSubscriptionsRequest_Encoding_DefaultBinary, new(DeleteSubscriptionsRequest))
	RegisterService(id.DeleteSubscriptionsResponse_Encoding_DefaultBinary, new(DeleteSubscriptionsResponse))
}
