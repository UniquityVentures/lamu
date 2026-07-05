package p_llm_assistant

import (
	"net/http"

	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/plugins/p_users"
	"golang.org/x/net/websocket"
)

func init() {
	registerPluginRoute("llm_assistant.DefaultRoute", lamu.Route{
		Path:    AppUrl,
		Handler: lamu.NewDynamicView("llm_assistant.ChatView"),
	})

	registerPluginRoute("llm_assistant.HistoryRoute", lamu.Route{
		Path:    AppUrl + "history/",
		Handler: lamu.NewDynamicView("llm_assistant.HistoryView"),
	})

	registerPluginRoute("llm_assistant.ChatSessionRoute", lamu.Route{
		Path:    AppUrl + "c/{id}/",
		Handler: lamu.NewDynamicView("llm_assistant.ChatSessionView"),
	})

	registerPluginRoute("llm_assistant.SidebarChatRoute", lamu.Route{
		Path:    AppUrl + "sidebar-chat/{id}/",
		Handler: lamu.NewDynamicView("llm_assistant.SidebarChatView"),
	})

	registerPluginRoute("llm_assistant.NewSessionRoute", lamu.Route{
		Path:    AppUrl + "new-session/",
		Handler: p_users.RequireAuth(http.HandlerFunc(handleNewSession)),
	})

	registerPluginRoute("llm_assistant.WSRoute", lamu.Route{
		Path: AppUrl + "ws/",
		Handler: p_users.RequireAuth(websocket.Server{
			Handler: assistantWebSocketConn,
		}),
	})

	registerPluginRoute("llm_assistant.SkillsListRoute", lamu.Route{
		Path:    AppUrl + "skills/",
		Handler: lamu.NewDynamicView("llm_assistant.SkillsListView"),
	})

	registerPluginRoute("llm_assistant.SkillsCreateRoute", lamu.Route{
		Path:    AppUrl + "skills/create/",
		Handler: lamu.NewDynamicView("llm_assistant.SkillsCreateView"),
	})

	registerPluginRoute("llm_assistant.SkillsDetailRoute", lamu.Route{
		Path:    AppUrl + "skills/{id}/",
		Handler: lamu.NewDynamicView("llm_assistant.SkillsDetailView"),
	})

	registerPluginRoute("llm_assistant.SkillsUpdateRoute", lamu.Route{
		Path:    AppUrl + "skills/{id}/update/",
		Handler: lamu.NewDynamicView("llm_assistant.SkillsUpdateView"),
	})

	registerPluginRoute("llm_assistant.SkillsDeleteRoute", lamu.Route{
		Path:    AppUrl + "skills/{id}/delete/",
		Handler: lamu.NewDynamicView("llm_assistant.SkillsDeleteView"),
	})

	registerPluginRoute("llm_assistant.SkillsExportRoute", lamu.Route{
		Path:    AppUrl + "skills/{id}/export/",
		Handler: p_users.RequireAuth(http.HandlerFunc(handleSkillExport)),
	})

	registerPluginRoute("llm_assistant.SkillsImportRoute", lamu.Route{
		Path:    AppUrl + "skills/import/",
		Handler: p_users.RequireAuth(http.HandlerFunc(handleSkillImportRoute)),
	})
}
