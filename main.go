package main

import (
	"context"
	"fmt"
	"log"
	"time"

	libvirt_commands "github.com/scottshotgg/libvirt-test/pkg/commands/libvirt"
)

func main() {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		log.Fatal("libvirt_commands.New err:", err)
	}

	info, err := lvc.CreateVM(context.Background(), 1337, "./test.xml", fmt.Sprintf("TEST-%d", time.Now().Unix()))
	if err != nil {
		log.Fatal("CreateVM err:", err)
	}

	fmt.Println("info:", info)
}
