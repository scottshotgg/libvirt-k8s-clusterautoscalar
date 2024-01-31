package test

// import (
// 	"encoding/xml"
// )

// type (
// 	Domain struct {
// 		XMLName  xml.Name `xml:"domain"`
// 		Text     string   `xml:",chardata"`
// 		Type     string   `xml:"type,attr"`
// 		ID       string   `xml:"id,attr"`
// 		Name     string   `xml:"name"`
// 		Uuid     string   `xml:"uuid"`
// 		Metadata Metadata `xml:"metadata"`
// 		Memory   struct {
// 			Text string `xml:",chardata"`
// 			Unit string `xml:"unit,attr"`
// 		} `xml:"memory"`
// 		CurrentMemory struct {
// 			Text string `xml:",chardata"`
// 			Unit string `xml:"unit,attr"`
// 		} `xml:"currentMemory"`
// 		Vcpu struct {
// 			Text      string `xml:",chardata"`
// 			Placement string `xml:"placement,attr"`
// 		} `xml:"vcpu"`
// 		Resource struct {
// 			Text      string `xml:",chardata"`
// 			Partition string `xml:"partition"`
// 		} `xml:"resource"`
// 		Os struct {
// 			Text string `xml:",chardata"`
// 			Type struct {
// 				Text    string `xml:",chardata"`
// 				Arch    string `xml:"arch,attr"`
// 				Machine string `xml:"machine,attr"`
// 			} `xml:"type"`
// 			Boot struct {
// 				Text string `xml:",chardata"`
// 				Dev  string `xml:"dev,attr"`
// 			} `xml:"boot"`
// 		} `xml:"os"`
// 		Features struct {
// 			Text   string `xml:",chardata"`
// 			Acpi   string `xml:"acpi"`
// 			Apic   string `xml:"apic"`
// 			Vmport struct {
// 				Text  string `xml:",chardata"`
// 				State string `xml:"state,attr"`
// 			} `xml:"vmport"`
// 		} `xml:"features"`
// 		Cpu struct {
// 			Text       string `xml:",chardata"`
// 			Mode       string `xml:"mode,attr"`
// 			Check      string `xml:"check,attr"`
// 			Migratable string `xml:"migratable,attr"`
// 			Topology   struct {
// 				Text    string `xml:",chardata"`
// 				Sockets string `xml:"sockets,attr"`
// 				Dies    string `xml:"dies,attr"`
// 				Cores   string `xml:"cores,attr"`
// 				Threads string `xml:"threads,attr"`
// 			} `xml:"topology"`
// 		} `xml:"cpu"`
// 		Clock struct {
// 			Text   string `xml:",chardata"`
// 			Offset string `xml:"offset,attr"`
// 			Timer  []struct {
// 				Text       string `xml:",chardata"`
// 				Name       string `xml:"name,attr"`
// 				Tickpolicy string `xml:"tickpolicy,attr"`
// 				Present    string `xml:"present,attr"`
// 			} `xml:"timer"`
// 		} `xml:"clock"`
// 		OnPoweroff string `xml:"on_poweroff"`
// 		OnReboot   string `xml:"on_reboot"`
// 		OnCrash    string `xml:"on_crash"`
// 		Pm         struct {
// 			Text         string `xml:",chardata"`
// 			SuspendToMem struct {
// 				Text    string `xml:",chardata"`
// 				Enabled string `xml:"enabled,attr"`
// 			} `xml:"suspend-to-mem"`
// 			SuspendToDisk struct {
// 				Text    string `xml:",chardata"`
// 				Enabled string `xml:"enabled,attr"`
// 			} `xml:"suspend-to-disk"`
// 		} `xml:"pm"`
// 		Devices struct {
// 			Text     string `xml:",chardata"`
// 			Emulator string `xml:"emulator"`
// 			Disk     struct {
// 				Text   string `xml:",chardata"`
// 				Type   string `xml:"type,attr"`
// 				Device string `xml:"device,attr"`
// 				Driver struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 					Type string `xml:"type,attr"`
// 				} `xml:"driver"`
// 				Source struct {
// 					Text  string `xml:",chardata"`
// 					File  string `xml:"file,attr"`
// 					Index string `xml:"index,attr"`
// 				} `xml:"source"`
// 				BackingStore string `xml:"backingStore"`
// 				Target       struct {
// 					Text string `xml:",chardata"`
// 					Dev  string `xml:"dev,attr"`
// 					Bus  string `xml:"bus,attr"`
// 				} `xml:"target"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text     string `xml:",chardata"`
// 					Type     string `xml:"type,attr"`
// 					Domain   string `xml:"domain,attr"`
// 					Bus      string `xml:"bus,attr"`
// 					Slot     string `xml:"slot,attr"`
// 					Function string `xml:"function,attr"`
// 				} `xml:"address"`
// 			} `xml:"disk"`
// 			Controller []struct {
// 				Text      string `xml:",chardata"`
// 				Type      string `xml:"type,attr"`
// 				Index     string `xml:"index,attr"`
// 				AttrModel string `xml:"model,attr"`
// 				Ports     string `xml:"ports,attr"`
// 				Alias     struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text          string `xml:",chardata"`
// 					Type          string `xml:"type,attr"`
// 					Domain        string `xml:"domain,attr"`
// 					Bus           string `xml:"bus,attr"`
// 					Slot          string `xml:"slot,attr"`
// 					Function      string `xml:"function,attr"`
// 					Multifunction string `xml:"multifunction,attr"`
// 				} `xml:"address"`
// 				Model struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"model"`
// 				Target struct {
// 					Text    string `xml:",chardata"`
// 					Chassis string `xml:"chassis,attr"`
// 					Port    string `xml:"port,attr"`
// 				} `xml:"target"`
// 			} `xml:"controller"`
// 			Interface struct {
// 				Text string `xml:",chardata"`
// 				Type string `xml:"type,attr"`
// 				Mac  struct {
// 					Text    string `xml:",chardata"`
// 					Address string `xml:"address,attr"`
// 				} `xml:"mac"`
// 				Source struct {
// 					Text   string `xml:",chardata"`
// 					Bridge string `xml:"bridge,attr"`
// 				} `xml:"source"`
// 				Target struct {
// 					Text string `xml:",chardata"`
// 					Dev  string `xml:"dev,attr"`
// 				} `xml:"target"`
// 				Model struct {
// 					Text string `xml:",chardata"`
// 					Type string `xml:"type,attr"`
// 				} `xml:"model"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text     string `xml:",chardata"`
// 					Type     string `xml:"type,attr"`
// 					Domain   string `xml:"domain,attr"`
// 					Bus      string `xml:"bus,attr"`
// 					Slot     string `xml:"slot,attr"`
// 					Function string `xml:"function,attr"`
// 				} `xml:"address"`
// 			} `xml:"interface"`
// 			Serial struct {
// 				Text   string `xml:",chardata"`
// 				Type   string `xml:"type,attr"`
// 				Source struct {
// 					Text string `xml:",chardata"`
// 					Path string `xml:"path,attr"`
// 				} `xml:"source"`
// 				Target struct {
// 					Text  string `xml:",chardata"`
// 					Type  string `xml:"type,attr"`
// 					Port  string `xml:"port,attr"`
// 					Model struct {
// 						Text string `xml:",chardata"`
// 						Name string `xml:"name,attr"`
// 					} `xml:"model"`
// 				} `xml:"target"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 			} `xml:"serial"`
// 			Console struct {
// 				Text   string `xml:",chardata"`
// 				Type   string `xml:"type,attr"`
// 				Tty    string `xml:"tty,attr"`
// 				Source struct {
// 					Text string `xml:",chardata"`
// 					Path string `xml:"path,attr"`
// 				} `xml:"source"`
// 				Target struct {
// 					Text string `xml:",chardata"`
// 					Type string `xml:"type,attr"`
// 					Port string `xml:"port,attr"`
// 				} `xml:"target"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 			} `xml:"console"`
// 			Channel []struct {
// 				Text   string `xml:",chardata"`
// 				Type   string `xml:"type,attr"`
// 				Source struct {
// 					Text string `xml:",chardata"`
// 					Mode string `xml:"mode,attr"`
// 					Path string `xml:"path,attr"`
// 				} `xml:"source"`
// 				Target struct {
// 					Text  string `xml:",chardata"`
// 					Type  string `xml:"type,attr"`
// 					Name  string `xml:"name,attr"`
// 					State string `xml:"state,attr"`
// 				} `xml:"target"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text       string `xml:",chardata"`
// 					Type       string `xml:"type,attr"`
// 					Controller string `xml:"controller,attr"`
// 					Bus        string `xml:"bus,attr"`
// 					Port       string `xml:"port,attr"`
// 				} `xml:"address"`
// 			} `xml:"channel"`
// 			Input []struct {
// 				Text  string `xml:",chardata"`
// 				Type  string `xml:"type,attr"`
// 				Bus   string `xml:"bus,attr"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text string `xml:",chardata"`
// 					Type string `xml:"type,attr"`
// 					Bus  string `xml:"bus,attr"`
// 					Port string `xml:"port,attr"`
// 				} `xml:"address"`
// 			} `xml:"input"`
// 			Graphics []struct {
// 				Text       string `xml:",chardata"`
// 				Type       string `xml:"type,attr"`
// 				Port       string `xml:"port,attr"`
// 				Autoport   string `xml:"autoport,attr"`
// 				AttrListen string `xml:"listen,attr"`
// 				Listen     struct {
// 					Text    string `xml:",chardata"`
// 					Type    string `xml:"type,attr"`
// 					Address string `xml:"address,attr"`
// 				} `xml:"listen"`
// 				Image struct {
// 					Text        string `xml:",chardata"`
// 					Compression string `xml:"compression,attr"`
// 				} `xml:"image"`
// 			} `xml:"graphics"`
// 			Sound struct {
// 				Text  string `xml:",chardata"`
// 				Model string `xml:"model,attr"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text     string `xml:",chardata"`
// 					Type     string `xml:"type,attr"`
// 					Domain   string `xml:"domain,attr"`
// 					Bus      string `xml:"bus,attr"`
// 					Slot     string `xml:"slot,attr"`
// 					Function string `xml:"function,attr"`
// 				} `xml:"address"`
// 			} `xml:"sound"`
// 			Audio struct {
// 				Text string `xml:",chardata"`
// 				ID   string `xml:"id,attr"`
// 				Type string `xml:"type,attr"`
// 			} `xml:"audio"`
// 			Video struct {
// 				Text  string `xml:",chardata"`
// 				Model struct {
// 					Text    string `xml:",chardata"`
// 					Type    string `xml:"type,attr"`
// 					Heads   string `xml:"heads,attr"`
// 					Primary string `xml:"primary,attr"`
// 				} `xml:"model"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text     string `xml:",chardata"`
// 					Type     string `xml:"type,attr"`
// 					Domain   string `xml:"domain,attr"`
// 					Bus      string `xml:"bus,attr"`
// 					Slot     string `xml:"slot,attr"`
// 					Function string `xml:"function,attr"`
// 				} `xml:"address"`
// 			} `xml:"video"`
// 			Redirdev []struct {
// 				Text  string `xml:",chardata"`
// 				Bus   string `xml:"bus,attr"`
// 				Type  string `xml:"type,attr"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text string `xml:",chardata"`
// 					Type string `xml:"type,attr"`
// 					Bus  string `xml:"bus,attr"`
// 					Port string `xml:"port,attr"`
// 				} `xml:"address"`
// 			} `xml:"redirdev"`
// 			Watchdog struct {
// 				Text   string `xml:",chardata"`
// 				Model  string `xml:"model,attr"`
// 				Action string `xml:"action,attr"`
// 				Alias  struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 			} `xml:"watchdog"`
// 			Memballoon struct {
// 				Text  string `xml:",chardata"`
// 				Model string `xml:"model,attr"`
// 				Alias struct {
// 					Text string `xml:",chardata"`
// 					Name string `xml:"name,attr"`
// 				} `xml:"alias"`
// 				Address struct {
// 					Text     string `xml:",chardata"`
// 					Type     string `xml:"type,attr"`
// 					Domain   string `xml:"domain,attr"`
// 					Bus      string `xml:"bus,attr"`
// 					Slot     string `xml:"slot,attr"`
// 					Function string `xml:"function,attr"`
// 				} `xml:"address"`
// 			} `xml:"memballoon"`
// 			Rng Rng `xml:"rng"`
// 		} `xml:"devices"`
// 		Seclabel SecLabel `xml:"seclabel"`
// 	}

// 	SecLabel struct {
// 		Text       string `xml:",chardata"`
// 		Type       string `xml:"type,attr"`
// 		Model      string `xml:"model,attr"`
// 		Relabel    string `xml:"relabel,attr"`
// 		Label      string `xml:"label"`
// 		Imagelabel string `xml:"imagelabel"`
// 	}

// 	Rng struct {
// 		Text    string  `xml:",chardata"`
// 		Model   string  `xml:"model,attr"`
// 		Backend Backend `xml:"backend"`
// 		Alias   Alias   `xml:"alias"`
// 		Address Address `xml:"address"`
// 	}

// 	Backend struct {
// 		Text  string `xml:",chardata"`
// 		Model string `xml:"model,attr"`
// 	}

// 	Alias struct {
// 		Text string `xml:",chardata"`
// 		Name string `xml:"name,attr"`
// 	}

// 	Address struct {
// 		Text     string `xml:",chardata"`
// 		Type     string `xml:"type,attr"`
// 		Domain   string `xml:"domain,attr"`
// 		Bus      string `xml:"bus,attr"`
// 		Slot     string `xml:"slot,attr"`
// 		Function string `xml:"function,attr"`
// 	}

// 	Metadata struct {
// 		Text            string       `xml:",chardata"`
// 		AttrLibosinfo   string       `xml:"libosinfo,attr"`
// 		CockpitMachines string       `xml:"cockpit_machines,attr"`
// 		Libosinfo       Libosinfo    `xml:"libosinfo"`
// 		Data            MetadataData `xml:"data"`
// 	}

// 	Libosinfo struct {
// 		Text string `xml:",chardata"`
// 		OS   OS     `xml:"os"`
// 	}

// 	OS struct {
// 		Text string `xml:",chardata"`
// 		ID   string `xml:"id,attr"`
// 	}

// 	MetadataData struct {
// 		Text              string `xml:",chardata"`
// 		HasInstallPhase   string `xml:"has_install_phase"`
// 		InstallSourceType string `xml:"install_source_type"`
// 		InstallSource     string `xml:"install_source"`
// 		OsVariant         string `xml:"os_variant"`
// 	}
// )
