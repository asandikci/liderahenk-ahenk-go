package osinfo

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Devices string  `json:"devices,omitempty"` // physical disk devices as string
	Total   float64 `json:"total,omitempty"`   // Total Disk size in GiB
	Used    float64 `json:"used,omitempty"`    // Used Disk size in GiB
	Free    float64 `json:"free,omitempty"`    // Free Disk size in GiB
}

func (h *Hardware) getDiskInfo() {

	for _, v := range GetDisks() {
		h.Disk.Devices += v.Device
	}
	h.Disk.Total = GetDiskUsage()["spaceTotal"]
	h.Disk.Used = GetDiskUsage()["usedTotal"]
	h.Disk.Free = GetDiskUsage()["usedFree"]
}

// return all physical disks, USB and CD-ROM devices
//
// An example output:
// device: /dev/example, mountpoint: /, fstype: ext4, opts:rw,noatime
func GetDisks() []disk.PartitionStat {
	parts, err := disk.Partitions(false)
	utils.Check(err)
	return parts
}

// return disk usage as GiB
//
// TODO different functions for all disks / a specified disk?
// FIXME Wrong Disk values for docker !!! (probably because counting also virtual mountpoints?)
func GetDiskUsage() map[string]float64 {
	var totalSize, freeSize, usedSize uint64
	for _, part := range GetDisks() {
		u, err := disk.Usage(part.Mountpoint)
		utils.Check(err)
		totalSize += u.Total
		freeSize += u.Free
		usedSize += u.Used
	}
	return map[string]float64{
		"spaceTotal": utils.Byte2GiB(totalSize),
		"freeTotal":  utils.Byte2GiB(freeSize),
		"usedTotal":  utils.Byte2GiB(usedSize),
	}
}
