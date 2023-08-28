package main

import (
	"log"

	"github.com/zcalusic/sysinfo"
)

type plug string

// exported plugin Symbol
var TmptestConnect plug

func (p plug) TmpTest() {
	var si sysinfo.SysInfo

	si.GetSysInfo()

	log.Println(si)
}

func (p plug) Info() map[string]string {
	inf := make(map[string]string)
	inf["name"] = "tmptest"
	inf["version"] = "0.0.1"
	inf["support"] = "debian"
	inf["description"] = "Temporary testing"
	inf["developer"] = "asandikci@aliberksandikci.com.tr"
	return inf
}
