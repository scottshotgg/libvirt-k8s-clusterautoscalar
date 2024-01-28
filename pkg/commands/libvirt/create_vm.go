package libvirt

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
)

func (c *Libvirt) CreateVM(ctx context.Context, groupID int, templatePath, name string) (*commands.VMInfo, error) {
	// TODO: this needs to happen on initialization
	var b, err = os.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	// TODO: just use filepath.Join here but im lazy right now

	absTemplatePath, err := filepath.Abs(templatePath)
	if err != nil {
		return nil, err
	}

	var (
		now   = time.Now()
		uuid  = uuid.New()
		id    = int(now.Unix())
		vmDir = fmt.Sprintf(baseVMsPath+"%s/", uuid)
		info  = commands.VMInfo{
			ID:         id,
			Name:       name,
			UUID:       uuid,
			DiskPath:   vmDir + diskName,
			MacAddress: macAddress(now),
		}
	)

	absVMDir, err := filepath.Abs(vmDir)
	if err != nil {
		return nil, err
	}

	var vmXML = fmt.Sprintf(
		string(b),
		info.Name,
		info.UUID,
		groupID,
		absTemplatePath,
		absVMDir,
		info.MacAddress,
		info.DiskPath,
		info.MacAddress,
	)

	err = os.MkdirAll(vmDir, 0777)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(baseImagePath+baseImageName, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	f2, err := os.OpenFile(info.DiskPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		// TODO: actually do DeleteVM
		var err2 = os.Remove(info.DiskPath)
		if err != nil {
			return nil, err2
		}

		return nil, err
	}

	defer f2.Close()

	// We could technically pre-allocate an HDD in the background
	// so that we don't need to do the copy here
	n, err := io.Copy(f2, f)
	if err != nil {
		// TODO: actually do DeleteVM
		var err2 = os.Remove(info.DiskPath)
		if err != nil {
			return nil, err2
		}

		return nil, err
	}

	fmt.Printf("Wrote %d bytes to %s!\n", n, info.DiskPath)

	domain, err := c.l.DomainDefineXML(vmXML)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(filepath.Join(vmDir, "config.xml"), []byte(vmXML), 0777)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(filepath.Join(vmDir, "template.xml"), b, 0777)
	if err != nil {
		return nil, err
	}

	// TODO: make API call to VyOS to register this VM to a certain IP

	// TODO: do we need this?
	_ = domain

	// TODO: just run this for now to edit the hostname
	// If we need to do more later then we can split out the 3 different commands into various functions here
	err = exec.
		Command("./set_hostname.sh", info.DiskPath, filepath.Join("/mnt/vm1", info.UUID.String()), info.Name).
		Run()

	if err != nil {
		return nil, err
	}

	return &info, nil
}
