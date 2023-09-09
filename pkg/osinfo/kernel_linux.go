//go:build linux && !windows

package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/zcalusic/sysinfo"
)

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
