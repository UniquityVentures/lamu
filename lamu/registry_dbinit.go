package lamu

import (
	"log"

	"github.com/UniquityVentures/lamu/registry"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBInitHook runs after core DB setup (migrations, callbacks). Hooks run in registration order.
type DBInitHook func(*gorm.DB) *gorm.DB

// RegistryDBInit stores DB init hooks; iterate with [registry.RegisterOrder] to preserve registration order.
// [AllStable] returns an internal cached slice — do not mutate it.
var RegistryDBInit *registry.ImmutableRegistry[DBInitHook] = &registry.ImmutableRegistry[DBInitHook]{}

func GetDbConn(config LamuConfig) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch config.DBType {
	case DBTypeSqlite:
		dialector = sqlite.New(*config.SqliteConfig)
	case DBTypePostgres:
		dialector = postgres.New(*config.PostgresConfig)
	default:
		log.Panicf("Unrecognized db type %s", config.DBType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}

	// Configure hard delete - skip soft delete and actually remove rows
	db.Callback().Delete().Before("gorm:delete").Register("lamu:hard_delete", func(db *gorm.DB) {
		// Set Unscoped to true to force hard delete instead of soft delete
		db.Statement.Unscoped = true
	})
	return db, nil
}

func InitDB(db *gorm.DB, config LamuConfig) error {
	var dialector gorm.Dialector

	switch config.DBType {
	case DBTypeSqlite:
		dialector = sqlite.New(*config.SqliteConfig)
	case DBTypePostgres:
		dialector = postgres.New(*config.PostgresConfig)
	default:
		log.Panicf("Unrecognized db type %s", config.DBType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}

	// Configure hard delete - skip soft delete and actually remove rows
	db.Callback().Delete().Before("gorm:delete").Register("lamu:hard_delete", func(db *gorm.DB) {
		// Set Unscoped to true to force hard delete instead of soft delete
		db.Statement.Unscoped = true
	})

	for _, p := range *RegistryDBInit.AllStable() {
		db = p.Value(db)
	}
	return nil
}
