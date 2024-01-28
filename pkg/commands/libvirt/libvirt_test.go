package libvirt_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	libvirt_commands "github.com/scottshotgg/libvirt-test/pkg/commands/libvirt"
)

const (
	testTemplatePath = "./template.xml"
	testUUIDStr      = "0f021ddf-833c-42e9-9e41-bac209281816"
	testGroupID      = "1337"
)

var (
	testUUID = uuid.MustParse(testUUIDStr)
)

func TestStartVM(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	var ctx = context.Background()

	infos, err := lvc.ListVMs(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(infos) == 0 {
		t.Fatal("No VMs to start")
	}

	var found bool
	for _, v := range infos {
		err = lvc.StartVM(ctx, v.UUID)
		if err != nil {
			t.Fatal(err)
		}

		found = true
		break
	}

	if !found {
		t.Fatal("Could not find a VM to start")
	}
}

func TestStopVM(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	var ctx = context.Background()

	infos, err := lvc.ListVMs(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(infos) == 0 {
		t.Fatal("No VMs to start")
	}

	var found bool
	for _, v := range infos {
		err = lvc.StopVM(ctx, v.UUID)
		if err != nil {
			t.Fatal(err)
		}

		found = true
		break
	}

	if !found {
		t.Fatal("Could not find a VM to stop")
	}
}

func TestCreateVM(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	info, err := lvc.CreateVM(context.Background(), testGroupID, testTemplatePath, fmt.Sprintf("TEST-%d", time.Now().Unix()))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("info:", info)
}

func TestListVMs(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	infos, err := lvc.ListVMs(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range infos {
		fmt.Printf("k, v: %d %+v\n", k, v)
	}
}

func TestGetVMInfo(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	var ctx = context.Background()

	infos, err := lvc.ListVMs(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(infos) == 0 {
		t.Fatal("No VMs to start")
	}

	for _, v := range infos {
		info, err := lvc.GetVMInfo(ctx, v.UUID)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("info: %+v\n", info)
		break
	}
}

func TestDeleteVM(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	var ctx = context.Background()

	infos, err := lvc.ListVMs(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(infos) == 0 {
		t.Fatal("No VMs to start")
	}

	for _, v := range infos {
		err = lvc.DeleteVM(ctx, v.UUID)
		if err != nil {
			t.Fatal(err)
		}

		break
	}
}

func TestDeleteAllVMs(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	vms, err := lvc.ListVMs(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, vm := range vms {
		err = lvc.DeleteVM(context.Background(), vm.UUID)
		if err != nil {
			if err != libvirt_commands.ErrNotManaged {
				t.Fatal(err)
			}
		}
	}
}

func TestScale(t *testing.T) {
	var lvc, err = libvirt_commands.New()
	if err != nil {
		t.Fatal(err)
	}

	info, err := lvc.Scale(context.Background(), testGroupID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Info: %+v\n", info)
}
