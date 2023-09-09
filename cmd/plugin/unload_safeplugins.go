//go:build safeplugins

package plugin

import "log"

func unloadPlugins(params []string) {
	log.Println("You can not unload plugins in safeplugin build/logic, skipping...")
}
