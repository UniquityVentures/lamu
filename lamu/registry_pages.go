package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/registry"
)

var RegistryPage *registry.ImmutableRegistry[components.PageInterface] = &registry.ImmutableRegistry[components.PageInterface]{}
