package commands

import (
	"context"

	"github.com/google/uuid"
)

// TODO: create VM state typedef and constants

type (
	Commands interface {
		ListVMs(ctx context.Context) ([]*VMInfo, error)
		GetVM(ctx context.Context, uuid uuid.UUID) (*VMInfo, error)
		CreateVM(ctx context.Context, groupID, templatePath, name string) (*VMInfo, error)
		DeleteVM(ctx context.Context, uuid uuid.UUID) (*VMInfo, error)

		StartVM(ctx context.Context, uuid uuid.UUID) (*VMInfo, error)
		StopVM(ctx context.Context, uuid uuid.UUID) (*VMInfo, error)

		Scale(ctx context.Context, groupID string) (*VMInfo, error)
	}

	VMResources struct {
		MaxMem     int
		CurrentMem int
		CPUs       int
		CPUTime    int
	}

	VMMetadata struct {
		GroupID      string `xml:"node-group-id"`
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
