package components

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/registry"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ClientTabsLayout controls tab ribbon orientation; content is always below the ribbon (column stack).
type ClientTabsLayout uint8

const (
	// ClientTabsLayoutResponsive: narrow view uses a horizontal tab row; md+ stacks tab buttons vertically (default).
	ClientTabsLayoutResponsive ClientTabsLayout = 0
	// ClientTabsLayoutVertical: tab buttons are stacked vertically; ribbon is above the panel content (column layout).
	ClientTabsLayoutVertical ClientTabsLayout = 1
	// ClientTabsLayoutHorizontal: tab buttons stay in a horizontal row (wrap on narrow widths).
	ClientTabsLayoutHorizontal ClientTabsLayout = 2
)

// ClientTabs renders client-side Alpine tabs: ribbon + [ClientMatchIf] content.
type ClientTabs struct {
	Page
	Tabs     []registry.Pair[string, getters.Getter[PageInterface]]
	Default  getters.Getter[string]
	StateKey string
	// Layout selects ribbon orientation; zero is [ClientTabsLayoutResponsive].
	Layout      ClientTabsLayout
	Attr        getters.Getter[Node]
	RibbonAttr  getters.Getter[Node]
	ContentAttr getters.Getter[Node]
	// DiscoveryChildren, when non-nil, is returned from [ClientTabs.GetChildren] so
	// [FindChildren] can locate nested components (e.g. filter forms) inside tab panels.
	// Use the same nodes referenced by [ClientTabs.Tabs] pair values.
	DiscoveryChildren []PageInterface
}

func (e ClientTabs) GetChildren() []PageInterface {
	return e.DiscoveryChildren
}

func (e ClientTabs) layoutClasses() (outer, ribbon, button string) {
	switch e.Layout {
	case ClientTabsLayoutVertical:
		// Stacked tab buttons; content is always below the ribbon (never beside).
		return "flex flex-col gap-4",
			"flex w-full flex-col gap-1 rounded-box border border-base-300 bg-base-100 p-1",
			"btn w-full justify-start"
	case ClientTabsLayoutHorizontal:
		return "flex flex-col gap-4",
			"flex w-full flex-row flex-wrap gap-1 rounded-box border border-base-300 bg-base-100 p-1",
			"btn flex-1 min-w-[5rem] justify-center"
	default:
		// Same column stack as other layouts; only the ribbon’s flex direction changes at md.
		return "flex flex-col gap-4",
			"flex w-full flex-row gap-1 rounded-box border border-base-300 bg-base-100 p-1 md:flex-col",
			"btn flex-1 md:flex-none md:w-full justify-center md:justify-start"
	}
}

func (e ClientTabs) Build(ctx context.Context) Node {
	if len(e.Tabs) == 0 {
		return Group{}
	}

	keys := make([]string, 0, len(e.Tabs))
	match := make(map[string]PageInterface, len(e.Tabs))
	for _, pair := range e.Tabs {
		key := pair.Key
		pageGetter := pair.Value
		if pageGetter == nil {
			continue
		}
		page, err := pageGetter(ctx)
		if err != nil {
			return ContainerError{Error: getters.Static(err)}.Build(ctx)
		}
		if page == nil {
			continue
		}
		keys = append(keys, key)
		match[key] = page
	}
	if len(keys) == 0 {
		return Group{}
	}

	stateKey := e.StateKey
	if stateKey == "" {
		stateKey = "tab"
	}

	defaultTab := keys[0]
	if e.Default != nil {
		if selected, err := e.Default(ctx); err != nil {
			return ContainerError{Error: getters.Static(err)}.Build(ctx)
		} else if _, ok := match[selected]; ok {
			defaultTab = selected
		}
	}
	xData, err := json.Marshal(map[string]string{stateKey: defaultTab})
	if err != nil {
		return ContainerError{Error: getters.Static(err)}.Build(ctx)
	}

	outerClass, ribbonClass, buttonClass := e.layoutClasses()

	ribbon := Group{}
	for _, key := range keys {
		ribbon = append(ribbon, Button(
			Type("button"),
			Class(buttonClass),
			Attr("@click", fmt.Sprintf("%s = %q", stateKey, key)),
			Attr(":class", fmt.Sprintf("%s === %q ? 'btn-primary' : 'btn-ghost'", stateKey, key)),
			Text(key),
		))
	}

	return Div(
		Class(outerClass),
		Attr("x-data", string(xData)),
		Iff(e.Attr != nil, func() Node {
			n, err := e.Attr(ctx)
			if err != nil {
				return ContainerError{Error: getters.Static(err)}.Build(ctx)
			}
			if n == nil {
				return Group{}
			}
			return n
		}),
		Div(
			Class(ribbonClass),
			Iff(e.RibbonAttr != nil, func() Node {
				n, err := e.RibbonAttr(ctx)
				if err != nil {
					return ContainerError{Error: getters.Static(err)}.Build(ctx)
				}
				if n == nil {
					return Group{}
				}
				return n
			}),
			ribbon,
		),
		Div(
			Class("min-w-0 flex-1"),
			Iff(e.ContentAttr != nil, func() Node {
				n, err := e.ContentAttr(ctx)
				if err != nil {
					return ContainerError{Error: getters.Static(err)}.Build(ctx)
				}
				if n == nil {
					return Group{}
				}
				return n
			}),
			Render(ClientMatchIf{
				Key:   getters.Static(stateKey),
				Match: getters.Static(match),
			}, ctx),
		),
	)
}

func (e ClientTabs) GetKey() string {
	return e.Key
}

func (e ClientTabs) GetRoles() []string {
	return e.Roles
}
