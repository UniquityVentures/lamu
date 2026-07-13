// Package quickstart contains explanations, quickstart guides, and code examples for bootstrapping a Lamu application.
//
// # Quickstart
//
// To bootstrap a Lamu application, create a main entry file (e.g. main.go) loading active plugins
// and calling the CLI bootstrapper:
//
//	package main
//
//	import (
//		"log"
//
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/plugins/p_dashboard"
//		"github.com/UniquityVentures/lamu/plugins/p_users"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	func main() {
//		// 1. Register the list of active plugins to load into the application kernel.
//		plugins := []registry.Pair[string, lamu.Plugin]{
//			p_dashboard.GetPlugin(),
//			p_users.GetPlugin(),
//		}
//
//		// 2. Load database settings, server addresses, and plugin parameters from config.toml.
//		config, err := lamu.LoadConfigFromFile("config.toml", plugins)
//		if err != nil {
//			log.Fatalf("failed loading configuration file: %v", err)
//		}
//
//		// 3. Build global registries and run the Cobra CLI bootstrapper.
//		if err := lamu.Start(config, plugins); err != nil {
//			log.Fatalf("failed executing application entry: %v", err)
//		}
//	}
//
// Line-by-line Breakdown:
//   - Step 1: Defines a slice of [registry.Pair] mapping plugin names to their [lamu.Plugin] configurations.
//   - Step 2: Calls [lamu.LoadConfigFromFile] to decode configurations, open GORM connections, and execute initial schemas.
//   - Step 3: Invokes [lamu.Start] to initialize the CLI command tree (Root starts the server, generate seeds, tui boots TUI).
//
// # Minimal Plugin
//
// Plugins are discrete packages exposing a GetPlugin function returning a [registry.Pair] of plugin name and [lamu.Plugin]:
//
//	package myplugin
//
//	import (
//		"net/url"
//
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	func GetPlugin() registry.Pair[string, lamu.Plugin] {
//		return registry.NewPair("myplugin", lamu.Plugin{
//			Type:        lamu.PluginTypeApp,
//			VerboseName: "Inventory Manager",
//			Icon:        "box",
//			URL: &url.URL{
//				Path: "/inventory/",
//			},
//		})
//	}
//
// Explanation:
//   - Type: Specifies [lamu.PluginTypeApp] for standalone logic, [lamu.PluginTypeAddon] (which hides the plugin from the dashboard's app grid), or [lamu.PluginTypeService].
//   - VerboseName & Icon: Defines the display name and icon used on the admin landing page.
//   - URL: Represents the primary landing URL path pointing to the plugin home view.
//
// # Adding Routes
//
// To register endpoint paths, implement a Route entry under [lamu.Plugin.Routes] returning page layouts:
//
//	package myplugin
//
//	import (
//		"context"
//		"net/http"
//
//		"github.com/UniquityVentures/lamu/components"
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//		. "maragu.dev/gomponents"
//		"maragu.dev/gomponents/html"
//	)
//
//	type HelloPage struct {
//		components.Page
//	}
//
//	func (p HelloPage) Build(ctx context.Context) Node {
//		return html.Div(html.H1(Text("Hello, World!")))
//	}
//
//	func GetPlugin() registry.Pair[string, lamu.Plugin] {
//		return registry.NewPair("hello_plugin", lamu.Plugin{
//			Type:        lamu.PluginTypeApp,
//			VerboseName: "Hello Plugin",
//			Pages: lamu.PluginStages(func() lamu.PluginFeatures[components.PageInterface] {
//				return lamu.PluginFeatures[components.PageInterface]{
//					Entries: []registry.Pair[string, components.PageInterface]{
//						registry.NewPair("myplugin.hello", HelloPage{}),
//					},
//				}
//			}),
//			Routes: lamu.PluginStages(func() lamu.PluginFeatures[lamu.Route] {
//				return lamu.PluginFeatures[lamu.Route]{
//					Entries: []registry.Pair[string, lamu.Route]{
//						registry.NewPair("hello_route", lamu.Route{
//							Path:    "/hello/",
//							Handler: lamu.GetPageView("myplugin.hello"),
//						}),
//					},
//				}
//			}),
//		})
//	}
//
// Explanation:
//   - [components.PageInterface] represents visual layout templates containing HTML elements structure.
//   - [lamu.GetPageView] constructs standard view controller handlers referencing page keys in the registry.
//   - [lamu.Route] maps paths to handlers, dynamically resolving wildcard parameters if path segments match.
//
// # Adding Views and Layers
//
// Views wrap page components in middleware pipeline layers. Here is an example of mapping custom middlewares:
//
//	package myplugin
//
//	import (
//		"log"
//		"net/http"
//
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//		"github.com/UniquityVentures/lamu/views"
//	)
//
//	type LogLayer struct{}
//
//	func (l LogLayer) Next(view views.View, next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			log.Printf("Executing view page: %s", view.PageName)
//			next.ServeHTTP(w, r)
//		})
//	}
//
//	func GetPlugin() registry.Pair[string, lamu.Plugin] {
//		myView := lamu.GetPageView("myplugin.hello").
//			WithLayer("logger", LogLayer{})
//
//		return registry.NewPair("views_plugin", lamu.Plugin{
//			Type:        lamu.PluginTypeApp,
//			VerboseName: "Views Plugin",
//			Views: lamu.PluginStages(func() lamu.PluginFeatures[*views.View] {
//				return lamu.PluginFeatures[*views.View]{
//					Entries: []registry.Pair[string, *views.View]{
//						registry.NewPair("hello_view", myView),
//					},
//				}
//			}),
//		})
//	}
//
// Other available view layers:
//   - [views.LayerDetail]: Loads database records by context key indicators.
//   - [views.LayerUpdate]: Updates database rows inside transactions on POST request actions.
//   - [views.LayerCreate]: Inserts new database records on POST request actions.
//   - [views.LayerDelete]: Deletes database records on POST request actions.
//   - [views.LayerList]: Queries database collections lists mapping them to data lists.
//   - [views.LayerSingleton]: Manages site configurations single-row databases entries.
//
// # Patching Existing Features of a Plugin
//
// Plugins can modify pages, views, or configurations registered by other plugins using Patches:
//
//	package myaddon
//
//	import (
//		"github.com/UniquityVentures/lamu/components"
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	func GetPlugin() registry.Pair[string, lamu.Plugin] {
//		return registry.NewPair("dashboard_decorator", lamu.Plugin{
//			Type:        lamu.PluginTypeAddon,
//			VerboseName: "Dashboard Decorator",
//			Pages: lamu.PluginStages(func() lamu.PluginFeatures[components.PageInterface] {
//				return lamu.PluginFeatures[components.PageInterface]{
//					Patches: []registry.Pair[string, func(components.PageInterface) components.PageInterface]{
//						registry.NewPair("p_dashboard.home", func(existing components.PageInterface) components.PageInterface {
//							// Verify insertion element doesn't exist yet to maintain idempotency
//							if components.HasChild(existing, "patched_banner") {
//								return existing
//							}
//							banner := components.Header{Key: "patched_banner", Title: "Patched Dashboard"}
//							return components.InsertChildFirst(existing, banner)
//						}),
//					},
//				}
//			}),
//		})
//	}
//
// Purity and Idempotency Rules:
//
//   - Pure: Do not mutate input arguments in place. Return a copy or new value if modifying pointer fields.
//   - Idempotent: Patch application must yield equivalent results if run multiple times (verify keys before appends).
//   - Merge Safety: Features merges execute in sequence. Package state variables must not be mutated.
//
// # Next Steps
//
// For a detailed breakdown of the application file structure, standard plugin files (app.go, config.go, pages.go, migrations.go, routes.go, models.go, views.go, commands.go), and architectural concepts (layers.go, components.go, querypatchers.go), refer to the documentation package: [github.com/UniquityVentures/lamu/docs].
package quickstart
