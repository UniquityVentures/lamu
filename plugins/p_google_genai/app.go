package p_google_genai

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const AppUrl = "/google-genai/"

// GetPlugin returns registry contributions for [lamu.BuildAllRegistries].
func GetPlugin() registry.Pair[string, lamu.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lamu.Plugin]{
		Key: "p_google_genai",
		Value: lamu.Plugin{
			// Addon: not listed on dashboard Apps grid; API key consumed by other plugins.
			Type:        lamu.PluginTypeAddon,
			Icon:        "sparkles",
			URL:         u,
			VerboseName: "Google GenAI",
			Configs:     pluginStages(pluginConfigs),
		},
	}
}
