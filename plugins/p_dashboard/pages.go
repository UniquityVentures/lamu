package p_dashboard

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/lamu"
	pcomps "github.com/UniquityVentures/lamu/plugins/p_dashboard/components"
	"github.com/UniquityVentures/lamu/registry"
)

func init() {
	components.RegistryTopbar.Register("dashboard.appsPageButton", components.ButtonLink{
		Icon:    "squares-2x2",
		Link:    lamu.RoutePath("dashboard.AppsPage", nil),
		Classes: "btn-sm btn-square btn-neutral",
	})
	components.RegistryTopbar.Register("dashboard.themeButton", pcomps.ThemeButton{
		Classes: "btn-sm btn-square btn-outline",
	})
	components.RegistryTopbar.Register("dashboard.userDropdown", pcomps.UserDropdown{})
}

func pluginPages() lamu.PluginFeatures[components.PageInterface] {
	return lamu.PluginFeatures[components.PageInterface]{
		Entries: []registry.Pair[string, components.PageInterface]{
			{Key: "dashboard.HomeRedirectStub", Value: &components.ContainerColumn{
				Page:     components.Page{Key: "dashboard.HomeRedirectStub"},
				Children: []components.PageInterface{},
			}},
			{Key: "dashboard.AppsPage", Value: &components.ShellTopbarScaffold{
				Children: []components.PageInterface{
					&components.LayoutSimple{
						Page: components.Page{Key: "dashboard.AppsPageLayout"},
						Children: []components.PageInterface{
							&pcomps.AppsGrid{
								Page: components.Page{Key: "dashboard.AppsGrid"},
							},
						},
					},
				},
			}},
		},
	}
}
