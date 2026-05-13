package p_filesystem

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pageEntriesMenus() []registry.Pair[string, components.PageInterface] {
	return []registry.Pair[string, components.PageInterface]{
		{Key: "filesystem.MainMenu", Value: &components.SidebarMenu{
			Title: getters.Static("Filesystem"),
			Back: &components.SidebarMenuItem{
				Title: getters.Static("Back to All Apps"),
				Url:   lamu.RoutePath("dashboard.AppsPage", nil),
			},
			Children: []components.PageInterface{
				&components.SidebarMenuItem{Title: getters.Static("All Files"), Url: lamu.RoutePath("filesystem.ListRoute", nil), Icon: "folder-open"},
				&components.SidebarMenuItem{Title: getters.Static("Create Item"), Url: lamu.RoutePath("filesystem.CreateRoute", nil), Icon: "plus"},
				&components.SidebarMenuItem{Title: getters.Static("Bulk Upload"), Url: lamu.RoutePath("filesystem.MultiUploadRoute", nil), Icon: "arrow-up-tray"},
			},
		}},
		{Key: "filesystem.VNodeMenu", Value: &components.SidebarMenu{
			Title: currentVNodeTitle(),
			Back: &components.SidebarMenuItem{
				Title: getters.Static("Back"),
				Url:   currentVNodeBackRoute(),
			},
			Children: []components.PageInterface{
				&components.SidebarMenuItem{Title: getters.Static("View Details"), Url: lamu.RoutePath("filesystem.DetailRoute", map[string]getters.Getter[any]{
					"id": getters.Any(getters.Key[uint]("vnode.ID")),
				}), Icon: "eye"},
				&components.SidebarMenuItem{Title: getters.Static("Edit"), Url: lamu.RoutePath("filesystem.UpdateRoute", map[string]getters.Getter[any]{
					"id": getters.Any(getters.Key[uint]("vnode.ID")),
				}), Icon: "pencil-square"},
				&components.SidebarMenuItem{Title: getters.Static("Move"), Url: lamu.RoutePath("filesystem.MoveRoute", map[string]getters.Getter[any]{
					"id": getters.Any(getters.Key[uint]("vnode.ID")),
				}), Icon: "arrow-right-circle"},
				&components.ShowIf{
					Getter: getters.Any(getters.Key[bool]("vnode.IsDirectory")),
					Children: []components.PageInterface{
						&components.SidebarMenuItem{Title: getters.Static("Browse Contents"), Url: lamu.RoutePath("filesystem.BrowseRoute", map[string]getters.Getter[any]{
							"parent_id": getters.Any(getters.Key[uint]("vnode.ID")),
						}), Icon: "folder-open"},
						&components.SidebarMenuItem{Title: getters.Static("Add New Item"), Url: lamu.RoutePath("filesystem.CreateChildRoute", map[string]getters.Getter[any]{
							"parent_id": getters.Any(getters.Key[uint]("vnode.ID")),
						}), Icon: "plus"},
						&components.SidebarMenuItem{Title: getters.Static("Bulk Upload"), Url: lamu.RoutePath("filesystem.MultiUploadChildRoute", map[string]getters.Getter[any]{
							"parent_id": getters.Any(getters.Key[uint]("vnode.ID")),
						}), Icon: "arrow-up-tray"},
					},
				},
			},
		}},
	}
}
