package api

import (
	"context"
	"errors"

	extgrpc_protos "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/externalgrpc/protos"
)

type (
	API struct {
		extgrpc_protos.UnimplementedCloudProviderServer
	}
)

var (
	_ extgrpc_protos.CloudProviderServer = (*API)(nil)
)

func New() *API {
	return &API{}
}

/*
	https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/cloudprovider/externalgrpc/protos/externalgrpc_grpc.pb.go
*/

func (a *API) NodeGroups(ctx context.Context, req *extgrpc_protos.NodeGroupsRequest) (*extgrpc_protos.NodeGroupsResponse, error) {
	// TODO: look through all domains and get distinct node groups
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupForNode(ctx context.Context, req *extgrpc_protos.NodeGroupForNodeRequest) (*extgrpc_protos.NodeGroupForNodeResponse, error) {
	// TODO: return node group ID
	return nil, errors.New("not implemented")
}

func (a *API) GPULabel(ctx context.Context, req *extgrpc_protos.GPULabelRequest) (*extgrpc_protos.GPULabelResponse, error) {
	// TODO: idk I don't think we need this for now but maybe this will conflict with our own operator to turn on nodes with GPUs
	return nil, errors.New("not implemented")
}

func (a *API) GetAvailableGPUTypes(ctx context.Context, req *extgrpc_protos.GetAvailableGPUTypesRequest) (*extgrpc_protos.GetAvailableGPUTypesResponse, error) {
	// TODO: Again - not really sure. Imo I think we should aim to keep these as "unimplemented" as possible without breaking shit
	return nil, errors.New("not implemented")
}

func (a *API) Cleanup(ctx context.Context, req *extgrpc_protos.CleanupRequest) (*extgrpc_protos.CleanupResponse, error) {
	// TODO: look at some other ones or something
	return nil, errors.New("not implemented")
}

func (a *API) Refresh(ctx context.Context, req *extgrpc_protos.RefreshRequest) (*extgrpc_protos.RefreshResponse, error) {
	// TODO: Not sure what this means in this context
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupTargetSizeRequest) (*extgrpc_protos.NodeGroupTargetSizeResponse, error) {
	// TODO: who knows
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupIncreaseSize(ctx context.Context, req *extgrpc_protos.NodeGroupIncreaseSizeRequest) (*extgrpc_protos.NodeGroupIncreaseSizeResponse, error) {
	// TODO: Call "Scale"
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupDeleteNodes(ctx context.Context, req *extgrpc_protos.NodeGroupDeleteNodesRequest) (*extgrpc_protos.NodeGroupDeleteNodesResponse, error) {
	// TODO: For loop over nodes to delete; DONT actually delete them - just shut them off
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupDecreaseTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupDecreaseTargetSizeRequest) (*extgrpc_protos.NodeGroupDecreaseTargetSizeResponse, error) {
	// TODO: who knows
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupNodes(ctx context.Context, req *extgrpc_protos.NodeGroupNodesRequest) (*extgrpc_protos.NodeGroupNodesResponse, error) {
	// Look through nodes and collect the ones that match the node group ID
	return nil, errors.New("not implemented")
}

// ---

func (a *API) PricingNodePrice(ctx context.Context, req *extgrpc_protos.PricingNodePriceRequest) (*extgrpc_protos.PricingNodePriceResponse, error) {
	// TODO: Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	return nil, errors.New("not implemented")
}

func (a *API) PricingPodPrice(ctx context.Context, req *extgrpc_protos.PricingPodPriceRequest) (*extgrpc_protos.PricingPodPriceResponse, error) {
	// TODO: Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupTemplateNodeInfo(ctx context.Context, req *extgrpc_protos.NodeGroupTemplateNodeInfoRequest) (*extgrpc_protos.NodeGroupTemplateNodeInfoResponse, error) {
	// TODO: Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupGetOptions(ctx context.Context, req *extgrpc_protos.NodeGroupAutoscalingOptionsRequest) (*extgrpc_protos.NodeGroupAutoscalingOptionsResponse, error) {
	// TODO: Implementation optional: if unimplemented return error code 12 (for `Unimplemented`)
	return nil, errors.New("not implemented")
}
