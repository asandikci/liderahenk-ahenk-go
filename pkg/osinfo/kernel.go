package osinfo

type Kernel struct {
	Sysname string `json:"sysname,omitempty"`
	Release string `json:"release,omitempty"`
	Version string `json:"version,omitempty"`
	Arch    string `json:"arch,omitempty"`
}
