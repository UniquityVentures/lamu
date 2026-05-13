package lamu

import (
	tea "charm.land/bubbletea/v2"
	"github.com/UniquityVentures/lamu/registry"
	"github.com/spf13/cobra"
)

// Start boots the cobra CLI and server paths. Populate registries before calling this via
// [LoadConfigFromFile(path, plugins)], which merges plugins and invokes [BuildAllRegistries];
// plugins must match the slice passed there.
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

	rootCmd.AddCommand(&cobra.Command{
		Use:   "tui",
		Short: "Launch the TUI instead of running the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := GetDbConn(config)
			if err != nil {
				return err
			}
			_, err = tea.NewProgram(initialModel(db)).Run()
			return err
		},
	})

	for _, pair := range *RegistryCommand.AllStable() {
		rootCmd.AddCommand(pair.Value(config))
	}

	return rootCmd.Execute()
}
