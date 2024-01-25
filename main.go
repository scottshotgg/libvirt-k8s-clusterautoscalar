package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

type (
	VMInfo struct {
		ID         int
		Name       string
		UUID       string
		DiskPath   string
		MacAddress string
	}
)

/*
	- id (num)
	- name
	- uuid
	- disk path
	- mac address
*/

func main() {
	const (
		unixProtocol = "unix"
		socketPath   = "/var/run/libvirt/libvirt-sock"
		dialTimeout  = 2 * time.Second
	)

	c, err := net.DialTimeout(unixProtocol, socketPath, dialTimeout)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	var l = libvirt.New(c)
	err = l.Connect()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// domains, err := l.Domains()
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	os.Exit(9)
	// }

	// for _, d := range domains {
	// 	fmt.Printf("%+v\n", d)
	// }

	// TODO: this needs to happen on initialization
	b, err := os.ReadFile("test.xml")
	if err != nil {
		fmt.Println("os.ReadFile err:", err)
		os.Exit(9)
	}

	var (
		now  = time.Now()
		info = VMInfo{
			ID:   int(now.Unix()),
			Name: fmt.Sprintf("TEST-%d", now.Unix()),
			UUID: uuid.New().String(),
			// - copy template image to new folder
			//	 - /mnt/smb/proxmox/images/109 -> /mnt/smb/proxmox/images/<name>
			// - don't autostart k8s
			// - hostname check-in
			DiskPath:   "REPLACE_ME_DADDY",
			MacAddress: macAddress(now),
		}

		x = fmt.Sprintf(
			string(b),
			info.ID,
			info.Name,
			info.UUID,
			"",
			info.DiskPath,
			info.MacAddress,
		)
	)

	domain, err := l.DomainDefineXML(x)
	if err != nil {
		fmt.Println("l.DomainDefineXML err:", err)
		os.Exit(9)
	}

	_ = domain
}

// Maybe make a random one later
func macAddress(now time.Time) string {
	var s = fmt.Sprintf("%x", now.UnixMicro())
	return strings.Join([]string{s[1:2] + "2", s[2:4], s[4:6], s[6:8], s[8:10], s[10:12]}, ":")
}
