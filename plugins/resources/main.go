package main

import (
	"ahenk-go/pkg/osinfo"
	"fmt"
	"runtime"
)

type greeting string

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
func (g greeting) AgentInfo() map[string]interface{} {
	data := map[string]interface{}{
		"System": runtime.GOOS, "Release": osinfo.GetKernelInfo()["Release"],
		// TODO 'agentVersion': self.get_agent_version(),
		"hostname": osinfo.GetKernelInfo()["Hostname"],
		// TODO 'ipAddresses': str(self.Hardware.Network.ip_addresses()).replace('[', '').replace(']', ''),
		// "osName":     osinfo.GetKernelInfo()["Sysname"], // TODO is it necessary?
		// "osNodeName": osinfo.GetKernelInfo()["NodeName"],// TODO is it necessary?
		// "osVersion":  osinfo.GetKernelInfo()["Version"], // TODO is it necessary?
		"osMachine": osinfo.GetKernelInfo()["Machine"],
		// TODO get distrubition name also (pardus,arch,debian,windows10 for example ...)
		// TODO 'macAddresses': str(self.Hardware.Network.mac_addresses()).replace('[', '').replace(']', ''),
		// TODO 'hardware.systemDefinitions': self.Hardware.system_definitions(),
		// TODO 'hardware.monitors': self.Hardware.monitors(),
		// TODO 'hardware.screens': self.Hardware.screens(),
		// TODO 'hardware.usbDevices': self.Hardware.usb_devices(),
		// TODO 'hardware.printers': self.Hardware.printers(),
		"diskTotal": osinfo.GetDiskUsage()["total"],
		"diskUsed":  osinfo.GetDiskUsage()["used"],
		"diskFree":  osinfo.GetDiskUsage()["free"],
		// TODO "diskUsage": osinfo.GetDiskUsage()["Usage"],
		// TODO 'memory': self.Hardware.Memory.total(),
		// TODO 'Device': device,
	}
	return data
}

func Info() map[string]string {
	inf := make(map[string]string)
	inf["name"] = "resources"
	inf["version"] = "0.0.1"
	inf["support"] = "debian"
	inf["description"] = "Resource Usage Information and Controls"
	inf["developer"] = "asandikci@aliberksandikci.com.tr"

	// inf["task"] = "true"
	// inf["user_oriented"] = "false"
	// inf["machine_oriented"] = "false"
	return inf
}

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}
func (g greeting) Myvar() {
	fmt.Println("I am here")
}

// this is exported
var Greeter greeting
