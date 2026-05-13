package p_users

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pageEntriesSelection() []registry.Pair[string, components.PageInterface] {
	return []registry.Pair[string, components.PageInterface]{
		{Key: "p_users.UserSelectionTable", Value: &components.Modal{
			UID: "user-selection-modal",
			Children: []components.PageInterface{
				&components.DataTable[User]{
					UID:     "user-selection-table",
					Title:   "Select User",
					Data:    getters.Key[components.ObjectList[User]]("users"),
					RowAttr: getters.RowAttrSelect("UserID", getters.Key[uint]("$row.ID"), getters.Key[string]("$row.Name")),
					Actions: []components.PageInterface{
						&components.TableButtonFilter{Child: lamu.DynamicPage{Name: "p_users.UserSelectionFilter"}},
						&components.ButtonModalForm{
							Name:        getters.Static("p_users.UserCreateForm"),
							Url:         lamu.RoutePath("p_users.CreateRoute", nil),
							FormPostURL: lamu.RoutePath("p_users.CreateRoute", nil),
							ModalUID:    "user-create-modal",
							Icon:        "plus",
							Classes:     "btn-square btn-outline btn-sm",
							Attr:        getters.ModalRefreshList(getters.Static(""), getters.Static("#user-selection-table")),
						},
					},
					Columns: []components.TableColumn{
						{Label: "Name", Name: "Name", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Name")},
						}},
						{Label: "Email", Name: "Email", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Email")},
						}},
						{Label: "Phone", Name: "Phone", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Phone")},
						}},
					},
				},
			},
		}},
		{Key: "p_users.RoleSelectionTable", Value: &components.Modal{
			UID: "role-selection-modal",
			Children: []components.PageInterface{
				&components.DataTable[Role]{
					UID:     "role-selection-table",
					Title:   "Select Role",
					Data:    getters.Key[components.ObjectList[Role]]("roles"),
					RowAttr: getters.RowAttrSelect("RoleID", getters.Key[uint]("$row.ID"), getters.Key[string]("$row.Name")),
					Actions: []components.PageInterface{
						&components.TableButtonFilter{Child: lamu.DynamicPage{Name: "p_users.RoleSelectionFilter"}},
						&components.ButtonModalForm{
							Name:        getters.Static("p_users.RoleCreateForm"),
							Url:         lamu.RoutePath("p_users.RoleCreateRoute", nil),
							FormPostURL: lamu.RoutePath("p_users.RoleCreateRoute", nil),
							ModalUID:    "role-create-modal",
							Icon:        "plus",
							Classes:     "btn-square btn-outline btn-sm",
							Attr:        getters.ModalRefreshList(getters.Static(""), getters.Static("#role-selection-table")),
						},
					},
					Columns: []components.TableColumn{
						{Label: "Name", Name: "Name", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Name")},
						}},
					},
				},
			},
		}},
	}
}
