package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/zcalusic/sysinfo"
)

type Node struct {
	Domainname string `json:"domainname,omitempty"`
	Hostname   string `json:"hostname,omitempty"`
	MachineID  string `json:"machineid,omitempty"`
	Hypervisor string `json:"hypervisor,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
}

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

// REVIEW Windows compatibility
