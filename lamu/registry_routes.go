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

// Route carries a ServeMux-compatible pattern ([net/http], Go 1.22+) and handler.
// Wildcards like {id} must each occupy a full path segment: use "/users/u/{id}/delete/"
// not "/users{id}/delete/". When building paths as base+suffix, the base should end
// with "/" before appending a segment that starts with "{" (e.g. const AppUrl = "/users/"
// so AppUrl+"{id}/" is valid). If sibling paths include fixed literals and {id} segments
// under the same prefix (e.g. /users/roles/… vs /users/{id}/…), add a disambiguating literal
// segment (e.g. /users/u/{id}/…).
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
