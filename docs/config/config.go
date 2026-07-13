// Package config contains explanations and code examples for plugin configurations in Lamu.
//
// # Plugin Configurations (config.go)
//
// Plugins can define their own configuration structs that map directly from settings in the main config.toml file.
// The configuration struct must implement the lamu.Config interface.
//
// # Example Config Definition
//
//	package myplugin
//
//	import (
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	type MyPluginConfig struct {
//		ApiToken    string `toml:"api_token"`
//		MaxRetries  int    `toml:"max_retries"`
//		EnableCache bool   `toml:"enable_cache"`
//	}
//
//	// PostConfig executes sanity checks and assigns default values after TOML values are loaded.
//	func (c *MyPluginConfig) PostConfig() {
//		if c.MaxRetries <= 0 {
//			c.MaxRetries = 3
//		}
//	}
//
//	var PluginConfig = &MyPluginConfig{}
//
//	func pluginConfigs() lamu.PluginFeatures[lamu.Config] {
//		return lamu.PluginFeatures[lamu.Config]{
//			Entries: []registry.Pair[string, lamu.Config]{
//				{Key: "myplugin", Value: PluginConfig},
//			},
//		}
//	}
//
//	// Registering inside lamu.Plugin
//	lamu.Plugin{
//		Configs: lamu.PluginStages(pluginConfigs),
//	}
package config
