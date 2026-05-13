package p_livereloading

import (
	"io"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"

	"golang.org/x/net/websocket"
)

func NilServer(ws *websocket.Conn) {
	io.Copy(io.Discard, ws)
}

func pluginRoutes() lamu.PluginFeatures[lamu.Route] {
	return lamu.PluginFeatures[lamu.Route]{
		Entries: []registry.Pair[string, lamu.Route]{
			{
				Key: "livereloading.ws",
				Value: lamu.Route{
					Path:    "/_livereload",
					Handler: websocket.Handler(NilServer),
				},
			},
		},
	}
}
