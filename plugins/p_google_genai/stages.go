package p_google_genai

import "github.com/UniquityVentures/lamu/lamu"

func pluginStages[T any](stage func() lamu.PluginFeatures[T]) []func() lamu.PluginFeatures[T] {
	return []func() lamu.PluginFeatures[T]{stage}
}
