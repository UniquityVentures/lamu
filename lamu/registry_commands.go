package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
	"github.com/spf13/cobra"
)

// CommandFactory represents a generator function that builds Cobra CLI commands mapped to a specific [LamuConfig].
//
// Use Cases:
//   - Defining custom CLI sub-commands inside application plugins (e.g., system diagnostics, database cleaner tasks).
//
// Example:
//
//	var BackupCmdFactory CommandFactory = func(config LamuConfig) *cobra.Command {
//		return &cobra.Command{
//			Use:   "backup",
//			Short: "Executes a database schema backup",
//			Run: func(cmd *cobra.Command, args []string) {
//				executeBackup(config)
//			},
//		}
//	}
//
//	// Register the command factory inside your lamu.Plugin configuration:
//	lamu.Plugin{
//		CommandFactories: lamu.PluginStages(func() PluginFeatures[CommandFactory] {
//			return PluginFeatures[CommandFactory]{
//				Entries: []registry.Pair[string, CommandFactory]{
//					registry.NewPair("backup_db", BackupCmdFactory),
//				},
//			}
//		}),
//	}
//
//	// Register a patch to modify an existing command in another plugin:
//	lamu.Plugin{
//		CommandFactories: lamu.PluginStages(func() PluginFeatures[CommandFactory] {
//			return PluginFeatures[CommandFactory]{
//				Patches: []registry.Pair[string, func(CommandFactory) CommandFactory]{
//					registry.NewPair("backup_db", func(existing CommandFactory) CommandFactory {
//						return func(config LamuConfig) *cobra.Command {
//							cmd := existing(config)
//							cmd.Short = "Patched: " + cmd.Short
//							return cmd
//						}
//					}),
//				},
//			}
//		}),
//	}
//
//	// Retrieve a registered command factory:
//	factory, ok := RegistryCommand.Get("backup_db")
type CommandFactory func(LamuConfig) *cobra.Command

// RegistryCommand represents the global immutable registry mapping custom plugin sub-commands to their CommandFactory builders.
var RegistryCommand *registry.ImmutableRegistry[CommandFactory] = &registry.ImmutableRegistry[CommandFactory]{}
