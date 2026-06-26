package views

import (
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// QueryPatcherSearch applies a case-insensitive contains match across multiple
// columns when the "search" query parameter is non-empty (?search=term).
// Matching uses OR across the listed DB column names.
type QueryPatcherSearch[T any] struct {
	Columns []string
}

func (p QueryPatcherSearch[T]) Patch(_ View, r *http.Request, db gorm.ChainInterface[T]) gorm.ChainInterface[T] {
	term := strings.TrimSpace(r.URL.Query().Get("search"))
	if term == "" || len(p.Columns) == 0 {
		return db
	}

	pattern := "%" + term + "%"
	clauses := make([]string, len(p.Columns))
	args := make([]any, len(p.Columns))
	for i, col := range p.Columns {
		clauses[i] = col + " ILIKE ?"
		args[i] = pattern
	}
	return db.Where(strings.Join(clauses, " OR "), args...)
}
