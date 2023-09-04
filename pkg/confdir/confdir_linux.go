//go:build linux && !windows

package confdir

import "os"

func getPaths() Path {
	var path Path
	path.Program = "/usr/bin/ahenk-go"
	path.UserConfig, _ = os.UserConfigDir()
	path.Data = "/etc/ahenk-go/"
	path.Plugins = path.Data + "plugins/"
	path.Logs = path.Data + "logs/"
	path.Pid = "/run/ahenk-go.pid"
	path.Version = path.Data + "version"
	return path
}
