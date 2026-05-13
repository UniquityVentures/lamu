package p_users

import (
	"github.com/UniquityVentures/lamu/components"
	"github.com/UniquityVentures/lamu/getters"
	"github.com/UniquityVentures/lamu/lamu"
	"github.com/UniquityVentures/lamu/registry"
)

func pageEntriesTables() []registry.Pair[string, components.PageInterface] {
	return []registry.Pair[string, components.PageInterface]{
		{Key: "users.UserTable", Value: &components.ShellScaffold{
			Sidebar: []components.PageInterface{
				lamu.DynamicPage{Name: "users.UserMenu"},
			},
			Children: []components.PageInterface{
				&components.DataTable[User]{
					UID:     "user-table",
					Classes: "w-full",
					Data:    getters.Key[components.ObjectList[User]]("users"),
					Actions: []components.PageInterface{
						&components.TableButtonFilter{Child: lamu.DynamicPage{Name: "users.UserFilter"}},
						&components.ButtonModalForm{
							Name:        getters.Static("users.UserCreateForm"),
							Url:         lamu.RoutePath("users.CreateRoute", nil),
							FormPostURL: lamu.RoutePath("users.CreateRoute", nil),
							ModalUID:    "user-create-modal",
							Icon:        "plus",
							Classes:     "btn-square btn-outline btn-sm",
						},
					},
					RowAttr: getters.RowAttrNavigate(lamu.RoutePath("users.DetailRoute", map[string]getters.Getter[any]{"id": getters.Any(getters.Key[uint]("$row.ID"))})),
					Columns: []components.TableColumn{
						{Label: "Name", Name: "Name", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Name")},
						}},
						{Label: "Email", Name: "Email", Children: []components.PageInterface{
							&components.FieldText{Getter: getters.Key[string]("$row.Email")},
						}},
						{Label: "Phone", Name: "Phone", Children: []components.PageInterface{
							&components.FieldPhone{Getter: getters.Key[string]("$row.Phone")},
						}},
					},
				},
			},
		}},
	}
}
