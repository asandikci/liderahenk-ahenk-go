package osinfo

import (
	"runtime"

	"github.com/zcalusic/sysinfo"
)

type OS struct {
	Name    string `json:"name,omitempty"`
	Distro  string `json:"distro,omitempty"`
	Version string `json:"version,omitempty"`
	Release string `json:"release,omitempty"`
	Arch    string `json:"architecture,omitempty"`
}

func (h *Hardware) getOSInfo() {
	if runtime.GOOS == "linux" {
		var si sysinfo.SysInfo
		si.GetSysInfo()

		h.OS.Name = si.OS.Name
		h.OS.Distro = si.OS.Vendor
		h.OS.Version = si.OS.Version
		h.OS.Release = si.OS.Release
		h.OS.Arch = si.OS.Architecture
	}
}
