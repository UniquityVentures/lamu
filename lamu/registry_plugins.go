package lamu

import (
	"net/http"
	"net/url"
	"slices"

	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
	"gorm.io/gorm"
)

type PluginType int

const (
	// For plugins that add new models and functionality, ideally independent of other plugins
	PluginTypeApp = iota
	// For plugins that add additional functionality to App
	PluginTypeAddon
	// For plugins that add a long running service
	PluginTypeService
)

// PluginFeatures collects registry entries plus optional patches keyed like Entries.
// See package doc for patch purity and idempotency requirements.
type PluginFeatures[T any] struct {
	Entries []registry.Pair[string, T]
	Patches []registry.Pair[string, func(T) T]
}

// Build returns registry pairs with patches applied in registration order.
//
// Patches must be pure and idempotent: they must not mutate their argument T in place, and
// applying the same patch again to its own output must yield an equivalent result. Registry
// assembly may invoke Build more than once while merging plugins.
func (f *PluginFeatures[T]) Build() []registry.Pair[string, T] {
	entries := slices.Clone(f.Entries)
	for i := range len(entries) {
		for _, v := range f.Patches {
			if v.Key != entries[i].Key {
				continue
			}
			entries[i].Value = v.Value(entries[i].Value)
		}
	}
	return entries
}

// Merge concatenates Entries and Patches. Order is preserved so patch application order stays deterministic.
func (f PluginFeatures[T]) Merge(others ...PluginFeatures[T]) PluginFeatures[T] {
	result := f
	for _, other := range others {
		result.Entries = append(result.Entries, other.Entries...)
		result.Patches = append(result.Patches, other.Patches...)
	}
	return result
}

type Plugin struct {
	Type        PluginType
	Icon        string
	URL         *url.URL
	VerboseName string
	Roles       []string
	// Feature callbacks are invoked during registry assembly and may be called
	// more than once. They must be deterministic and repeat-safe; see package
	// documentation for patch purity and idempotency requirements.
	Migrations       []func() PluginFeatures[UsefulFilesystem]
	Views            []func() PluginFeatures[*views.View]
	Routes           []func() PluginFeatures[Route]
	Pages            []func() PluginFeatures[components.PageInterface]
	Models           []func() PluginFeatures[any]
	Layers           []func() PluginFeatures[views.GlobalLayer]
	Generators       []func() PluginFeatures[Generator]
	DBInitHooks      []func() PluginFeatures[DBInitHook]
	Configs          []func() PluginFeatures[Config]
	CommandFactories []func() PluginFeatures[CommandFactory]
}

var RegistryPlugin *registry.ImmutableRegistry[Plugin] = &registry.ImmutableRegistry[Plugin]{}

func CorePlugin(db *gorm.DB, config LamuConfig) registry.Pair[string, Plugin] {
	layers := PluginFeatures[views.GlobalLayer]{}
	layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.AttachRequestLayer", Value: views.AttachRequestLayer{}})
	layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.DbLayer", Value: DBLayer{DB: db}})
	if config.Debug {
		layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.LoggingLayer", Value: LoggingLayer{}})
		layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.CacheDisableLayer", Value: CacheDisableLayer{}})
	}
	layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.HtmxBoostLayer", Value: HtmxBoostLayer{}})
	layers.Entries = append(layers.Entries, registry.Pair[string, views.GlobalLayer]{Key: "core.EnvironmentLayer", Value: EnvironmentLayer{}})

	return registry.Pair[string, Plugin]{
		Key: "core", Value: Plugin{
			Type: PluginTypeAddon,
			URL: &url.URL{
				Path: "/",
			},
			VerboseName: "Core",
			Roles:       []string{"superuser", "admin"},
			Views: []func() PluginFeatures[*views.View]{
				func() PluginFeatures[*views.View] {
					return PluginFeatures[*views.View]{
						Entries: []registry.Pair[string, *views.View]{
							{Key: "core.HomeView", Value: GetPageView("core.HomePage")},
						},
					}
				},
			},
			Pages: []func() PluginFeatures[components.PageInterface]{
				func() PluginFeatures[components.PageInterface] {
					return PluginFeatures[components.PageInterface]{
						Entries: []registry.Pair[string, components.PageInterface]{
							{Key: "core.HomePage", Value: components.ShellBase{}},
						},
					}
				},
			},
			Layers: []func() PluginFeatures[views.GlobalLayer]{
				func() PluginFeatures[views.GlobalLayer] {
					return layers
				},
			},
			Routes: []func() PluginFeatures[Route]{
				func() PluginFeatures[Route] {
					return PluginFeatures[Route]{
						Entries: []registry.Pair[string, Route]{
							{Key: "core.HomeRoute", Value: Route{Path: "/", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
								w.WriteHeader(http.StatusOK)
								w.Write([]byte("Hello, World!"))
							})}},
						},
					}
				},
			},
		},
	}
}
