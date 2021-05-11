package todo_actions

import (
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/store"
	atl "github.com/nodejayes/action-transport-layer/pkg"
)

type AllToDoResponse struct {
	Result []store.ToDo `json:"result"`
	Error  error
}

func (r AllToDoResponse) GetResult() interface{} {
	return r.Result
}

func (r AllToDoResponse) GetError() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

func AllToDoHandler(cl *atl.Client, payload map[string]interface{}) error {
	items, err := store.ReadToDos(func(item store.ToDo) bool {
		return true
	})
	if err != nil {
		return err
	}

	cl.Send("AllToDo", AllToDoResponse{Result: items})
	return nil
}
