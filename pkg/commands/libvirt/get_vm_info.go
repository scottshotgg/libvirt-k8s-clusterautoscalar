package libvirt

import (
	"context"
	"encoding/xml"
	"errors"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
)

// Change this to return better info
func (c *Libvirt) GetVMInfo(ctx context.Context, uuid uuid.UUID) (*commands.VMInfo, error) {
	var d, err = c.l.DomainLookupByUUID(libvirt.UUID(uuid))
	if err != nil {
		return nil, err
	}

	state, maxMem, memory, vCPUs, cpuTime, err := c.l.DomainGetInfo(d)
	if err != nil {
		return nil, err
	}

	metadata, err := c.l.DomainGetMetadata(d, 2, uris, 0)
	if err != nil {
		var lvErr libvirt.Error
		if !errors.As(err, &lvErr) {
			return nil, err
		}

		if lvErr.Code != uint32(libvirt.ErrNoDomainMetadata) {
			// TODO: maybe this will return a 'not found'?
			return nil, err
		}
	}

	var md *commands.VMMetadata
	if metadata != "" {
		err = xml.Unmarshal([]byte(metadata), &md)
		if err != nil {
			return nil, err
		}
	}

	return &commands.VMInfo{
		ID:       int(d.ID),
		Name:     d.Name,
		UUID:     uuid,
		State:    int(state),
		Metadata: md,
		Resources: commands.VMResources{
			MaxMem:     int(maxMem),
			CurrentMem: int(memory),
			CPUs:       int(vCPUs),
			CPUTime:    int(cpuTime),
		},
	}, nil
}
