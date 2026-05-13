package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
)

func BuildAllRegistries(allPlugins []registry.Pair[string, Plugin]) {
	commandFactories := PluginFeatures[CommandFactory]{}
	configs := PluginFeatures[Config]{}
	dbInitHooks := PluginFeatures[DBInitHook]{}
	generators := PluginFeatures[Generator]{}
	layers := PluginFeatures[views.GlobalLayer]{}
	migrations := PluginFeatures[UsefulFilesystem]{}
	models := PluginFeatures[any]{}
	pages := PluginFeatures[components.PageInterface]{}
	routes := PluginFeatures[Route]{}
	views := PluginFeatures[*views.View]{}

	for _, pair := range allPlugins {
		plugin := pair.Value

		commandFactories.Merge(plugin.CommandFactories)
		configs.Merge(plugin.Configs)
		dbInitHooks.Merge(plugin.DBInitHooks)
		generators.Merge(plugin.Generators)
		layers.Merge(plugin.Layers)
		migrations.Merge(plugin.Migrations)
		models.Merge(plugin.Models)
		pages.Merge(plugin.Pages)
		routes.Merge(plugin.Routes)
		views.Merge(plugin.Views)
	}

	*RegistryCommand = registry.NewImmutableRegistry(commandFactories.Build())
	*RegistryConfig = registry.NewImmutableRegistry(configs.Build())
	*RegistryDBInit = registry.NewImmutableRegistry(dbInitHooks.Build())
	*RegistryGenerator = registry.NewImmutableRegistry(generators.Build())
	*RegistryLayer = registry.NewImmutableRegistry(layers.Build())
	*RegistryMigrations = registry.NewImmutableRegistry(migrations.Build())
	*RegistryModel = registry.NewImmutableRegistry(models.Build())
	*RegistryPage = registry.NewImmutableRegistry(pages.Build())
	*RegistryPlugin = registry.NewImmutableRegistry(allPlugins)
	*RegistryRoute = registry.NewImmutableRegistry(routes.Build())
	*RegistryView = registry.NewImmutableRegistry(views.Build())
}
