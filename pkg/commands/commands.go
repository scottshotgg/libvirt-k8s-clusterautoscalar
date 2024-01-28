package commands

import (
	"github.com/google/uuid"
)

// TODO: create VM state typedef and constants

type (
	Commands interface {
		ListVMs() ([]*VMInfo, error)
		GetVM(info *VMInfo) (*VMInfo, error)
		CreateVM() (*VMInfo, error)
		DeleteVM(info *VMInfo) (*VMInfo, error)

		StartVM(uuid uuid.UUID) (*VMInfo, error)
		StopVM(uuid uuid.UUID) (*VMInfo, error)

		Scale(groupID int) (*VMInfo, error)
	}

	VMResources struct {
		MaxMem     int
		CurrentMem int
		CPUs       int
		CPUTime    int
	}

	VMMetadata struct {
		NodeGroup    int    `xml:"node-group"`
		TemplatePath string `xml:"template-path"`
		RootDir      string `xml:"root-dir"`
		MacAddress   string `xml:"mac-address"`
	}

	VMInfo struct {
		ID         int
		Name       string
		UUID       uuid.UUID
		DiskPath   string
		MacAddress string
		State      int
		Resources  VMResources
		Metadata   *VMMetadata
	}
)
