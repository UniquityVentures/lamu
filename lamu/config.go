package lamu

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/UniquityVentures/lamu/registry"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

type LamuConfig struct {
	Debug          bool
	DBType         DBType
	SqliteConfig   *sqlite.Config
	PostgresConfig *postgres.Config
	Address        string
	UDS            string
	GeneratorOrder []string
	TrustedOrigins []string
	Plugins        map[string]toml.Primitive
}

type DBType string

const (
	DBTypeSqlite   = DBType("Sqlite")
	DBTypePostgres = DBType("Postgres")
)

// LoadConfigFromFile decodes the top-level config TOML, then decodes each [Plugins.<key>] section
// into the plugin config pointer registered for that key. It calls [BuildAllRegistries] first
// so RegistryConfig is populated; pass the same plugins slice you pass to [Start]. For tools that
// only need core DB settings, pass nil plugins.
func LoadConfigFromFile(path string, plugins []registry.Pair[string, Plugin]) (LamuConfig, error) {
	var config LamuConfig

	if path == "" {
		return config, fmt.Errorf("config path is empty")
	}

	resolvedPath := path
	if !filepath.IsAbs(resolvedPath) {
		exe, err := os.Executable()
		if err != nil {
			slog.Error("failed resolving executable path for config file", "err", err, "configPath", path)
			return config, err
		}
		resolvedPath = filepath.Join(filepath.Dir(exe), resolvedPath)
	}

	md, err := toml.DecodeFile(resolvedPath, &config)
	if err != nil {
		slog.Error("failed decoding config file", "err", err, "configPath", path, "resolvedPath", resolvedPath)
		return config, err
	}

	db, err := GetDbConn(config)
	if err != nil {
		return config, err
	}
	BuildAllRegistries(append([]registry.Pair[string, Plugin]{CorePlugin(db, config)}, plugins...))
	if err := InitDB(db, config); err != nil {
		return config, err
	}

	for key, cfgPointer := range RegistryConfig.All() {
		if prim, ok := config.Plugins[key]; ok {
			err = md.PrimitiveDecode(prim, cfgPointer)
			if err != nil {
				slog.Error("failed decoding plugin config", "err", err, "plugin", key)
				return config, err
			}
		}
		// Run even when the app has no [Plugins.<key>] table, so plugins can require fields
		// (e.g. panic if mandatory secrets are missing) instead of silently skipping validation.
		cfgPointer.PostConfig()
	}

	return config, nil
}
