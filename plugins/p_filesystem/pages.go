package p_filesystem

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pluginPages() lamu.PluginFeatures[components.PageInterface] {
	var entries []registry.Pair[string, components.PageInterface]
	entries = append(entries, pageEntriesMenus()...)
	entries = append(entries, pageEntriesFilters()...)
	entries = append(entries, pageEntriesTables()...)
	entries = append(entries, pageEntriesDetail()...)
	entries = append(entries, pageEntriesForms()...)
	entries = append(entries, pageEntriesSelection()...)
	return lamu.PluginFeatures[components.PageInterface]{Entries: entries}
}
