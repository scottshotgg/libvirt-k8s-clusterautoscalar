package api

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
	"google.golang.org/protobuf/types/known/anypb"
	extgrpc_protos "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/externalgrpc/protos"
)

type (
	API struct {
		extgrpc_protos.UnimplementedCloudProviderServer

		c commands.Commands
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
	// var vms, err = a.c.
	var vms, err = a.c.ListVMs(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Collect distinct node groups
	_ = vms

	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupForNode(ctx context.Context, req *extgrpc_protos.NodeGroupForNodeRequest) (*extgrpc_protos.NodeGroupForNodeResponse, error) {
	// TODO: might need to get the UUID from the annotations/labels or just use Name in the commands
	// req.GetNode().GetName()
	var vm, err = a.c.GetVM(ctx, uuid.New())
	if err != nil {
		return nil, err
	}

	// TODO: return node group ID
	return &extgrpc_protos.NodeGroupForNodeResponse{
		NodeGroup: &extgrpc_protos.NodeGroup{
			Id: vm.Metadata.GroupID,
			// TODO: track these values in the "cloud" and add controllability here
			MinSize: 1,
			MaxSize: 5,
			Debug:   "debug_urself_kid",
		},
	}, nil
}

func (a *API) GPULabel(ctx context.Context, req *extgrpc_protos.GPULabelRequest) (*extgrpc_protos.GPULabelResponse, error) {
	// TODO: idk I don't think we need this for now but maybe this will conflict with our own operator to turn on nodes with GPUs
	return &extgrpc_protos.GPULabelResponse{
		Label: "nvidia_beep_boop",
	}, nil
}

func (a *API) GetAvailableGPUTypes(ctx context.Context, req *extgrpc_protos.GetAvailableGPUTypesRequest) (*extgrpc_protos.GetAvailableGPUTypesResponse, error) {
	// TODO: Again - not really sure. Imo I think we should aim to keep these as "unimplemented" as possible without breaking shit
	return &extgrpc_protos.GetAvailableGPUTypesResponse{
		GpuTypes: map[string]*anypb.Any{},
	}, nil
}

func (a *API) Cleanup(ctx context.Context, req *extgrpc_protos.CleanupRequest) (*extgrpc_protos.CleanupResponse, error) {
	// TODO: look at some other ones or something
	return &extgrpc_protos.CleanupResponse{}, nil
}

func (a *API) Refresh(ctx context.Context, req *extgrpc_protos.RefreshRequest) (*extgrpc_protos.RefreshResponse, error) {
	// TODO: Not sure what this means in this context
	return &extgrpc_protos.RefreshResponse{}, nil
}

func (a *API) NodeGroupTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupTargetSizeRequest) (*extgrpc_protos.NodeGroupTargetSizeResponse, error) {
	// TODO: who knows
	return nil, errors.New("not implemented")
	// return &extgrpc_protos.NodeGroupTargetSizeResponse{
	// 	TargetSize: 1,
	// }, nil
}

func (a *API) NodeGroupIncreaseSize(ctx context.Context, req *extgrpc_protos.NodeGroupIncreaseSizeRequest) (*extgrpc_protos.NodeGroupIncreaseSizeResponse, error) {
	for i := 0; i < int(req.GetDelta()); i++ {
		var _, err = a.c.Scale(ctx, req.GetId())
		if err != nil {
			return nil, err
		}
	}

	return &extgrpc_protos.NodeGroupIncreaseSizeResponse{}, nil

	// return nil, errors.New("not implemented")
}

func (a *API) NodeGroupDeleteNodes(ctx context.Context, req *extgrpc_protos.NodeGroupDeleteNodesRequest) (*extgrpc_protos.NodeGroupDeleteNodesResponse, error) {
	// TODO: For loop over nodes to delete; DONT actually delete them - just shut them off
	for _, v := range req.GetNodes() {
		// TODO: might need to get the UUID from the annotations/labels or just use Name in the commands

		// TODO: use v.Name or v.Labels.UUID or w/e
		_ = v

		var _, err = a.c.StopVM(ctx, uuid.New())
		if err != nil {
			return nil, err
		}
	}

	return &extgrpc_protos.NodeGroupDeleteNodesResponse{}, nil

	// return nil, errors.New("not implemented")
}

func (a *API) NodeGroupDecreaseTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupDecreaseTargetSizeRequest) (*extgrpc_protos.NodeGroupDecreaseTargetSizeResponse, error) {
	// TODO: who knows - not really sure what to do here? *shrug*
	return nil, errors.New("not implemented")
}

func (a *API) NodeGroupNodes(ctx context.Context, req *extgrpc_protos.NodeGroupNodesRequest) (*extgrpc_protos.NodeGroupNodesResponse, error) {
	// Look through nodes and collect the ones that match the node group ID
	var res, err = a.NodeGroups(ctx, &extgrpc_protos.NodeGroupsRequest{})
	if err != nil {
		return nil, err
	}

	// TODO: have a central function that organizes these into maps ... and then just pluck the one we need
	_ = res

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
