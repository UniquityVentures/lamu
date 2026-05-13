package p_export

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginPages() lamu.PluginFeatures[components.PageInterface] {
	return lamu.PluginFeatures[components.PageInterface]{
		Entries: []registry.Pair[string, components.PageInterface]{
			{Key: "export.Menu", Value: components.SidebarMenu{
				Title: getters.Static("Export"),
				Back: &components.SidebarMenuItem{
					Title: getters.Static("Back to All Apps"),
					Url:   lamu.RoutePath("dashboard.AppsPage", nil),
				},
				Children: []components.PageInterface{
					components.SidebarMenuItem{
						Title:  getters.Static("XLSX Export"),
						Url:    lamu.RoutePath("export.PageRoute", nil),
						Active: true,
					},
				},
			}},
			{Key: "export.Page", Value: &components.ShellScaffold{
				Sidebar: []components.PageInterface{
					lamu.DynamicPage{Name: "export.Menu"},
				},
				Children: []components.PageInterface{
					exportPickerPage{},
				},
			}},
		},
	}
}
