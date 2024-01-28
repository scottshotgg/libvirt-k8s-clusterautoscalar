package libvirt

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func (c *Libvirt) StopVM(ctx context.Context, uuid uuid.UUID) error {
	var (
		ticker = time.NewTicker(1 * time.Second)
		// timer  = time.NewTimer(30 * time.Second)
	)

	defer ticker.Stop()
	// defer timer.Stop()

	for range ticker.C {
		// select {
		// case <-ctx.Done():
		// 	return ctx.Err()

		// case <-ticker.C:

		// 	// TODO: need to write a force stop in here
		// 	// case <-timer.C:

		// }

		var info, err = c.GetVMInfo(ctx, uuid)
		if err != nil {
			return err
		}

		if info.State != 1 {
			// fmt.Printf("VM is not running: %d\n", info.State)
			break
		}

		err = c.l.DomainShutdown(libvirt.Domain{
			UUID: libvirt.UUID(uuid),
		})

		if err != nil {
			fmt.Println("err:", err)
			return err
		}
	}

	return nil
}
