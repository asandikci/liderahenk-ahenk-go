package osinfo

import (
	"syscall"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const KB = uint64(1024)

// return kernel information as string map with calling syscall (like uname command)
func GetKernelInfo() map[string]string {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err != nil {
		panic(err)
	}
	return map[string]string{
		"Sysname":    utils.Byte2String(uname.Sysname[:]),
		"Nodename":   utils.Byte2String(uname.Nodename[:]),
		"Release":    utils.Byte2String(uname.Release[:]),
		"Version":    utils.Byte2String(uname.Version[:]),
		"Machine":    utils.Byte2String(uname.Machine[:]),
		"Domainname": utils.Byte2String(uname.Domainname[:]),
	}
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
		"total": utils.Byte2GiB(totalSize),
		"free":  utils.Byte2GiB(freeSize),
		"used":  utils.Byte2GiB(usedSize),
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
