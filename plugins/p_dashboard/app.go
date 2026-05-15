package p_dashboard

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const AppUrl = "/dashboard/"

// GetPlugin returns the registry contributions for this plugin (views, pages, routes) for
// [lamu.BuildAllRegistries]. Callers that assemble the full plugin list should include
// a pair with key "p_dashboard" and this value.
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_dashboard", Value: lamu.Plugin{
			Type:        lamu.PluginTypeAddon,
			Icon:        "dashboard",
			URL:         u,
			VerboseName: "Dashboard",
			Views:       pluginStages(pluginViews),
			Pages:       pluginStages(pluginPages),
			Routes:      pluginStages(pluginRoutes),
		},
	}
}
