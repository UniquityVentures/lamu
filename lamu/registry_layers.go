package lamu

import (
	"github.com/UniquityVentures/lamu/registry"
	"github.com/UniquityVentures/lamu/views"
)

var RegistryLayer *registry.ImmutableRegistry[views.GlobalLayer] = &registry.ImmutableRegistry[views.GlobalLayer]{}
