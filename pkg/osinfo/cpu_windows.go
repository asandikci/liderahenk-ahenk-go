//go:build windows

package osinfo

// TODO Get windows cpu information
// FIXME Causes error on windows(wine)
func (h *System) getCPUInfo() {
	// var cpuInfo []cpu.InfoStat
	// var err error
	// cpuInfo, err = cpu.Info()
	// utils.Check(err)
	// h.CPU.Vendor = cpuInfo[0].VendorID
	// h.CPU.Vendor = cpuInfo[0].Model
}
