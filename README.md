# Lamu Web Framework

[![Go Reference](https://pkg.go.dev/badge/github.com/UniquityVentures/lamu.svg)](https://pkg.go.dev/github.com/UniquityVentures/lamu)

Lamu is a modular, plugin-based web application framework for Go. It features dynamic registry-based layouts, hot-reloadable plugin features, transactional views, and database schema migrations managed per plugin.

For detailed package documentation, check out the [Lamu Go Reference](https://pkg.go.dev/github.com/UniquityVentures/lamu/lamu).

## Quickstart

Create a empty go project named lamu_test

```bash
mkdir lamu_test
go mod init lamu_test
go get github.com/UniquityVentures/lamu@latest
```

Create an empty main.go, and an empty config.toml file

```bash
touch main.go
```

To configure for a basic sqlite deployment, put the following in config.toml:
```toml
Debug = true
DBType = "Sqlite"
Address = ":42069"

[SqliteConfig]
  DSN = "lamu_database.db"

[plugins.p_pwa]
  # If set, /serviceworker.js will serve this file. If empty, p_pwa serves a minimal default.
  serviceWorkerPath = ""

  # If set, /offline will render this view key. If empty, p_pwa serves a minimal offline HTML page.
  offlineViewName = ""

  staticDir = "./pwa_static/"

  PWA_APP_NAME = "Lamu Test"
  PWA_APP_DESCRIPTION = "Test app for lamu"
  PWA_APP_THEME_COLOR = "#0A0302"
  PWA_APP_BACKGROUND_COLOR = "#ffffff"
  PWA_APP_DISPLAY = "standalone"
  PWA_APP_SCOPE = "/"
  PWA_APP_ORIENTATION = "any"
  PWA_APP_START_URL = "/"
  PWA_APP_STATUS_BAR_COLOR = "default"
  PWA_APP_DIR = "ltr"
  PWA_APP_LANG = "en-US"

```

To initialize a Lamu application by registering active plugins, loading configuration values from a TOML file, and executing the CLI entrypoint, put the following in main.go

```go
package main

import (
	"log"
	"gorm.io/driver/sqlite"

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
	// Load database settings, server addresses, and plugin parameters from config.toml.
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

To run, 

```bash
go mod tidy
go run main.go
```

## Features

- **Plugin Registries**: Package database models, pages, API routes, and configs inside modular plugin boundaries.
- **Transactional View Layers**: Compose request pipelines with built-in or custom middleware layers to handle detail loading, form updates, and deletions.
- **Goose Migrations**: Keep SQL database migrations decoupled and isolated inside plugin subdirectory systems.
