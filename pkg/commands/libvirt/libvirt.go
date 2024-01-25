package impl

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
	"github.com/scottshotgg/libvirt-test/pkg/commands"
)

// This package uses libvirt - maybe rename the pkg to that

type (
	Libvirt struct {
		l *libvirt.Libvirt
	}
)

const (
	unixProtocol     = "unix"
	socketPath       = "/var/run/libvirt/libvirt-sock"
	dialTimeout      = 2 * time.Second
	baseVMsPath      = "/home/scottshotgg/VMs/"
	baseImagePath    = baseVMsPath + "base/"
	baseImageName    = "k8s-worker-arch.qcow2"
	newImageTmplName = "%d-disk-0.qcow2"
)

func New() (*Libvirt, error) {
	var c, err = net.DialTimeout(unixProtocol, socketPath, dialTimeout)
	if err != nil {
		return nil, err
	}

	var l = libvirt.New(c)
	err = l.Connect()
	if err != nil {
		return nil, err
	}

	return &Libvirt{
		l: l,
	}, nil
}

func (c *Libvirt) ListVMs() ([]*commands.VMInfo, error) {
	// domains, err := l.Domains()
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	os.Exit(9)
	// }

	// for _, d := range domains {
	// 	fmt.Printf("%+v\n", d)
	// }

	return nil, errors.New("not implemented")
}

func (c *Libvirt) GetVMByID(id int) (*commands.VMInfo, error) {
	return nil, errors.New("not implemented")
}

func (c *Libvirt) CreateVM() (*commands.VMInfo, error) {
	// TODO: this needs to happen on initialization
	var b, err = os.ReadFile("test.xml")
	if err != nil {
		return nil, err
	}

	// - copy template image to new folder
	//	 - /mnt/smb/proxmox/images/109 -> /mnt/smb/proxmox/images/<name>
	// - don't autostart k8s
	// - hostname check-in

	var (
		now   = time.Now()
		id    = int(now.Unix())
		vmDir = fmt.Sprintf(baseVMsPath+"%d/", id)
		info  = commands.VMInfo{
			ID:         id,
			Name:       fmt.Sprintf("TEST-%d", now.Unix()),
			UUID:       uuid.New().String(),
			DiskPath:   vmDir + fmt.Sprintf(newImageTmplName, id),
			MacAddress: macAddress(now),
		}

		x = fmt.Sprintf(
			string(b),
			info.ID,
			info.Name,
			info.UUID,
			info.DiskPath,
			info.MacAddress,
		)
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
		return nil, err
	}

	defer f2.Close()

	n, err := io.Copy(f2, f)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Wrote %d bytes to %s!\n", n, info.DiskPath)

	domain, err := c.l.DomainDefineXML(x)
	if err != nil {
		return nil, err
	}

	// TODO: do we need this?
	_ = domain

	return &info, nil
}

func (c *Libvirt) DeleteVMByID(id int) (*commands.VMInfo, error) {
	return nil, errors.New("not implemented")
}

// Maybe make a random one later
func macAddress(now time.Time) string {
	var s = fmt.Sprintf("%x", now.UnixMicro())
	return strings.Join([]string{s[1:2] + "2", s[2:4], s[4:6], s[6:8], s[8:10], s[10:12]}, ":")
}
