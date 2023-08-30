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
	var hardware osinfo.Hardware
	hardware.GetHardwareInfo()

	data := map[string]interface{}{
		// General System Information
		"System":  runtime.GOOS,
		"Release": osinfo.GetLinuxInfo()["Kernel"]["Release"], //needs REVIEW for windows
		"Version": osinfo.GetLinuxInfo()["OS"]["Version"],     //needs REVIEW for windows
		"Machine": osinfo.GetLinuxInfo()["OS"]["Arch"],        //needs REVIEW for windows

		// CPU Information
		"CPU Physical Core Count": hardware.CPU.Cores,
		// "CPU Logical Core Count":  hardware.CPU.Logical_core_count, // TODO
		// "CPU Actual Hz":           hardware.CPU.ActualHz, // TODO
		// "CPU Advertised Hz":       hardware.CPU.Hz_advertised, // TODO
		// "Processor":               hardware.CPU.Brand, // TODO

		// Memory Information
		"Total Memory": hardware.Memory.Total,
		"Usage":        hardware.Memory.Used,

		// Disk Information
		"Total Disk": hardware.Disk.Total,
		"Usage Disk": hardware.Disk.Used,
		"Device":     hardware.Disk.Devices,
	}

	// TODO see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/resource_info_fetcher.py
	return data
}

// return general Agent system information
//
// these values changes rarely, see ResourceUsage() function for instant resource usage information
func (p plug) AgentInfo() map[string]interface{} {
	var hardware osinfo.Hardware
	hardware.GetHardwareInfo()

	// Common data
	data := map[string]interface{}{
		"System":         runtime.GOOS,
		"DiskSpaceTotal": hardware.Disk.Total,
		"MemoryTotal":    hardware.Memory.Total,
		// TODO "AhenkVersion": get Ahenk self version here
	}

	// Linux specific data
	if runtime.GOOS == "linux" {
		data["Name"] = osinfo.GetLinuxInfo()["OS"]["Name"]
		data["Distribution"] = osinfo.GetLinuxInfo()["OS"]["Vendor"]
		data["Architecture"] = osinfo.GetLinuxInfo()["OS"]["Arch"]
		data["Version"] = osinfo.GetLinuxInfo()["OS"]["Version"]

		data["NodeHostname"] = osinfo.GetLinuxInfo()["Node"]["Hostname"]

		data["Architecture"] = osinfo.GetLinuxInfo()["Kernel"]["Machine"]
		data["KernelVersion"] = osinfo.GetLinuxInfo()["Kernel"]["Version"]
	}

	// LINK see https://github.com/golang/go/blob/master/src/go/build/syslist.go#L14 for all possible Operating systems
	// and https://go.dev/doc/install/source#environment for all possible combinations

	// REVIEW is calling all functions one by one slow downs code?

	// TODO see https://github.com/Pardus-LiderAhenk/ahenk/blob/master/src/plugins/resource-usage/agent_info.py
	return data
}

func (p plug) Info() map[string]string {
	inf := make(map[string]string)
	inf["name"] = "resources"
	inf["version"] = "0.0.2"
	inf["support"] = "debian"
	inf["description"] = "Resource Usage Information and Controls"
	inf["developer"] = "asandikci@aliberksandikci.com.tr"
	return inf
}
