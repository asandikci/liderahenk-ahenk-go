//go:build linux && !windows

package osinfo

import "github.com/zcalusic/sysinfo"

func (h *System) getOSInfo() { //REVIEW review info in pardus
	var si sysinfo.SysInfo
	si.GetSysInfo()

	h.OS.Name = si.OS.Name
	h.OS.Distro = si.OS.Vendor
	h.OS.Version = si.OS.Version
	h.OS.Release = si.OS.Release
	h.OS.Arch = si.OS.Architecture
}
