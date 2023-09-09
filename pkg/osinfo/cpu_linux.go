//go:build linux && !windows

package osinfo

import (
	"github.com/zcalusic/sysinfo"
)

func (h *System) getCPUInfo() {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	h.CPU.Vendor = si.CPU.Vendor
	h.CPU.Model = si.CPU.Model
	h.CPU.Speed = si.CPU.Speed
	h.CPU.Cache = si.CPU.Cache
	h.CPU.Cpus = si.CPU.Cpus
	h.CPU.Cores = si.CPU.Cores
	h.CPU.Threads = si.CPU.Threads
}
