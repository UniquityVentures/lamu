package p_dashboard

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

// pluginRoutes serves the apps grid at [AppUrl]. Registry key matches
// [lamu.RoutePath]("dashboard.AppsPage", …) used across menu and redirects.
func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{Key: "dashboard.AppsPage", Value: lamu.Route{
				Path:    AppUrl,
				Handler: lamu.NewDynamicView("dashboard.AppsView"),
			}},
		},
	}
}
