package lamu

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"strings"

	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/registry"
)

var RegistryRoute *registry.ImmutableRegistry[Route] = &registry.ImmutableRegistry[Route]{}

type Route struct {
	Path    string
	Handler http.Handler
}

func GetRouter(config LamuConfig) *http.ServeMux {
	baseRouter := http.NewServeMux()
	if config.Debug {
		baseRouter.Handle("/pprof/", pprof.Handler("heap"))
	}
	routes := RegistryRoute.All()
	for _, route := range routes {
		// Keep exact-match behavior for "directory-like" routes that end with "/"
		// (so "/foo/" doesn't also match "/foo/bar"). For non-slash routes like
		// "/app.webmanifest", register them directly since "/app.webmanifest{$}"
		// is not a valid ServeMux pattern.
		if strings.HasSuffix(route.Path, "/") {
			baseRouter.Handle(route.Path+"{$}", route.Handler)
		} else {
			baseRouter.Handle(route.Path, route.Handler)
		}
	}
	return baseRouter
}

// RoutePath returns a Getter that resolves to the route's Path string.
func RoutePath(name string, args map[string]getters.Getter[any]) getters.Getter[string] {
	return func(ctx context.Context) (string, error) {
		if route, ok := RegistryRoute.Get(name); ok {
			r := route.Path
			for k, g := range args {
				v, err := g(ctx)
				if err != nil {
					return "", err
				}
				r = strings.ReplaceAll(r, fmt.Sprintf("{%s}", k), fmt.Sprintf("%v", v))
			}
			return r, nil
		}
		return "", fmt.Errorf("Route for %s not found", name)
	}
}
