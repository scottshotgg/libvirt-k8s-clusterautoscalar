package main

import (
	"fmt"
	"log"

	libvirt_commands "github.com/scottshotgg/libvirt-test/pkg/commands/libvirt"
)

func main() {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		log.Fatal("libvirt_commands.New err:", err)
	}

	info, err := lvc.CreateVM()
	if err != nil {
		log.Fatal("CreateVM err:", err)
	}

	fmt.Println("info:", info)
}
