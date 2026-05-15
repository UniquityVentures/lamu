package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/views"
)

func GetPageView(pageName string) *views.View {
	return &views.View{
		PageName: pageName,
		PageLookup: func(name string) (components.PageInterface, bool) {
			return RegistryPage.Get(name)
		},
	}
}
