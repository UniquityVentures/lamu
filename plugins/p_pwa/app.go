package p_pwa

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

// GetPlugin returns registry contributions for [lamu.BuildAllRegistries].
// Shell head registrations for the manifest link remain in init() (see views.go).
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse("/")
	if err != nil {
		log.Panic(err)
	}
	return registry.Pair[string, lamu.Plugin]{
		Key: "p_pwa",
		Value: lamu.Plugin{
			Type:        lamu.PluginTypeAddon,
			Icon:        "cpu-chip",
			URL:         u,
			VerboseName: "PWA",
			Configs: pluginConfigs,
			Views:   pluginViews,
			Routes:  pluginRoutes,
		},
	}
}
