package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
	"github.com/spf13/cobra"
)

type CommandFactory func(LamuConfig) *cobra.Command

var RegistryCommand *registry.ImmutableRegistry[CommandFactory] = &registry.ImmutableRegistry[CommandFactory]{}
