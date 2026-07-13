// Package models contains explanations and code examples for database models in Lamu.
//
// # Database Models (models.go)
//
// Database schemas in Lamu are defined as GORM models and auto-migrated on startup.
// Define your database entities as structs and register them inside the plugin.
//
// # Model Definition and Registration Example
//
//	package myplugin
//
//	import (
//		"time"
//		"github.com/UniquityVentures/lamu/lamu"
//		"github.com/UniquityVentures/lamu/registry"
//	)
//
//	type BlogPost struct {
//		ID        uint      `gorm:"primaryKey"`
//		Title     string    `gorm:"size:255;not null"`
//		Body      string    `gorm:"type:text"`
//		CreatedAt time.Time
//		UpdatedAt time.Time
//	}
//
//	func pluginModels() lamu.PluginFeatures[any] {
//		return lamu.PluginFeatures[any]{
//			Entries: []registry.Pair[string, any]{
//				{Key: "blog.post", Value: &BlogPost{}},
//			},
//		}
//	}
//
//	// Registering inside lamu.Plugin
//	lamu.Plugin{
//		Models: lamu.PluginStages(pluginModels),
//	}
//
// # GORM Reference
//
// For details on database tag properties, associations, and queries, refer to the [gorm.io/gorm] package documentation.
package models
