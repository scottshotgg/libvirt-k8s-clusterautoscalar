package commands

type (
	Commands interface {
		ListVMs() ([]*VMInfo, error)
		GetVMByID(id int) (*VMInfo, error)
		CreateVM() (*VMInfo, error)
		DeleteVMByID(id int) (*VMInfo, error)
	}

	VMInfo struct {
		ID         int
		Name       string
		UUID       string
		DiskPath   string
		MacAddress string
	}
)
