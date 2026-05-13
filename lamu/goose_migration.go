package lamu

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"unicode"

	"github.com/pressly/goose/v3"
)

// GooseVersionTableName returns a goose version table name for a migrations registry key:
// deterministic, Postgres/SQLite-safe, one table per plugin.
func GooseVersionTableName(registryKey string) string {
	var b strings.Builder
	b.WriteString("goose_migrations__")
	prevUnderscore := false
	for _, r := range registryKey {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(unicode.ToLower(r))
			prevUnderscore = false
		case r == '.' || r == '-' || r == '_':
			if !prevUnderscore {
				b.WriteByte('_')
				prevUnderscore = true
			}
		default:
			if !prevUnderscore {
				b.WriteByte('_')
				prevUnderscore = true
			}
		}
	}
	s := strings.Trim(b.String(), "_")
	if len(s) > len("goose_migrations__") {
		return s
	}
	return "goose_migrations__default"
}

func gooseDialect(t DBType) (goose.Dialect, error) {
	switch t {
	case DBTypeSqlite:
		return goose.DialectSQLite3, nil
	case DBTypePostgres:
		return goose.DialectPostgres, nil
	default:
		return "", fmt.Errorf("unsupported DBType for goose: %v", t)
	}
}

// gooseUpPluginMigrations runs goose Up per plugin migrations FS ([RegistryMigrations.AllStable]
// order), separate version table per key so numeric versions can overlap across plugins.
func gooseUpPluginMigrations(ctx context.Context, sqlDB *sql.DB, config LamuConfig) error {
	pairs := *RegistryMigrations.AllStable()
	if len(pairs) == 0 {
		return nil
	}
	dialect, err := gooseDialect(config.DBType)
	if err != nil {
		return err
	}
	for _, pair := range pairs {
		sub, err := fs.Sub(pair.Value, "migrations")
		if err != nil {
			return fmt.Errorf("migrations subdirectory for %q: %w", pair.Key, err)
		}
		p, err := goose.NewProvider(
			dialect,
			sqlDB,
			sub,
			goose.WithTableName(GooseVersionTableName(pair.Key)),
			goose.WithDisableGlobalRegistry(true),
		)
		if errors.Is(err, goose.ErrNoMigrations) {
			continue
		}
		if err != nil {
			return fmt.Errorf("goose NewProvider(%q): %w", pair.Key, err)
		}
		if _, err := p.Up(ctx); err != nil {
			return fmt.Errorf("goose Up for %q: %w", pair.Key, err)
		}
	}
	return nil
}
