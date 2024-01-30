package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
	libvirt_commands "github.com/scottshotgg/libvirt-test/pkg/commands/libvirt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (a *API) NodeGroupIncreaseSize(ctx context.Context, req *extgrpc_protos.NodeGroupIncreaseSizeRequest) (*extgrpc_protos.NodeGroupIncreaseSizeResponse, error) {
	for i := 0; i < int(req.GetDelta()); i++ {
		var _, err = a.c.Scale(ctx, req.GetId())
		if err != nil {
			return nil, err
		}
	}

	return &extgrpc_protos.NodeGroupIncreaseSizeResponse{}, nil
}

func (a *API) NodeGroupDeleteNodes(ctx context.Context, req *extgrpc_protos.NodeGroupDeleteNodesRequest) (*extgrpc_protos.NodeGroupDeleteNodesResponse, error) {
	for _, v := range req.GetNodes() {
		var uuidLabel, ok = v.GetLabels()["libvirt-api.uuid"]
		if !ok {
			return nil, fmt.Errorf("no UUID label on provided node: %s", v.GetName())
		}

		uid, err := uuid.Parse(uuidLabel)
		if err != nil {
			return nil, err
		}

		_, err = a.c.StopVM(ctx, uid)
		if err != nil {
			return nil, err
		}
	}

	return &extgrpc_protos.NodeGroupDeleteNodesResponse{}, nil
}

func (a *API) NodeGroupNodes(ctx context.Context, req *extgrpc_protos.NodeGroupNodesRequest) (*extgrpc_protos.NodeGroupNodesResponse, error) {
	var vms, err = a.c.ListVMs(ctx)
	if err != nil {
		return nil, err
	}

	var instances []*extgrpc_protos.Instance

	for _, v := range vms {
		if v.Metadata.GroupID != req.Id || v.State != libvirt_commands.VMState_Running {
			continue
		}

		instances = append(instances, &extgrpc_protos.Instance{
			Id: v.UUID.String(),
			Status: &extgrpc_protos.InstanceStatus{
				InstanceState: extgrpc_protos.InstanceStatus_instanceRunning,
			},
		})
	}

	return &extgrpc_protos.NodeGroupNodesResponse{
		Instances: instances,
	}, nil
}

func (a *API) NodeGroups(ctx context.Context, req *extgrpc_protos.NodeGroupsRequest) (*extgrpc_protos.NodeGroupsResponse, error) {
	// var vms, err = a.c.
	var vms, err = a.c.ListVMs(ctx)
	if err != nil {
		return nil, err
	}

	var (
		nodeGroupsMap = map[string]*extgrpc_protos.NodeGroup{}
		nodeGroups    []*extgrpc_protos.NodeGroup
	)

	for _, vm := range vms {
		var _, ok = nodeGroupsMap[vm.Metadata.GroupID]
		if ok {
			continue
		}

		var ng = extgrpc_protos.NodeGroup{
			Id:      vm.Metadata.GroupID,
			MinSize: 1,
			MaxSize: 5,
			Debug:   "hippity hoo blah",
		}

		nodeGroupsMap[vm.Metadata.GroupID] = &ng
		nodeGroups = append(nodeGroups, &ng)
	}

	return &extgrpc_protos.NodeGroupsResponse{
		NodeGroups: nodeGroups,
	}, nil
}

func (a *API) NodeGroupForNode(ctx context.Context, req *extgrpc_protos.NodeGroupForNodeRequest) (*extgrpc_protos.NodeGroupForNodeResponse, error) {
	var node = req.GetNode()

	var uuidLabel, ok = node.GetLabels()["libvirt-api.uuid"]
	if !ok {
		return nil, fmt.Errorf("no UUID label on provided node: %s", node.GetName())
	}

	uid, err := uuid.Parse(uuidLabel)
	if err != nil {
		return nil, err
	}

	vm, err := a.c.GetVM(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &extgrpc_protos.NodeGroupForNodeResponse{
		NodeGroup: &extgrpc_protos.NodeGroup{
			Id: vm.Metadata.GroupID,
			// TODO: track these values in the "cloud" and add controllability here
			// FIXME: look at some other ones or something
			MinSize: 1,
			MaxSize: 5,
			Debug:   "debug_urself_kid",
		},
	}, nil
}

func (a *API) PricingNodePrice(ctx context.Context, req *extgrpc_protos.PricingNodePriceRequest) (*extgrpc_protos.PricingNodePriceResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (a *API) PricingPodPrice(ctx context.Context, req *extgrpc_protos.PricingPodPriceRequest) (*extgrpc_protos.PricingPodPriceResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (a *API) NodeGroupTemplateNodeInfo(ctx context.Context, req *extgrpc_protos.NodeGroupTemplateNodeInfoRequest) (*extgrpc_protos.NodeGroupTemplateNodeInfoResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (a *API) NodeGroupGetOptions(ctx context.Context, req *extgrpc_protos.NodeGroupAutoscalingOptionsRequest) (*extgrpc_protos.NodeGroupAutoscalingOptionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

func (a *API) GPULabel(ctx context.Context, req *extgrpc_protos.GPULabelRequest) (*extgrpc_protos.GPULabelResponse, error) {
	// TODO: idk I don't think we need this for now but maybe this will conflict with our own operator to turn on nodes with GPUs
	// FIXME: look at some other ones or something
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
	// TODO: Not sure what this means in this context - just leave this here for now
	// FIXME: look at some other ones or something
	return &extgrpc_protos.CleanupResponse{}, nil
}

func (a *API) Refresh(ctx context.Context, req *extgrpc_protos.RefreshRequest) (*extgrpc_protos.RefreshResponse, error) {
	// TODO: Not sure what this means in this context - just leave this here for now
	// FIXME: look at some other ones or something
	return &extgrpc_protos.RefreshResponse{}, nil
}

func (a *API) NodeGroupTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupTargetSizeRequest) (*extgrpc_protos.NodeGroupTargetSizeResponse, error) {
	// TODO: who knows
	// FIXME: look at some other ones or something
	return nil, errors.New("not implemented")
	// return &extgrpc_protos.NodeGroupTargetSizeResponse{
	// 	TargetSize: 1,
	// }, nil
}

func (a *API) NodeGroupDecreaseTargetSize(ctx context.Context, req *extgrpc_protos.NodeGroupDecreaseTargetSizeRequest) (*extgrpc_protos.NodeGroupDecreaseTargetSizeResponse, error) {
	// TODO: who knows - not really sure what to do here? *shrug*
	return nil, errors.New("not implemented")
}
