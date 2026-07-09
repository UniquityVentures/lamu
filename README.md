# Lamu Web Framework

[![Go Reference](https://pkg.go.dev/badge/github.com/UniquityVentures/lamu.svg)](https://pkg.go.dev/github.com/UniquityVentures/lamu)

Lamu is a modular, plugin-based web application framework for Go. It features dynamic registry-based layouts, hot-reloadable plugin features, transactional views, and database schema migrations managed per plugin.

For detailed package documentation, check out the [Lamu Go Reference](https://pkg.go.dev/github.com/UniquityVentures/lamu/lamu).

## Quickstart

Initialize a Lamu application by registering active plugins, loading configuration values from a TOML file, and executing the CLI entrypoint:

```go
package main

import (
	"log"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/plugins/p_dashboard"
	"github.com/UniquityVentures/lamu/plugins/p_users"
	"github.com/UniquityVentures/lamu/registry"
)

func main() {
	// 1. Register the list of active plugins to load into the application kernel.
	plugins := []registry.Pair[string, lamu.Plugin]{
		registry.NewPair("dashboard", p_dashboard.GetPlugin()),
		registry.NewPair("users", p_users.GetPlugin()),
	}

	// 2. Load database settings, server addresses, and plugin parameters from config.toml.
	config, err := lamu.LoadConfigFromFile("config.toml", plugins)
	if err != nil {
		log.Fatalf("failed loading configuration file: %v", err)
	}

	// 3. Build global registries and run the Cobra CLI bootstrapper.
	if err := lamu.Start(config, plugins); err != nil {
		log.Fatalf("failed executing application entry: %v", err)
	}
}
```

## Features

- **Plugin Registries**: Package database models, pages, API routes, and configs inside modular plugin boundaries.
- **Transactional View Layers**: Compose request pipelines with built-in or custom middleware layers to handle detail loading, form updates, and deletions.
- **Goose Migrations**: Keep SQL database migrations decoupled and isolated inside plugin subdirectory systems.
