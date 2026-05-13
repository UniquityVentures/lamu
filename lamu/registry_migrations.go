package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
)

var RegistryMigrations *registry.ImmutableRegistry[UsefulFilesystem] = &registry.ImmutableRegistry[UsefulFilesystem]{}
