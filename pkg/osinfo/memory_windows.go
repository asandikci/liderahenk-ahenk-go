//go:build windows

package osinfo

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/v3/mem"
)

// TODO also implement cached usage
func (h *System) getMemoryInfo() {
	memInfo, err := mem.VirtualMemory()
	utils.Check(err)
	h.Memory.Used = utils.Byte2GiB(memInfo.Used)
	h.Memory.Total = utils.Byte2GiB(memInfo.Total)
}
