package libvirt

import (
	"context"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func (c *Libvirt) StartVM(ctx context.Context, uuid uuid.UUID) error {
	var t = time.NewTicker(1 * time.Second)
	defer t.Stop()

	for range t.C {
		var info, err = c.GetVMInfo(ctx, uuid)
		if err != nil {
			return err
		}

		if info.State == VMState_Running {
			// fmt.Printf("VM is running: %d\n", info.State)
			break
		}

		err = c.l.DomainCreate(libvirt.Domain{
			UUID: libvirt.UUID(uuid),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
