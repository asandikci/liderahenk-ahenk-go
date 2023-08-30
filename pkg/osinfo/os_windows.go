//go:build windows

package osinfo

import (
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/pkg/utils"
	"golang.org/x/sys/windows/registry"
)

func (h *System) getOSInfo() {

	reg, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	utils.Check(err)
	defer reg.Close()

	pn, _, err := reg.GetStringValue("ProductName")
	utils.Check(err)

	h.OS.Name = pn
	// TODO add other h.OS values
}
