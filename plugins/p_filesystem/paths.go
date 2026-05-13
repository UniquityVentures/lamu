package p_filesystem

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{Key: "filesystem.ListRoute", Value: lamu.Route{
				Path:    AppUrl,
				Handler: lamu.NewDynamicView("filesystem.ListView"),
			}},
			{Key: "filesystem.BrowseRoute", Value: lamu.Route{
				Path:    AppUrl + "browse/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.BrowseView"),
			}},
			{Key: "filesystem.SelectRoute", Value: lamu.Route{
				Path:    AppUrl + "select/",
				Handler: lamu.NewDynamicView("filesystem.SelectView"),
			}},
			{Key: "filesystem.SelectChildRoute", Value: lamu.Route{
				Path:    AppUrl + "select/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.SelectChildView"),
			}},
			{Key: "filesystem.MultiSelectRoute", Value: lamu.Route{
				Path:    AppUrl + "multi-select/",
				Handler: lamu.NewDynamicView("filesystem.MultiSelectView"),
			}},
			{Key: "filesystem.MultiSelectChildRoute", Value: lamu.Route{
				Path:    AppUrl + "multi-select/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.MultiSelectChildView"),
			}},
			{Key: "filesystem.MoveSelectRoute", Value: lamu.Route{
				Path:    AppUrl + "move-select/",
				Handler: lamu.NewDynamicView("filesystem.MoveSelectView"),
			}},
			{Key: "filesystem.MoveSelectChildRoute", Value: lamu.Route{
				Path:    AppUrl + "move-select/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.MoveSelectChildView"),
			}},
			{Key: "filesystem.CreateRoute", Value: lamu.Route{
				Path:    AppUrl + "create/",
				Handler: lamu.NewDynamicView("filesystem.CreateView"),
			}},
			{Key: "filesystem.CreateChildRoute", Value: lamu.Route{
				Path:    AppUrl + "create/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.CreateChildView"),
			}},
			{Key: "filesystem.MultiUploadRoute", Value: lamu.Route{
				Path:    AppUrl + "upload/",
				Handler: lamu.NewDynamicView("filesystem.MultiUploadView"),
			}},
			{Key: "filesystem.MultiUploadChildRoute", Value: lamu.Route{
				Path:    AppUrl + "upload/in/{parent_id}/",
				Handler: lamu.NewDynamicView("filesystem.MultiUploadChildView"),
			}},
			{Key: "filesystem.DetailRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/",
				Handler: lamu.NewDynamicView("filesystem.DetailView"),
			}},
			{Key: "filesystem.DownloadRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/download/",
				Handler: lamu.NewDynamicView("filesystem.DownloadView"),
			}},
			{Key: "filesystem.UpdateRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/edit/",
				Handler: lamu.NewDynamicView("filesystem.UpdateView"),
			}},
			{Key: "filesystem.DeleteRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/delete/",
				Handler: lamu.NewDynamicView("filesystem.DeleteView"),
			}},
			{Key: "filesystem.MoveRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/move/",
				Handler: lamu.NewDynamicView("filesystem.MoveView"),
			}},
		},
	}
}
