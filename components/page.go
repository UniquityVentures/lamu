package components

import (
	"context"
	"reflect"
	"slices"

	"maragu.dev/gomponents"
)

// PageInterface represents the standard component interface within the Lamu UI framework.
// Every custom page layout or interactive component must satisfy this interface.
type PageInterface interface {
	// Build compiles the component using context values and returns a gomponents HTML node structure.
	Build(context.Context) gomponents.Node
	// GetKey returns the unique key identifying this specific component.
	GetKey() string
	// GetRoles returns the authorized roles allowed to view or interact with this component.
	GetRoles() []string
}

// Page struct defines common properties embedded in all component structs.
// It carries the unique component key and routing roles configuration.
type Page struct {
	// Key represents the unique component key identifier.
	Key   string
	// Roles represents a slice of authorized roles required to view this component.
	Roles []string
}

// Render compiles the page component if the role in ctx (under key "$role") matches the required roles.
// If the user's role is unauthorized, it returns an empty gomponents.Group node instead of the rendered output.
func Render(p PageInterface, ctx context.Context) gomponents.Node {
	roles := GetRequiredRoles(p)
	currentRole, _ := ctx.Value("$role").(string)
	if roles == nil {
		return p.Build(ctx)
	}

	if slices.Contains(roles, currentRole) {
		return p.Build(ctx)
	}
	return gomponents.Group{}
}

// GetRequiredRoles extracts the required roles list configured in the component's embedded Page structure using reflection.
func GetRequiredRoles(p PageInterface) []string {
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	page, ok := v.FieldByName("Page").Interface().(Page)
	if !ok {
		return nil
	}
	return page.Roles
}
