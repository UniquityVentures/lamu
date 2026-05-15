package p_otp

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const AppURL = "/otp/preferences/"

// GetPlugin returns registry contributions for [lamu.BuildAllRegistries].
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppURL)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_otp",
		Value: lamu.Plugin{
			Type:        lamu.PluginTypeApp,
			Icon:        "key",
			URL:         u,
			VerboseName: "OTP Preferences",
			Roles:       []string{""},
			Migrations:  pluginStages(pluginMigrations),
			Views:       pluginStages(pluginViews),
			Pages:       pluginStages(pluginPages),
			Routes:      pluginStages(pluginRoutes),
			Models:      pluginStages(pluginModels),
		},
	}
}
