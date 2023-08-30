package osinfo

import (
	"runtime"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/mem"
	"github.com/zcalusic/sysinfo"
)

type Memory struct {
	Type  string  `json:"type,omitempty"`
	Speed uint    `json:"speed,omitempty"` // RAM data rate in MT/s
	Total float64 `json:"total,omitempty"` // Total RAM size in GiB
	Used  float64 `json:"used,omitempty"`  // Used RAM size in GiB
}

func (h *Hardware) getMemoryInfo() {
	memInfo, err := mem.VirtualMemory()
	utils.Check(err)
	h.Memory.Used = utils.Byte2GiB(memInfo.Used)
	h.Memory.Total = utils.Byte2GiB(memInfo.Total)

	if runtime.GOOS == "linux" {
		var si sysinfo.SysInfo
		si.GetSysInfo()

		h.Memory.Type = si.Memory.Type
		h.Memory.Speed = si.Memory.Speed
	}
}
