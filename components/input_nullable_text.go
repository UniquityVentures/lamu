package components

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strings"

	"github.com/UniquityVentures/lamu/getters"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// InputNullableText renders a text input for optional *string model fields.
// Parse returns nil when the submitted value is empty so cleared fields persist as SQL NULL.
type InputNullableText struct {
	Page
	Label    string
	Name     string
	Getter   getters.Getter[*string]
	Required bool
	Classes  string
	Hidden   bool
	Attr     getters.Getter[Node]
}

func (e InputNullableText) GetKey() string {
	return e.Key
}

func (e InputNullableText) GetRoles() []string {
	return e.Roles
}

func (e InputNullableText) Build(ctx context.Context) Node {
	var valueNode Node = Value("")
	if e.Getter != nil {
		value, err := e.Getter(ctx)
		if err != nil {
			slog.Error("InputNullableText getter failed", "error", err, "key", e.Key)
		} else if value != nil {
			valueNode = Value(*value)
		}
	}

	wrapClass := fmt.Sprintf("my-1 %s", e.Classes)
	if e.Hidden {
		wrapClass += " hidden"
	}
	return Div(Class(wrapClass),
		Label(Class("label text-sm font-bold flex flex-col items-start gap-1"),
			If(!e.Hidden, Text(e.Label)),
			Input(If(!e.Hidden, Type("text")), If(e.Hidden, Type("hidden")), Name(e.Name),
				valueNode,
				Class(fmt.Sprintf("input input-bordered w-full %s", e.Classes)),
				If(e.Required, Required()),
				Iff(e.Attr != nil, func() (out Node) {
					out = Raw("")
					defer func() {
						if r := recover(); r != nil {
							slog.Error("InputNullableText attr getter panicked", "panic", r, "key", e.Key)
						}
					}()
					n, err := e.Attr(ctx)
					if err != nil {
						slog.Error("InputNullableText attr getter failed", "error", err, "key", e.Key)
						return out
					}
					if n == nil {
						return out
					}
					v := reflect.ValueOf(n)
					if (v.Kind() == reflect.Pointer || v.Kind() == reflect.Map || v.Kind() == reflect.Slice || v.Kind() == reflect.Interface || v.Kind() == reflect.Func) && v.IsNil() {
						return out
					}
					return n
				}),
			),
		),
	)
}

func (e InputNullableText) Parse(v any, _ context.Context) (any, error) {
	vals, _ := v.([]string)
	if len(vals) == 0 {
		return (*string)(nil), nil
	}
	raw := strings.TrimSpace(vals[0])
	if raw == "" {
		return (*string)(nil), nil
	}
	return &raw, nil
}

func (e InputNullableText) GetName() string {
	return e.Name
}
