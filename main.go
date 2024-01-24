package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/google/uuid"
)

func main() {
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	domains, err := l.Domains()
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(9)
	}

	for _, d := range domains {
		fmt.Printf("%+v", d)
	}

	b, err := os.ReadFile("test.xml")
	if err != nil {
		fmt.Println("os.ReadFile err:", err)
		os.Exit(9)
	}

	var now = time.Now().Unix()
	var name = fmt.Sprintf("TEST-%d", now)
	var uid = uuid.New().String()

	domain, err := l.DomainDefineXML(fmt.Sprintf(string(b), now, name, uid, now, name))
	if err != nil {
		fmt.Println("l.DomainDefineXML err:", err)
		os.Exit(9)
	}

	fmt.Println(domain)

	// err = l.DomainCreate(libvirt.Domain{
	// 	Name: "something",
	// 	UUID: libvirt.UUID(uuid.New()),
	// 	ID:   -1,
	// })

	// if err != nil {
	// 	fmt.Println("woah an error:", err)
	// }
}
