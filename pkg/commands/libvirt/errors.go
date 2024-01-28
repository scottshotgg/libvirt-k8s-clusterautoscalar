package libvirt

import (
	"errors"
)

var (
	ErrNotManaged        = errors.New("This VM is not managed by libvirt-api")
	ErrNoVMsFound        = errors.New("No VMs found")
	ErrNoVMGroupsToScale = errors.New("No VM groups to scale")
	ErrVMGroupNotFound   = errors.New("VM group not found")
)
