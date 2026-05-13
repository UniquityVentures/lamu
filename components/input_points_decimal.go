package components

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/UniquityVentures/lamu/fields"
	"github.com/UniquityVentures/lamu/getters"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// InputPointsDecimal is a decimal-amount field for form Values that round-trip
// through [PointsDecimal] so [views.PopulateFromMap] receives the correct type
// (unlike [InputText], which yields a string and breaks mapstructure decode).
type InputPointsDecimal struct {
	Page
	Label    string
	Name     string
	Getter   getters.Getter[fields.DecimalSix]
	Required bool
	Classes  string
	Hidden   bool
}

func (e InputPointsDecimal) GetKey() string { return e.Key }

func (e InputPointsDecimal) GetRoles() []string { return e.Roles }

func (e InputPointsDecimal) Build(ctx context.Context) Node {
	text := ""
	if e.Getter != nil {
		pd, err := e.Getter(ctx)
		if err != nil {
			slog.Error("InputPointsDecimal getter failed", "error", err, "key", e.Key)
		} else {
			text = pd.String()
		}
	}
	wrapClass := fmt.Sprintf("my-1 %s", e.Classes)
	if e.Hidden {
		wrapClass += " hidden"
	}
	valueNode := Value(text)
	return Div(Class(wrapClass),
		Label(Class("label text-sm font-bold flex flex-col items-start gap-1"),
			If(!e.Hidden, Text(e.Label)),
			Input(If(!e.Hidden, Type("text")), If(e.Hidden, Type("hidden")), Name(e.Name),
				valueNode,
				Attr("inputmode", "decimal"),
				Class(fmt.Sprintf("input input-bordered w-full %s", e.Classes)),
				If(e.Required, Required()),
			),
		),
	)
}

func (e InputPointsDecimal) Parse(v any, _ context.Context) (any, error) {
	vals, _ := v.([]string)
	if len(vals) == 0 || strings.TrimSpace(vals[0]) == "" {
		var out fields.DecimalSix
		if err := out.UnmarshalText([]byte("")); err != nil {
			return fields.DecimalSix{}, err
		}
		return out, nil
	}
	var out fields.DecimalSix
	if err := out.UnmarshalText([]byte(strings.TrimSpace(vals[0]))); err != nil {
		return fields.DecimalSix{}, err
	}
	return out, nil
}

func (e InputPointsDecimal) GetName() string { return e.Name }
