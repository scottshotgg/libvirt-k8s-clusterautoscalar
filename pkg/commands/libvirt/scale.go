package libvirt

import (
	"context"
	"fmt"
	"time"

	"github.com/scottshotgg/libvirt-test/pkg/commands"
)

func (c *Libvirt) Scale(ctx context.Context, groupID string) (*commands.VMInfo, error) {
	var vms, err = c.ListVMs(ctx)
	if err != nil {
		return nil, err
	}

	if len(vms) == 0 {
		return nil, ErrNoVMsFound
	}

	var (
		info         *commands.VMInfo
		templatePath string
	)

	// Look through the VMs, find one that is in a node group
	// If one is found then just start it, else create a new one
	for _, vm := range vms {
		var vminfo, err = c.GetVMInfo(ctx, vm.UUID)
		if err != nil {
			// TODO: might just need to continue here
			return nil, err
		}

		if vminfo.Metadata == nil {
			continue
		}

		// We have found a VM in the same group that isn't
		// started so we don't need to create a new one
		if vminfo.Metadata.GroupID == groupID {
			templatePath = vminfo.Metadata.TemplatePath
		}

		if vminfo.State == VMState_Shutdown {
			fmt.Printf("Found VM: %+v\n", vminfo)
			info = vminfo

			break
		}
	}

	// If we don't have a template path then this means that we didn't find a node group and we can't install
	if templatePath == "" {
		return nil, ErrVMGroupNotFound
	}

	if info == nil {
		fmt.Println("Creating VM ...")

		// TODO: think of something else for the name, maybe use sqlite or something idk
		info, err = c.CreateVM(ctx, groupID, templatePath, fmt.Sprintf("k8s_worker_node_%d", time.Now().Unix()))
		if err != nil {
			return nil, err
		}

		fmt.Println("Created VM:", info)
	}

	fmt.Println("Starting VM ...")

	return info, c.StartVM(ctx, info.UUID)
}
