package osinfo

type Node struct {
	Domainname string `json:"domainname,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}
