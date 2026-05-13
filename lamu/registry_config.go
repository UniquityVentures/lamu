package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
)

type Config interface {
	PostConfig()
}

var RegistryConfig *registry.ImmutableRegistry[Config] = &registry.ImmutableRegistry[Config]{}
