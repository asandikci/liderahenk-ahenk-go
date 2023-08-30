package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/zcalusic/sysinfo"
)

type Kernel struct {
	Sysname string `json:"sysname,omitempty"`
	Release string `json:"release,omitempty"`
	Version string `json:"version,omitempty"`
	Arch    string `json:"arch,omitempty"`
}

func (h *System) getKernelInfo() {

	var si sysinfo.SysInfo
	si.GetSysInfo()

	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		panic(err)
	}
	h.Kernel.Sysname = utils.Byte2String(uname.Sysname[:])
	h.Kernel.Release = si.Kernel.Release
	h.Kernel.Version = si.Kernel.Version
	h.Kernel.Arch = si.Kernel.Architecture
}

// REVIEW Windows compatibility
