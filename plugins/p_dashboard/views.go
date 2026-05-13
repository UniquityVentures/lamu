package p_dashboard

import (
	"context"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/plugins/p_users"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
)

func pluginViews() lamu.PluginFeatures[*views.View] {
	return lamu.PluginFeatures[*views.View]{
		Entries: []registry.Pair[string, *views.View]{
			{Key: "dashboard.AppsView", Value: lamu.GetPageView("dashboard.AppsPage").WithLayer("p_users.auth", p_users.AuthenticationLayer{})},
		},
		Patches: []registry.Pair[string, func(*views.View) *views.View]{
			{Key: "p_users.LoginSuccessView", Value: func(_ *views.View) *views.View {
				return lamu.RedirectView(lamu.RoutePath("dashboard.AppsPage", nil))
			}},
			// core.HomeView: core.HomeRoute renders this; dashboard sends logged-in users to apps, others to login.
			{Key: "core.HomeView", Value: func(_ *views.View) *views.View {
				return lamu.GetPageView("dashboard.HomeRedirectStub").
					WithLayer("p_users.optional_auth", p_users.OptionalAuthLayer{}).
					WithLayer("dashboard.home_root_redirect", lamu.RedirectLayer{URLGetter: func(ctx context.Context) (string, error) {
						if p_users.UserPresentInContext(ctx) {
							return lamu.RoutePath("dashboard.AppsPage", nil)(ctx)
						}
						return lamu.RoutePath("p_users.LoginRoute", nil)(ctx)
					}})
			}},
		},
	}
}
