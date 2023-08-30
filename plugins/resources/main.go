package main

import (
	"runtime"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/osinfo"
)

type plug string

// exported plugin Symbol
var ResourcesConnect plug

// return instant resource usage information
func (p plug) ResourceUsage() map[string]interface{} {
	var system osinfo.System
	system.GetSystemInfo()

	data := map[string]interface{}{
		// CPU Information
		"CPU Physical Core Count": system.CPU.Cores,
		// "CPU Logical Core Count":  system.CPU.Logical_core_count, // TODO
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

	// TODO see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/resource_info_fetcher.py
	return data
}

// return general Agent system information
//
// these values changes rarely, see ResourceUsage() function for instant resource usage information
func (p plug) AgentInfo() map[string]interface{} {
	var system osinfo.System
	system.GetSystemInfo()

	// Common data
	data := map[string]interface{}{
		"System":         runtime.GOOS,
		"DiskSpaceTotal": system.Disk.Total,
		"MemoryTotal":    system.Memory.Total,
		// TODO "AhenkVersion": get Ahenk self version here

		"Name":          system.OS.Name,
		"Distribution":  system.OS.Distro,
		"Arch":          system.OS.Arch,
		"Version":       system.OS.Version,
		"Hostname":      system.Node.Hostname,
		"KernelVersion": system.Kernel.Version,
		"KernelRelease": system.Kernel.Release,
	}

	// REVIEW is calling all functions one by one slow downs code?
	// TODO see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/agent_info.py
	return data
}
