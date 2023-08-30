package osinfo

import (
	"runtime"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/zcalusic/sysinfo"
)

type CPU struct {
	Vendor  string `json:"vendor,omitempty"`
	Model   string `json:"model,omitempty"`
	Speed   uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores   uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Threads uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
}

func (h *Hardware) getCPUInfo() {
	if runtime.GOOS == "linux" {
		var si sysinfo.SysInfo
		si.GetSysInfo()

		h.CPU.Vendor = si.CPU.Vendor
		h.CPU.Model = si.CPU.Model
		h.CPU.Speed = si.CPU.Speed
		h.CPU.Cache = si.CPU.Cache
		h.CPU.Cpus = si.CPU.Cpus
		h.CPU.Cores = si.CPU.Cores
		h.CPU.Threads = si.CPU.Threads
	} else {
		cpuInfo, err := cpu.Info()
		utils.Check(err)
		h.CPU.Vendor = cpuInfo[0].VendorID
		h.CPU.Vendor = cpuInfo[0].Model
		// h.CPU.Vendor = cpuInfo[0].Speed // TODO

	}
}
