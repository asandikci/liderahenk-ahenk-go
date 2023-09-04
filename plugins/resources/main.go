package resources

import (
	"os"
	"runtime"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/confdir"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/osinfo"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
)

type plug string

// exported plugin Symbol
var ResourcesConnect plug

// return instant resource usage information
func (p plug) ResourceUsage() map[string]interface{} { // ANCHOR[id=ResourceUsage]
	var system osinfo.System
	system.GetSystemInfo()

	return map[string]interface{}{
		// CPU Information
		"CPU Physical Core Count": system.CPU.Cores,
		"CPU Logical Core Count":  system.CPU.Threads,
		// "CPU Actual Hz":           system.CPU.ActualHz, // TODO
		// "CPU Advertised Hz":       system.CPU.Hz_advertised, // TODO
		// "Processor":               system.CPU.Brand, // TODO

		// Memory Information
		"Total Memory": system.Memory.Total,
		"Usage":        system.Memory.Used,

		// Disk Information
		"Total Disk": system.Disk.Total,
		"Usage Disk": system.Disk.Used,
		"Device":     system.Disk.Devices,
	}
	// see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/resource_info_fetcher.py
}

// return general Agent system information
//
// these values changes rarely, see ResourceUsage() function for instant resource usage information
func (p plug) AgentInfo() map[string]interface{} { // ANCHOR[id=AgentInfo]
	var system osinfo.System
	system.GetSystemInfo()
	ver, err := os.ReadFile(confdir.Paths.Version)
	utils.Check(err)

	// Common data
	return map[string]interface{}{
		"System":        runtime.GOOS,
		"AhenkVersion":  string(ver),
		"Hostname":      system.Node.Hostname,
		"KernelVersion": system.Kernel.Version,
		"KernelRelease": system.Kernel.Release,
		"Name":          system.OS.Name,
		"Distribution":  system.OS.Distro,
		"Arch":          system.OS.Arch,
		"Version":       system.OS.Version,

		"DiskSpaceTotal": system.Disk.Total,
		"MemoryTotal":    system.Memory.Total,
	}
	// see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/agent_info.py
}
