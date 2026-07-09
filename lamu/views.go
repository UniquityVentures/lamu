package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/views"
)

// GetPageView initializes and returns a standard view controller wrapper [views.View] that resolves and renders page layouts from the global [RegistryPage].
//
// Use Cases:
//   - Defining basic HTML routes or static pages that require page registry lookups without custom middleware layers.
//
// Example:
//
//	var DashboardHomeView = lamu.GetPageView("dashboard.home")
func GetPageView(pageName string) *views.View {
	return &views.View{
		PageName: pageName,
		PageLookup: func(name string) (components.PageInterface, bool) {
			return RegistryPage.Get(name)
		},
	}
}
