//go:build safeplugins

// IMPORTED SAFE PLUGINS, TEST PURPOSE

package plugin

import (
	"fmt"
	"os"

	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/plugins/resources"
	"git.aliberksandikci.com.tr/Liderahenk/ahenk-go/plugins/tmptest"
)

func connectPlugins(params []string) {
	// INFO
	tmptest.TmptestConnect.Info()
	resources.ResourcesConnect.Info()

	// FUNCTIONS
	tmptest.TmptestConnect.TmpTest()
	LogPlugin("Agent Info - SAFE", resources.ResourcesConnect.AgentInfo(), true)
	LogPlugin("Resource Usage - SAFE", resources.ResourcesConnect.ResourceUsage(), true)

	fmt.Println("THERE WAS NO ERROR, RUN make local OR make local_wine NOW FOR CONTINUE TESTING WITH GOLADER PLUGIN LOGIC")
	os.Exit(0)
}
