package p_google_genai

import (
	"strings"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

type Config struct {
	APIKey string `toml:"apiKey"`
}

var GoogleGenAIConfig = &Config{}

func (c *Config) PostConfig() {
	if c == nil {
		return
	}
	c.APIKey = strings.TrimSpace(c.APIKey)
}

func pluginConfigs() lamu.PluginFeatures[lamu.Config] {
	return lamu.PluginFeatures[lamu.Config]{
		Entries: []registry.Pair[string, lamu.Config]{
			{Key: "p_google_genai", Value: GoogleGenAIConfig},
		},
	}
}
