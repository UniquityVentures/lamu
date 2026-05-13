package p_otp

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/plugins/p_users"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginPages() lamu.PluginFeatures[components.PageInterface] {
	auth := pageEntriesOtpAuth()
	prefs := pageEntriesOtpPreferences()
	entries := make([]registry.Pair[string, components.PageInterface], 0, len(auth)+len(prefs))
	entries = append(entries, auth...)
	entries = append(entries, prefs...)

	return lamu.PluginFeatures[components.PageInterface]{
		Entries: entries,
		Patches: []registry.Pair[string, func(components.PageInterface) components.PageInterface]{
			{Key: "p_users.LoginPage", Value: patchUsersLoginPageWithOtpForgotLink},
		},
	}
}

func patchUsersLoginPageWithOtpForgotLink(page components.PageInterface) components.PageInterface {
	if scaffold, ok := page.(*components.ShellAuthScaffold); ok {
		components.InsertChildAfter(scaffold,
			"p_users.AuthForm",
			func(*components.FormComponent[p_users.User]) *components.ButtonLink {
				return &components.ButtonLink{
					Label: "Forgot password?",
					Link:  lamu.RoutePath("otp.ForgotPasswordRoute", nil),
				}
			})
		return scaffold
	}
	panic("Base page for login page was not ShellAuthScaffold")
}
