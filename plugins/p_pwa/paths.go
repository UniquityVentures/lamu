package p_pwa

import (
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{
				Key: "pwa.ManifestRoute",
				Value: lamu.Route{
					Path:    "/app.webmanifest",
					Handler: lamu.NewDynamicView(manifestViewKey),
				},
			},
			{
				Key: "pwa.ServiceWorkerRoute",
				Value: lamu.Route{
					Path:    "/serviceworker.js",
					Handler: lamu.NewDynamicView(serviceWorkerViewKey),
				},
			},
			{
				Key: "pwa.OfflineRoute",
				Value: lamu.Route{
					Path:    "/offline",
					Handler: lamu.NewDynamicView(offlineViewKey),
				},
			},
			{
				Key: "pwa.assetLinks",
				Value: lamu.Route{
					Path:    "/.well-known/assetlinks.json",
					Handler: lamu.NewDynamicView(assetLinksViewKey),
				},
			},
			{
				Key: "pwa.StaticPwaBaseRoute",
				Value: lamu.Route{
					Path:    "/static/pwa/",
					Handler: lamu.NewDynamicView(staticPwaViewKey),
				},
			},
			{
				Key: "pwa.StaticPwaFilesRoute",
				Value: lamu.Route{
					Path:    "/static/pwa/{path...}",
					Handler: lamu.NewDynamicView(staticPwaViewKey),
				},
			},
		},
	}
}
