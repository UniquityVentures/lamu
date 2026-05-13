package p_users

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

const (
	// Trailing slashes so AppUrl+"{…}/" concatenation yields whole wildcard segments for ServeMux.
	AppUrl  = "/users/"
	RoleUrl = "/users/roles/"
	// Routes keyed by DB user ID live under …/u/{id}/… so literals like …/roles/… never match …/{id}/….
	UserIDRoutePrefix = AppUrl + "u/"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Patches: []registry.Pair[string, func(lamu.Route) lamu.Route]{
			{
				Key: "core.HomeRoute",
				Value: func(old lamu.Route) lamu.Route {
					old.Path = "/"
					old.Handler = lamu.NewDynamicView("core.HomeView")
					return old
				},
			},
		},
		Entries: []registry.Pair[string, lamu.Route]{
			{Key: "p_users.ListRoute", Value: lamu.Route{
				Path:    AppUrl,
				Handler: lamu.NewDynamicView("p_users.ListView"),
			}},
			{Key: "p_users.CreateRoute", Value: lamu.Route{
				Path:    AppUrl + "create/",
				Handler: lamu.NewDynamicView("p_users.CreateView"),
			}},
			{Key: "p_users.DetailRoute", Value: lamu.Route{
				Path:    UserIDRoutePrefix + "{id}/",
				Handler: lamu.NewDynamicView("p_users.DetailView"),
			}},
			{Key: "p_users.UpdateRoute", Value: lamu.Route{
				Path:    UserIDRoutePrefix + "{id}/edit/",
				Handler: lamu.NewDynamicView("p_users.UpdateView"),
			}},
			{Key: "p_users.SelfDetailRoute", Value: lamu.Route{
				Path:    AppUrl + "self/",
				Handler: lamu.NewDynamicView("p_users.SelfDetailView"),
			}},
			{Key: "p_users.SelfUpdateRoute", Value: lamu.Route{
				Path:    AppUrl + "self/edit/",
				Handler: lamu.NewDynamicView("p_users.SelfUpdateView"),
			}},
			{Key: "p_users.SelfChangePasswordRoute", Value: lamu.Route{
				Path:    AppUrl + "self/change-password/",
				Handler: lamu.NewDynamicView("p_users.SelfChangePasswordView"),
			}},
			{Key: "p_users.DeleteRoute", Value: lamu.Route{
				Path:    UserIDRoutePrefix + "{id}/delete/",
				Handler: lamu.NewDynamicView("p_users.DeleteView"),
			}},
			{Key: "p_users.ChangePasswordRoute", Value: lamu.Route{
				Path:    UserIDRoutePrefix + "{id}/change-password/",
				Handler: lamu.NewDynamicView("p_users.ChangePasswordView"),
			}},
			{Key: "p_users.SelectRoute", Value: lamu.Route{
				Path:    AppUrl + "select/",
				Handler: lamu.NewDynamicView("p_users.SelectView"),
			}},
			{Key: "p_users.RoleSelectRoute", Value: lamu.Route{
				Path:    RoleUrl + "select/",
				Handler: lamu.NewDynamicView("p_users.RoleSelectView"),
			}},
			{Key: "p_users.RoleListRoute", Value: lamu.Route{
				Path:    RoleUrl,
				Handler: lamu.NewDynamicView("p_users.RoleListView"),
			}},
			{Key: "p_users.RoleCreateRoute", Value: lamu.Route{
				Path:    RoleUrl + "create/",
				Handler: lamu.NewDynamicView("p_users.RoleCreateView"),
			}},
			{Key: "p_users.RoleDetailRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/",
				Handler: lamu.NewDynamicView("p_users.RoleDetailView"),
			}},
			{Key: "p_users.RoleUpdateRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/edit/",
				Handler: lamu.NewDynamicView("p_users.RoleUpdateView"),
			}},
			{Key: "p_users.RoleDeleteRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/delete/",
				Handler: lamu.NewDynamicView("p_users.RoleDeleteView"),
			}},
			{Key: "p_users.LoginRoute", Value: lamu.Route{
				Path:    AppUrl + "login/",
				Handler: lamu.NewDynamicView("p_users.LoginView"),
			}},
			{Key: "p_users.SignupRoute", Value: lamu.Route{
				Path:    AppUrl + "signup/",
				Handler: lamu.NewDynamicView("p_users.SignupView"),
			}},
			{Key: "p_users.LoginSuccessRoute", Value: lamu.Route{
				Path:    AppUrl + "success/",
				Handler: lamu.NewDynamicView("p_users.LoginSuccessView"),
			}},
			{Key: "p_users.UnauthenticatedRoute", Value: lamu.Route{
				Path:    AppUrl + "unauthenticated/",
				Handler: lamu.NewDynamicView("p_users.UnauthenticatedView"),
			}},
			{Key: "p_users.LogoutRoute", Value: lamu.Route{
				Path:    AppUrl + "logout/",
				Handler: lamu.NewDynamicView("p_users.LogoutView"),
			}},
		},
	}
}
