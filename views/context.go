package views

import (
	"context"
	"maps"

	"github.com/UniquityVentures/lamu/getters"
)

// ContextWithMap updates a map inside a context under the specified key.
// It retrieves the existing map from the context (or initializes a new one if missing),
// copies the provided keys and values into it using [maps.Copy], and returns the updated context.
func ContextWithMap[K comparable, V any](ctx context.Context, m map[K]V, key any) context.Context {
	ctxM, _ := ctx.Value(key).(map[K]V)
	if ctxM == nil {
		ctxM = map[K]V{}
	}
	maps.Copy(ctxM, m)
	return context.WithValue(ctx, key, ctxM)
}

// ContextWithErrorsAndValues merges input form parameter values and validation errors maps into the request context.
// It maps values to [getters.ContextKeyIn] and errors to [getters.ContextKeyError] so form input elements can retrieve them during render phases.
func ContextWithErrorsAndValues(ctx context.Context, values map[string]any, errors map[string]error) context.Context {
	return ContextWithMap(ContextWithMap(ctx, values, getters.ContextKeyIn), errors, getters.ContextKeyError)
}
