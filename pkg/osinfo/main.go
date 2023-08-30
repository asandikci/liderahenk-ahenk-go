package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/mem"
	"github.com/zcalusic/sysinfo"
)

const KB = uint64(1024)

type Hardware struct {
	CPU    CPU    `json:"cpu"`
	Memory Memory `json:"memory"`
	Disk   Disk   `json:"disk"`
	// Node   Node   `json:"node"`
	// Kernel Kernel `json:"kernel"`
	OS OS `json:"OS"`
}

func (h *Hardware) GetHardwareInfo() {
	h.getCPUInfo()
	h.getMemoryInfo()
	h.getDiskInfo()
	h.getOSInfo()
}

// return linux Kernel, Node and System Information
//
// REVIEW are there any command that also compatible with windows?
func GetLinuxInfo() map[string]map[string]interface{} {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		panic(err)
	}

	var si sysinfo.SysInfo
	si.GetSysInfo()
	return map[string]map[string]interface{}{
		"Kernel": {
			"Sysname": utils.Byte2String(uname.Sysname[:]),
			"Release": si.Kernel.Release,
			"Version": si.Kernel.Version,
			"Machine": si.Kernel.Architecture,
		},
		"Node": {
			"Domainname": utils.Byte2String(uname.Domainname[:]),
			"Hostname":   si.Node.Hostname,
			"MachineID":  si.Node.MachineID,
			"Timezone":   si.Node.Timezone,
		},
		"OS": { //REVIEW review info in pardus
			"Name":    si.OS.Name,
			"Vendor":  si.OS.Vendor,
			"Arch":    si.OS.Architecture,
			"Version": si.OS.Version,
			"Release": si.OS.Release,
		},
	}
}

// return memory usage as GiB
//
// TODO also implement swap usage
func GetMemoryUsage() map[string]float64 {
	v, _ := mem.VirtualMemory()
	return map[string]float64{
		"total":     utils.Byte2GiB(v.Total),
		"free":      utils.Byte2GiB(v.Free),
		"used":      utils.Byte2GiB(v.Used),
		"avaliable": utils.Byte2GiB(v.Available),
		"cached":    utils.Byte2GiB(v.Cached),
	}
}
