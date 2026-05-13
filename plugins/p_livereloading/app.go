package p_livereloading

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

// GetPlugin returns routes and metadata for [lamu.BuildAllRegistries].
// Shell head snippet registration stays in init() (see pages.go).
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse("/")
	if err != nil {
		log.Panic(err)
	}
	return registry.Pair[string, lamu.Plugin]{
		Key: "p_livereloading",
		Value: lamu.Plugin{
			Type:        lamu.PluginTypeAddon,
			Icon:        "arrow-path",
			URL:         u,
			VerboseName: "Live reload",
			Routes:      pluginRoutes,
		},
	}
}
