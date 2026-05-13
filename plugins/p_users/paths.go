package p_users

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{Key: "base.HomeRoute", Value: lamu.Route{
				Path:    "/",
				Handler: lamu.NewDynamicView("base.HomeView"),
			}},
			{Key: "users.ListRoute", Value: lamu.Route{
				Path:    AppUrl,
				Handler: lamu.NewDynamicView("users.ListView"),
			}},
			{Key: "users.CreateRoute", Value: lamu.Route{
				Path:    AppUrl + "create/",
				Handler: lamu.NewDynamicView("users.CreateView"),
			}},
			{Key: "users.DetailRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/",
				Handler: lamu.NewDynamicView("users.DetailView"),
			}},
			{Key: "users.UpdateRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/edit/",
				Handler: lamu.NewDynamicView("users.UpdateView"),
			}},
			{Key: "users.SelfDetailRoute", Value: lamu.Route{
				Path:    AppUrl + "self/",
				Handler: lamu.NewDynamicView("users.SelfDetailView"),
			}},
			{Key: "users.SelfUpdateRoute", Value: lamu.Route{
				Path:    AppUrl + "self/edit/",
				Handler: lamu.NewDynamicView("users.SelfUpdateView"),
			}},
			{Key: "users.SelfChangePasswordRoute", Value: lamu.Route{
				Path:    AppUrl + "self/change-password/",
				Handler: lamu.NewDynamicView("users.SelfChangePasswordView"),
			}},
			{Key: "users.DeleteRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/delete/",
				Handler: lamu.NewDynamicView("users.DeleteView"),
			}},
			{Key: "users.ChangePasswordRoute", Value: lamu.Route{
				Path:    AppUrl + "{id}/change-password/",
				Handler: lamu.NewDynamicView("users.ChangePasswordView"),
			}},
			{Key: "users.SelectRoute", Value: lamu.Route{
				Path:    AppUrl + "select/",
				Handler: lamu.NewDynamicView("users.SelectView"),
			}},
			{Key: "users.RoleSelectRoute", Value: lamu.Route{
				Path:    RoleUrl + "select/",
				Handler: lamu.NewDynamicView("users.RoleSelectView"),
			}},
			{Key: "users.RoleListRoute", Value: lamu.Route{
				Path:    RoleUrl,
				Handler: lamu.NewDynamicView("users.RoleListView"),
			}},
			{Key: "users.RoleCreateRoute", Value: lamu.Route{
				Path:    RoleUrl + "create/",
				Handler: lamu.NewDynamicView("users.RoleCreateView"),
			}},
			{Key: "users.RoleDetailRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/",
				Handler: lamu.NewDynamicView("users.RoleDetailView"),
			}},
			{Key: "users.RoleUpdateRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/edit/",
				Handler: lamu.NewDynamicView("users.RoleUpdateView"),
			}},
			{Key: "users.RoleDeleteRoute", Value: lamu.Route{
				Path:    RoleUrl + "{id}/delete/",
				Handler: lamu.NewDynamicView("users.RoleDeleteView"),
			}},
			{Key: "users.LoginRoute", Value: lamu.Route{
				Path:    AppUrl + "login/",
				Handler: lamu.NewDynamicView("users.LoginView"),
			}},
			{Key: "users.SignupRoute", Value: lamu.Route{
				Path:    AppUrl + "signup/",
				Handler: lamu.NewDynamicView("users.SignupView"),
			}},
			{Key: "users.LoginSuccessRoute", Value: lamu.Route{
				Path:    AppUrl + "success/",
				Handler: lamu.NewDynamicView("users.LoginSuccessView"),
			}},
			{Key: "users.UnauthenticatedRoute", Value: lamu.Route{
				Path:    AppUrl + "unauthenticated/",
				Handler: lamu.NewDynamicView("users.UnauthenticatedView"),
			}},
			{Key: "users.LogoutRoute", Value: lamu.Route{
				Path:    AppUrl + "logout/",
				Handler: lamu.NewDynamicView("users.LogoutView"),
			}},
		},
	}
}
