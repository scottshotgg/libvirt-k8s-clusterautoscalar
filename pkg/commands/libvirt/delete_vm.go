package libvirt

import (
	"context"
	"os"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func (c *Libvirt) DeleteVM(ctx context.Context, uuid uuid.UUID) error {
	var info, err = c.GetVMInfo(ctx, uuid)
	if err != nil {
		return err
	}

	if info.Metadata == nil {
		return ErrNotManaged
	}

	err = c.l.DomainUndefine(libvirt.Domain{
		UUID: libvirt.UUID(info.UUID),
	})

	if err != nil {
		return err
	}

	err = os.RemoveAll(info.Metadata.RootDir)
	if err != nil {
		return err
	}

	return nil
}
