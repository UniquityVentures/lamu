package p_filesystem

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const AppUrl = "/filesystem/"

// GetPlugin returns the registry contributions for this plugin for [lamu.BuildAllRegistries].
// Callers assembling the full plugin list should include a pair with key "p_filesystem" and this value.
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_filesystem", Value: lamu.Plugin{
			Type:        lamu.PluginTypeApp,
			Icon:        "folder",
			URL:         u,
			VerboseName: "Filesystem",
			Roles:       []string{"superuser", "admin"},
			Migrations:  pluginMigrations,
			Views:       pluginViews,
			Pages:       pluginPages,
			Routes:      pluginRoutes,
			Configs:     pluginConfigs,
			Generators:  pluginGenerators,
		},
	}
}
