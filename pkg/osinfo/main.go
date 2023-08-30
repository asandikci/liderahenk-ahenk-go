package osinfo

const KB = uint64(1024)

type System struct {
	Kernel Kernel `json:"kernel"`
	OS     OS     `json:"OS"`
	Node   Node   `json:"node"`

	CPU    CPU    `json:"cpu"`
	Memory Memory `json:"memory"`
	Disk   Disk   `json:"disk"`
}

// initializes all hardware structs that gets software and hardware (system) information
func (h *System) GetSystemInfo() {
	h.getKernelInfo()
	h.getOSInfo()
	h.getNodeInfo()

	h.getCPUInfo()
	h.getMemoryInfo()
	h.getDiskInfo()
}

// TODO seperate file structure (as in gopsutil)
/*
main.go
disk/
--disk.go
--disk_linux.go
--disk_windows.go
cpu/
--cpu.go
--cpu_linux.go
--cpu_windows.go
memory/ ...
...
*/
