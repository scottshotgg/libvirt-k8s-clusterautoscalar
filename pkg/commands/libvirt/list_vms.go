package libvirt

import (
	"context"

	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
)

func (c *Libvirt) ListVMs(ctx context.Context) ([]*commands.VMInfo, error) {
	var domains, err = c.l.Domains()
	if err != nil {
		return nil, err
	}

	var vminfos []*commands.VMInfo
	for _, v := range domains {
		// TODO: I don't think we need to make _all_ of these calls but meh
		info, err := c.GetVMInfo(ctx, uuid.UUID(v.UUID))
		if err != nil {
			return nil, err
		}

		if info.Metadata == nil {
			continue
		}

		vminfos = append(vminfos, info)
	}

	return vminfos, nil
}
