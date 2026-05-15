package p_export

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const AppUrl = "/export/"

// GetPlugin returns the registry contributions for this plugin (views, pages, routes) for
// [lamu.BuildAllRegistries]. Callers that assemble the full plugin list should include
// a pair with key "p_export" and this value.
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_export", Value: lamu.Plugin{
			Type:        lamu.PluginTypeApp,
			Icon:        "arrow-down-tray",
			URL:         u,
			VerboseName: "Export",
			Views:       pluginStages(pluginViews),
			Pages:       pluginStages(pluginPages),
			Routes:      pluginStages(pluginRoutes),
		},
	}
}
