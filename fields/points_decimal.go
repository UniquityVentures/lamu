package fields

import (
	"database/sql/driver"
	"encoding"
	"fmt"
	"math/big"
	"strings"
)

// DecimalSix holds a monetary-style amount with exactly 6 decimal places
// persisted as NUMERIC and represented in Go with *big.Rat.
type DecimalSix struct {
	R *big.Rat
}

var (
	_ encoding.TextMarshaler   = DecimalSix{}
	_ encoding.TextUnmarshaler = (*DecimalSix)(nil)
	_ driver.Valuer            = DecimalSix{}
)

// NormalizeSixDecimals rounds R to exactly 6 decimal places
func (p DecimalSix) NormalizeDecimals() DecimalSix {
	r := new(big.Rat)
	if p.R == nil {
		r = big.NewRat(0, 1)
	} else {
		r = r.Set(p.R)
	}
	r = r.Mul(r, big.NewRat(1000000, 1))
	r.SetInt(new(big.Int).Div(r.Num(), r.Denom()))
	r = r.Quo(r, big.NewRat(1000000, 1))
	return DecimalSix{R: r}
}

// MarshalText implements encoding.TextMarshaler.
func (p DecimalSix) MarshalText() ([]byte, error) {
	r := p.NormalizeDecimals().R
	return []byte(r.FloatString(6)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler (mapstructure form binding).
func (p *DecimalSix) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))
	if s == "" {
		p.R = big.NewRat(0, 1)
		return nil
	}
	r := new(big.Rat)
	if _, ok := r.SetString(s); !ok {
		return fmt.Errorf("invalid points value %q", s)
	}
	p0 := p.NormalizeDecimals()
	p = &p0
	return nil
}

// Value implements driver.Valuer for GORM / SQL.
func (p DecimalSix) Value() (driver.Value, error) {
	return p.NormalizeDecimals().R.FloatString(6), nil
}

// Scan implements sql.Scanner.
func (p *DecimalSix) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		p.R = big.NewRat(0, 1)
		return nil
	case []byte:
		return p.UnmarshalText(v)
	case string:
		return p.UnmarshalText([]byte(v))
	case int64:
		p.R = big.NewRat(v, 1)
		return nil
	default:
		return fmt.Errorf("cannot scan %T into PointsDecimal", src)
	}
}

// String returns a fixed 6-decimal string for UI.
func (p DecimalSix) String() string {
	b, err := p.MarshalText()
	if err != nil {
		return "0.000000"
	}
	return string(b)
}
