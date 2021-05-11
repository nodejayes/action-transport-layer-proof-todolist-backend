package todo_actions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/store"
	atl "github.com/nodejayes/action-transport-layer/pkg"
)

// AddToDoResponse is the Response Format of the AddToDo Action
type AddToDoResponse struct {
	Result store.ToDo `json:"result"`
	Error  error
}

// GetResult returns the Result that was send back to the Client when the Action was triggered
func (r AddToDoResponse) GetResult() interface{} {
	return r.Result
}

// GetError returns the Error String when there was one and this was send back to the Client
func (r AddToDoResponse) GetError() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

// AddToDoHandler is the Action Handler of the AddToDo Action
func AddToDoHandler(cl *atl.Client, payload map[string]interface{}) error {
	var params store.ToDo
	err := mapstructure.Decode(payload, &params)
	if err != nil {
		return err
	}

	item, err := store.CreateToDo(params)
	if err != nil {
		return err
	}
	cl.Send("AddToDo", AddToDoResponse{
		Result: item,
	})
	return nil
}
