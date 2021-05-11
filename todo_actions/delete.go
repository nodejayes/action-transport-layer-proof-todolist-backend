package todo_actions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/store"
	atl "github.com/nodejayes/action-transport-layer/pkg"
)

type DeleteToDoPayload struct {
	Item store.ToDo `json:"id"`
}

type DeleteToDoResponse struct {
	Result int `json:"result"`
	Error  error
}

func (r DeleteToDoResponse) GetResult() interface{} {
	return r.Result
}

func (r DeleteToDoResponse) GetError() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

func DeleteToDoHandler(cl *atl.Client, payload map[string]interface{}) error {
	var param DeleteToDoPayload
	err := mapstructure.Decode(payload, &param)
	if err != nil {
		return err
	}

	err = store.DeleteToDo(param.Item)
	if err != nil {
		return err
	}
	cl.Send("DeleteToDo", DeleteToDoResponse{Result: param.Item.ID})
	return nil
}
