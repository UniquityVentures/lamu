// Package app contains explanations and code examples for the plugin app.go file in Lamu.
//
// # Plugin Definition (app.go)
//
// Every plugin must define an app.go file that implements the Plugin entrypoint.
// This is done by implementing a GetPlugin function returning a registry.Pair wrapping
// the plugin's key and its lamu.Plugin configuration struct.
//
// The plugin registers lifecycle stages like Views, Pages, Routes, Models, CommandFactories, and Migrations.
//
// # Example GetPlugin Signature
//
//	package myplugin
//
//	import (
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	// GetPlugin returns the registry contributions for this plugin.
//	func GetPlugin() registry.Pair[string, lamu.Plugin] {
//		return registry.Pair[string, lamu.Plugin]{
//			Key: "myplugin",
//			Value: lamu.Plugin{
//				Type:             lamu.PluginTypeApp, // or lamu.PluginTypeAddon
//				VerboseName:      "My Custom Feature",
//				Views:            lamu.PluginStages(pluginViews),
//				Pages:            lamu.PluginStages(pluginPages),
//				Routes:           lamu.PluginStages(pluginRoutes),
//				Models:           lamu.PluginStages(pluginModels),
//				Migrations:       lamu.PluginStages(pluginMigrations),
//				CommandFactories: lamu.PluginStages(pluginCommandFactories),
//			},
//		}
//	}
package app
