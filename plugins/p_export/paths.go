package p_export

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{Key: "export.PageRoute", Value: lamu.Route{
				Path:    AppUrl,
				Handler: lamu.NewDynamicView("export.PageView"),
			}},
			{Key: "export.DownloadRoute", Value: lamu.Route{
				Path:    AppUrl + "download/",
				Handler: lamu.NewDynamicView("export.DownloadView"),
			}},
		},
	}
}
