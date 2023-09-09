//go:build linux && !windows

package osinfo

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zcalusic/sysinfo"
)

func (h *System) getMemoryInfo() { // TODO also implement swap usage
	var si sysinfo.SysInfo
	si.GetSysInfo()

	memInfo, err := mem.VirtualMemory()
	utils.Check(err)

	h.Memory.Type = si.Memory.Type
	h.Memory.Speed = si.Memory.Speed
	h.Memory.Used = utils.Byte2GiB(memInfo.Used)
	h.Memory.Total = utils.Byte2GiB(memInfo.Total)
}
