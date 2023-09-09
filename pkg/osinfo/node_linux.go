//go:build linux && !windows

package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/zcalusic/sysinfo"
)

func (h *System) getNodeInfo() {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		panic(err)
	}

	h.Node.Domainname = utils.Byte2String(uname.Domainname[:])
	h.Node.Hostname = si.Node.Hostname
	h.Node.MachineID = si.Node.MachineID
	h.Node.Hypervisor = si.Node.Hypervisor
	h.Node.Timezone = si.Node.Timezone
}
