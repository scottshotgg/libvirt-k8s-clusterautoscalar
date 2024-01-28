package libvirt

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/digitalocean/go-libvirt"
)

// This package uses libvirt - maybe rename the pkg to that

type (
	Libvirt struct {
		l *libvirt.Libvirt
	}
)

const (
	// TODO: need to be able to supply all of these values
	unixProtocol  = "unix"
	socketPath    = "/var/run/libvirt/libvirt-sock"
	dialTimeout   = 2 * time.Second
	baseVMsPath   = "/home/scottshotgg/VMs/"
	baseImagePath = baseVMsPath + "base/"
	baseImageName = "k8s-worker-arch-comp.qcow2"
	diskName      = "disk.qcow2"
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

// Maybe make a random one later
func macAddress(now time.Time) string {
	var s = fmt.Sprintf("%x", now.UnixMicro())
	return strings.Join([]string{s[1:2] + "2", s[2:4], s[4:6], s[6:8], s[8:10], s[10:12]}, ":")
}
