package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
)

// FillRegistry merges feature bundles from each plugin, then assigns an immutable registry after
// each merge by calling [PluginFeatures.Build]. Patches must be pure and idempotent; see package lamu doc.
func FillRegistry[T any](features []func() PluginFeatures[T], targetRegistry *registry.ImmutableRegistry[T]) {
	finalFeatures := PluginFeatures[T]{}
	for _, feature := range features {
		if feature == nil {
			continue
		}
		finalFeatures = finalFeatures.Merge(feature())
		*targetRegistry = registry.NewImmutableRegistry(finalFeatures.Build())
	}
}

func MapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}

func BuildAllRegistries(allPlugins []registry.Pair[string, Plugin]) {
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[UsefulFilesystem] {
		return pair.Value.Migrations
	}), RegistryMigrations)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[DBInitHook] {
		return pair.Value.DBInitHooks
	}), RegistryDBInit)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[Config] {
		return pair.Value.Configs
	}), RegistryConfig)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[Generator] {
		return pair.Value.Generators
	}), RegistryGenerator)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[CommandFactory] {
		return pair.Value.CommandFactories
	}), RegistryCommand)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[any] {
		return pair.Value.Models
	}), RegistryModel)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[views.GlobalLayer] {
		return pair.Value.Layers
	}), RegistryLayer)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[components.PageInterface] {
		return pair.Value.Pages
	}), RegistryPage)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[*views.View] {
		return pair.Value.Views
	}), RegistryView)
	FillRegistry(MapSlice(allPlugins, func(pair registry.Pair[string, Plugin]) func() PluginFeatures[Route] {
		return pair.Value.Routes
	}), RegistryRoute)

	// Installed-plugin metadata for tools like dashboard.AppsGrid (PluginType filter, RBAC tiles).
	*RegistryPlugin = registry.NewImmutableRegistry(allPlugins)
}
