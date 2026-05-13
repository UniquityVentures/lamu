package lamu

import (
	"github.com/UniquityVentures/lamu/components"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func init() {
	_ = components.RegistryShellHeadNodes.Register("core.Title", html.TitleEl(gomponents.Text("Lamu")))
}
