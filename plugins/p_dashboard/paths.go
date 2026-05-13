package p_dashboard

import "github.com/UniquityVentures/lamu/lamu"

// pluginRoutes returns HTTP routes for this plugin. The dashboard is view-driven;
// add [lamu.Route] entries here when this plugin exposes standalone handlers.
func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{}
}
