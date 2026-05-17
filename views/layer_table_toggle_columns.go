package views

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/getters"
)

// LayerTableToggleColumns parses a URL query parameter (name from QueryParam getter) as a
// comma-separated list of [components.TableColumn.Name] values, builds map[string]bool, and stores it on
// context under the key from ContextKey getter. If the query parameter is absent, context is unchanged so
// [components.DataTable] getters can return (nil, nil) and show every column.
//
// When the parameter is present with an empty value (e.g. cols=), the stored map is empty and every
// named column is hidden for tables using [components.GetterEnabledColumnsFromContext].
type LayerTableToggleColumns struct {
	QueryParam getters.Getter[string]
	ContextKey getters.Getter[string]
}

func (m LayerTableToggleColumns) Next(_ View, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		paramName, err := m.QueryParam(ctx)
		if err != nil {
			slog.Error("views: LayerTableToggleColumns: query param name", "error", err)
			next.ServeHTTP(w, r)
			return
		}
		keyName, err := m.ContextKey(ctx)
		if err != nil {
			slog.Error("views: LayerTableToggleColumns: context key", "error", err)
			next.ServeHTTP(w, r)
			return
		}
		q := r.URL.Query()
		if _, ok := q[paramName]; !ok {
			next.ServeHTTP(w, r)
			return
		}
		raw := ""
		if vals := q[paramName]; len(vals) > 0 {
			raw = vals[0]
		}
		parsed := components.ParseEnabledTableColumnsParam(raw)
		ctx = context.WithValue(ctx, keyName, parsed)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
