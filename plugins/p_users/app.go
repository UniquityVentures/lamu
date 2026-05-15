package p_users

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

// GetPlugin returns the registry contributions for this plugin for [lamu.BuildAllRegistries].
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_users",
		Value: lamu.Plugin{
			Type:             lamu.PluginTypeApp,
			Icon:             "users",
			URL:              u,
			VerboseName:      "Users",
			Migrations:       pluginStages(pluginMigrations),
			Views:            pluginStages(pluginViews),
			Pages:            pluginStages(pluginPages),
			Routes:           pluginStages(pluginRoutes),
			Models:           pluginStages(pluginModels),
			Generators:       pluginStages(pluginGenerators),
			DBInitHooks:      pluginStages(pluginDBInitHooks),
			Configs:          pluginStages(pluginAuthConfigs),
			CommandFactories: pluginStages(pluginCommandFactories),
		},
	}
}
