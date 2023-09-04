package confdir

type Path struct {
	Program    string `json:"Program"`
	UserConfig string `json:"User Config"`
	Data       string `json:"Data"`
	Plugins    string `json:"Plugins"`
	Logs       string `json:"Logs"`
	Pid        string `json:"pid"`
	Version    string `json:"version"`
}

var Paths Path

func GetPaths() {
	Paths = getPaths()
}
