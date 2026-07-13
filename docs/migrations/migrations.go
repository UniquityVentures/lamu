// Package migrations contains explanations and code examples for database migrations in Lamu.
//
// # Database Migrations (migrations.go)
//
// Database schema migrations in Lamu are managed sequentially using SQL scripts compatible with the goose engine.
// Migrations are isolated per plugin, embedded using Go's go:embed directive, and registered within the Plugin struct.
//
// # Generating SQL Migration Files
//
// To create a new goose SQL migration file in the plugin's migrations folder, execute the goose CLI command from the repository root:
//
//	goose -dir plugins/<plugin_name>/migrations create <migration_name> sql
//
// Example:
//
//	goose -dir plugins/blog/migrations create create_posts_table sql
//
// This produces a SQL script (e.g. plugins/blog/migrations/20260713161200_create_posts_table.sql) with:
//
//	-- +goose Up
//	CREATE TABLE posts ( ... );
//
//	-- +goose Down
//	DROP TABLE posts;
//
// # Example Migrations File
//
//	package myplugin
//
//	import (
//		"embed"
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	//go:embed migrations
//	var migrationsFS embed.FS
//
//	func pluginMigrations() lamu.PluginFeatures[lamu.UsefulFilesystem] {
//		return lamu.PluginFeatures[lamu.UsefulFilesystem]{
//			Entries: []registry.Pair[string, lamu.UsefulFilesystem]{
//				{Key: "myplugin.migrations", Value: migrationsFS},
//			},
//		}
//	}
package migrations
