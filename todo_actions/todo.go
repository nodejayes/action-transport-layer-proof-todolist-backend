package todo_actions

import atl "github.com/nodejayes/action-transport-layer/pkg"

func GetMapping() map[string]func(cl *atl.Client, payload map[string]interface{}) error {
	return map[string]func(cl *atl.Client, payload map[string]interface{}) error{
		"AllToDo":               AllToDoHandler,
		"AddToDo":               AddToDoHandler,
		"UpdateToDoDescription": UpdateToDoDescriptionHandler,
		"ToggleToDo":            ToggleToDoHandler,
		"DeleteToDo":            DeleteToDoHandler,
	}
}
