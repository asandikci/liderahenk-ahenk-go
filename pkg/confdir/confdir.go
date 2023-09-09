package confdir

type Path struct {
	Program    string `json:"Program"`     // executable path
	UserConfig string `json:"User Config"` // user config directory path
	Data       string `json:"Data"`        // program data directory path
	Plugins    string `json:"Plugins"`     // plugin directory path
	Logs       string `json:"Logs"`        // log directory path
	Pid        string `json:"pid"`         // linux pid file path
	Version    string `json:"version"`     // version file path
}

var Paths Path

func GetPaths() {
	Paths = getPaths()
}
