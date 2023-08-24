package main

import (
	"runtime"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/osinfo"
)

type plug string

// exported plugin Symbol
var ResourcesConnect plug

// return instant resource usage information
func ResourceUsage() map[string]string {
	data := map[string]string{

		// {'System': self.Os.name(), 'Release': self.Os.kernel_release(),
		// 	'Version': self.Os.distribution_version(), 'Machine': self.Os.architecture(),
		// 	'CPU Physical Core Count': self.Hardware.Cpu.physical_core_count(),
		// 	'Total Memory': self.Hardware.Memory.total(),
		// 	'Usage': self.Hardware.Memory.used(),
		// 	'Total Disc': self.Hardware.Disk.total(),
		// 	'Usage Disc': self.Hardware.Disk.used(),
		// 	'Processor': self.Hardware.Cpu.brand(),
		// 	'Device': device,
		// 	'CPU Logical Core Count': self.Hardware.Cpu.logical_core_count(),
		// 	'CPU Actual Hz': self.Hardware.Cpu.hz_actual(),
		// 	'CPU Advertised Hz': self.Hardware.Cpu.hz_advertised()
		// 	}
	}
	return data
}

// return general Agent information (that changes rarely)
func (p plug) AgentInfo() map[string]interface{} {
	data := map[string]interface{}{
		"System": runtime.GOOS, "Kernel": osinfo.GetKernelInfo()["Release"],
		"hostname":        osinfo.GetKernelInfo()["Hostname"],
		"osMachine":       osinfo.GetKernelInfo()["Machine"],
		"diskTotal":       osinfo.GetDiskUsage()["total"],
		"diskUsed":        osinfo.GetDiskUsage()["used"],
		"diskFree":        osinfo.GetDiskUsage()["free"],
		"diskUsage":       osinfo.GetDiskUsage()["used"] / osinfo.GetDiskUsage()["total"],
		"memoryTotal":     osinfo.GetMemoryUsage()["total"],
		"memoryFree":      osinfo.GetMemoryUsage()["free"],
		"memoryAvaliable": osinfo.GetMemoryUsage()["avaliable"],
		"memoryUsed":      osinfo.GetMemoryUsage()["used"],
		"memoryCached":    osinfo.GetMemoryUsage()["cached"],
		"memoryUsage":     (osinfo.GetMemoryUsage()["used"] + osinfo.GetMemoryUsage()["cached"]) / osinfo.GetMemoryUsage()["total"], //REVIEW just used/total ?

		// TODO is calling all functions one by one slow downs code?

		// TODO 'agentVersion': self.get_agent_version(),
		// TODO 'ipAddresses': str(self.Hardware.Network.ip_addresses()).replace('[', '').replace(']', ''),
		// TODO  get distrubition name also (pardus,arch,debian,windows10 for example ...)
		// TODO 'macAddresses': str(self.Hardware.Network.mac_addresses()).replace('[', '').replace(']', ''),
		// TODO 'hardware.systemDefinitions': self.Hardware.system_definitions(),
		// TODO 'hardware.monitors': self.Hardware.monitors(),
		// TODO 'hardware.screens': self.Hardware.screens(),
		// TODO 'hardware.usbDevices': self.Hardware.usb_devices(),
		// TODO 'hardware.printers': self.Hardware.printers(),
		// TODO 'Device': device,
	}
	return data
}

func Info() map[string]string {
	inf := make(map[string]string)
	inf["name"] = "resources"
	inf["version"] = "0.0.2"
	inf["support"] = "debian"
	inf["description"] = "Resource Usage Information and Controls"
	inf["developer"] = "asandikci@aliberksandikci.com.tr"
	return inf
}
