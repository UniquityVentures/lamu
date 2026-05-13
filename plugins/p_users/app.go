package p_users

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const (
	AppUrl  = "/users"
	RoleUrl = "/users/roles"
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
			Migrations:       pluginMigrations,
			Views:            pluginViews,
			Pages:            pluginPages,
			Routes:           pluginRoutes,
			Models:           pluginModels,
			Generators:       pluginGenerators,
			DBInitHooks:      pluginDBInitHooks,
			Configs:          pluginAuthConfigs,
			CommandFactories: pluginCommandFactories,
		},
	}
}
