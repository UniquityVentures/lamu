package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
	"github.com/spf13/cobra"
)

// Start initializes and executes the Cobra CLI application, acting as the main entrypoint for any Lamu application.
//
// CLI Command Scopes:
//   - Root Command: Starts the HTTP web server via [StartServer].
//   - generate: Runs database seed generators via [RunGenerators].
//   - tui: Launches the Bubble Tea terminal user interface.
//   - Plugin Commands: Resolves and registers custom commands dynamically loaded from [RegistryCommand].
//
// Registries and configurations must be populated before invoking this function (e.g. using LoadConfigFromFile).
//
// Use Cases:
//   - Initializing the CLI bootstrapper in the main execution block of a Go application.
//
// Example:
//
//	func main() {
//		config := lamu.LamuConfig{
//			Port: 8080,
//			DB:   lamu.DBConfig{Driver: "postgres", DSN: "postgresql://..."},
//		}
//		plugins := []registry.Pair[string, lamu.Plugin]{
//			registry.NewPair("dashboard", p_dashboard.New()),
//		}
//		if err := lamu.Start(config, plugins); err != nil {
//			log.Fatal(err)
//		}
//	}
func Start(config LamuConfig, plugins []registry.Pair[string, Plugin]) error {
	_ = plugins
	rootCmd := &cobra.Command{
		Use:   "lamu",
		Short: "Lamu web framework",
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartServer(config)
		},
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "generate",
		Short: "Run data generators to seed the database",
		RunE: func(cmd *cobra.Command, args []string) error {
			RunGenerators(config)
			return nil
		},
	})

	for _, pair := range *RegistryCommand.AllStable() {
		rootCmd.AddCommand(pair.Value(config))
	}

	return rootCmd.Execute()
}
