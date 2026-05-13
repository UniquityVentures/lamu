package p_otp

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{
				Key: "otp.ForgotPasswordRoute",
				Value: lamu.Route{
					Path:    "/otp/forgot-password/",
					Handler: lamu.NewDynamicView("otp.ForgotPasswordView"),
				},
			},
			{
				Key: "otp.PhoneOtpRequestRoute",
				Value: lamu.Route{
					Path:    "/otp/login/sms/",
					Handler: lamu.NewDynamicView("otp.PhoneOtpRequestView"),
				},
			},
			{
				Key: "otp.EmailOtpRequestRoute",
				Value: lamu.Route{
					Path:    "/otp/login/email/",
					Handler: lamu.NewDynamicView("otp.EmailOtpRequestView"),
				},
			},
			{
				Key: "otp.OtpVerifyRoute",
				Value: lamu.Route{
					Path:    "/otp/verify/",
					Handler: lamu.NewDynamicView("otp.OtpVerifyView"),
				},
			},
			{
				Key: "otp.OTPPreferencesRoute",
				Value: lamu.Route{
					Path:    "/otp/preferences/",
					Handler: lamu.NewDynamicView("otp.OTPPreferencesView"),
				},
			},
		},
	}
}
