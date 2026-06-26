package components

import (
	"context"
	"fmt"
	"strings"

	"github.com/UniquityVentures/lamu/registry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

var RegistryTopbar = registry.NewRegistry[PageInterface]()

type SidebarItem struct {
	Icon    string
	Content PageInterface
}

var RegistryRightSidebar = registry.NewRegistry[SidebarItem]()

type LayoutTopbar struct {
	Page
	Children []PageInterface
}

func (e LayoutTopbar) Build(ctx context.Context) gomponents.Node {
	topbarItems := gomponents.Group{}

	for _, comp := range *RegistryTopbar.AllStable(registry.RegisterOrder[PageInterface]{}) {
		item := comp.Value
		topbarItems = append(topbarItems, Render(item, ctx))
	}

	// Fetch entries from RegistryRightSidebar
	rightSidebarEntries := *RegistryRightSidebar.AllStable(registry.RegisterOrder[SidebarItem]{})

	// Generate AlpineJS persistent state and control tags if there are sidebar items
	var xData string
	var keysJS string
	var defaultKey string
	if len(rightSidebarEntries) > 0 {
		defaultKey = rightSidebarEntries[0].Key

		var keysBuilder strings.Builder
		keysBuilder.WriteString("[")
		for i, entry := range rightSidebarEntries {
			if i > 0 {
				keysBuilder.WriteString(",")
			}
			keysBuilder.WriteString(fmt.Sprintf("%q", entry.Key))
		}
		keysBuilder.WriteString("]")
		keysJS = keysBuilder.String()

		xData = fmt.Sprintf(`{
			showRight: $persist(true).as('right-sidebar-show'),
			activeTab: $persist(%q).as('right-sidebar-active'),
			init() {
				const keys = %s;
				if (!keys.includes(this.activeTab) && keys.length > 0) {
					this.activeTab = keys[0];
				}
			},
			toggleRight() {
				this.showRight = !this.showRight;
			},
			setActiveTab(key) {
				this.activeTab = key;
			}
		}`, defaultKey, keysJS)

		// Add toggle button to the topbar navigation menu if at least one item exists
		topbarItems = append(topbarItems, html.Button(
			html.Class("btn btn-sm btn-square btn-ghost"),
			gomponents.Attr("@click", "toggleRight()"),
			Render(Icon{Name: "bars-3-bottom-right"}, ctx),
		))
	}

	childGroup := gomponents.Group{}
	for _, child := range e.Children {
		childGroup = append(childGroup, Render(child, ctx))
	}

	// Build the main layout
	var mainLayout gomponents.Node
	if len(rightSidebarEntries) > 0 {
		var asideAttrs []gomponents.Node
		asideAttrs = append(asideAttrs,
			html.Class("flex-none w-80 border-l border-base-300 bg-base-100 flex flex-col h-full overflow-hidden"),
			gomponents.Attr("x-show", "showRight"),
		)

		// Tab Buttons Row (only if more than 1 item)
		var tabRow gomponents.Node = gomponents.Group{}
		if len(rightSidebarEntries) > 1 {
			var tabButtons []gomponents.Node
			for _, entry := range rightSidebarEntries {
				tabButtons = append(tabButtons, html.Button(
					html.Class("btn btn-sm btn-square"),
					gomponents.Attr(":class", fmt.Sprintf("activeTab === %q ? 'btn-primary' : 'btn-ghost'", entry.Key)),
					gomponents.Attr("@click", fmt.Sprintf("setActiveTab(%q)", entry.Key)),
					Render(Icon{Name: entry.Value.Icon}, ctx),
				))
			}
			tabRow = html.Div(
				html.Class("flex items-center gap-2 border-b border-base-300 p-2 overflow-x-auto flex-none"),
				gomponents.Group(tabButtons),
			)
		}

		// Content Panes
		var contentPanels []gomponents.Node
		for _, entry := range rightSidebarEntries {
			contentPanels = append(contentPanels, html.Div(
				gomponents.Attr("x-show", fmt.Sprintf("activeTab === %q", entry.Key)),
				html.Class("h-full overflow-y-auto p-4"),
				Render(entry.Value.Content, ctx),
			))
		}
		contentArea := html.Div(
			html.Class("flex-1 overflow-hidden relative"),
			gomponents.Group(contentPanels),
		)

		mainLayout = html.Div(html.Class("flex-1 flex overflow-hidden"),
			html.Div(html.Class("flex-1 overflow-hidden"),
				childGroup,
			),
			html.Aside(
				append(asideAttrs,
					tabRow,
					contentArea,
				)...,
			),
		)
	} else {
		mainLayout = html.Div(html.Class("flex-1 overflow-hidden"),
			childGroup,
		)
	}

	rootAttrs := []gomponents.Node{
		html.Class("h-screen flex flex-col overflow-hidden"),
	}
	if len(rightSidebarEntries) > 0 {
		rootAttrs = append(rootAttrs, gomponents.Attr("x-data", xData))
	}
	rootAttrs = append(rootAttrs,
		html.Div(html.Class("navbar bg-base-100 border-b border-base-300 px-4 flex justify-between items-center flex-none"),
			html.Div(html.Class("flex-1")),
			html.Div(html.Class("flex-none flex items-center gap-2"),
				topbarItems,
			),
		),
		mainLayout,
	)

	return html.Div(rootAttrs...)
}

func (e LayoutTopbar) GetKey() string {
	return e.Key
}

func (e LayoutTopbar) GetRoles() []string {
	return e.Roles
}

func (e LayoutTopbar) GetChildren() []PageInterface {
	return e.Children
}

func (e *LayoutTopbar) SetChildren(children []PageInterface) {
	e.Children = children
}
