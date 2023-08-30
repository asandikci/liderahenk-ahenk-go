package osinfo

type OS struct {
	Name    string `json:"name,omitempty"`
	Distro  string `json:"distro,omitempty"`
	Version string `json:"version,omitempty"`
	Release string `json:"release,omitempty"`
	Arch    string `json:"architecture,omitempty"`
}
