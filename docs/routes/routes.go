// Package routes contains explanations and code examples for HTTP routing in Lamu.
//
// # HTTP Routing (routes.go)
//
// Lamu routes map request path patterns directly to handlers.
// The routing layer is built on top of Go 1.22+'s native net/http router (ServeMux), supporting path variables natively.
//
// # 1. Registering Routes
//
// Define routes inside the plugin and register them into the lamu.Plugin stages:
//
//	package myplugin
//
//	import (
//		"net/http"
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	func handleDetails(w http.ResponseWriter, r *http.Request) {
//		id := r.PathValue("id")
//		w.Write([]byte("Fetching details for: " + id))
//	}
//
//	func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
//		return lamu.PluginFeatures[lamu.Route]{
//			Entries: []registry.Pair[string, lamu.Route]{
//				{
//					Key: "myplugin.details",
//					Value: lamu.Route{
//						Path:    "/details/{id}/",
//						Handler: http.HandlerFunc(handleDetails),
//					},
//				},
//			},
//		}
//	}
//
// # 2. Patching Routes
//
// You can patch or override routes declared by other plugins to inject custom middleware or change path details:
//
//	func patchExistingRoutes() lamu.PluginFeatures[lamu.Route] {
//		return lamu.PluginFeatures[lamu.Route]{
//			Patches: []registry.Pair[string, func(lamu.Route) lamu.Route]{
//				{
//					Key: "core.HomeRoute",
//					Value: func(existing lamu.Route) lamu.Route {
//						return lamu.Route{
//							Path:    existing.Path,
//							Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//								w.Write([]byte("Patched landing page!"))
//							}),
//						}
//					},
//				},
//			},
//		}
//	}
//
// # Go Router Reference
//
// For native ServeMux path wildcard rules, see standard library [net/http#ServeMux] documentation.
package routes
