package todo_actions

import (
	"github.com/mitchellh/mapstructure"
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/store"
	atl "github.com/nodejayes/action-transport-layer/pkg"
)

// -------------------------------------------------------------------------------------------------
// Toggle Item
// -------------------------------------------------------------------------------------------------

type ToggleToDoPayload struct {
	ID int `json:"id"`
}

type ToggleToDoResponse struct {
	Result store.ToDo `json:"result"`
	Error  error
}

func (r ToggleToDoResponse) GetResult() interface{} {
	return r.Result
}

func (r ToggleToDoResponse) GetError() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

func ToggleToDoHandler(cl *atl.Client, payload map[string]interface{}) error {
	var param ToggleToDoPayload
	err := mapstructure.Decode(payload, &param)
	if err != nil {
		return err
	}

	item, err := store.ReadToDo(param.ID)
	if err != nil {
		return err
	}
	item.Finish = !item.Finish
	_, err = store.UpdateToDo(item)
	cl.Send("ToggleToDo", ToggleToDoResponse{Result: item})
	return nil
}

// -------------------------------------------------------------------------------------------------
// Update Item Description
// -------------------------------------------------------------------------------------------------

type UpdateToDoDescriptionPayload struct {
	ID          int    `json:"id"`
	Description string `json:"string"`
}

type UpdateToDoDescriptionResponse struct {
	Result store.ToDo `json:"result"`
	Error  error
}

func (r UpdateToDoDescriptionResponse) GetResult() interface{} {
	return r.Result
}

func (r UpdateToDoDescriptionResponse) GetError() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

func UpdateToDoDescriptionHandler(cl *atl.Client, payload map[string]interface{}) error {
	var param UpdateToDoDescriptionPayload
	err := mapstructure.Decode(payload, &param)
	if err != nil {
		cl.Send("UpdateToDoDescription", UpdateToDoDescriptionResponse{Result: store.ToDo{}, Error: err})
		return err
	}

	item, err := store.ReadToDo(param.ID)
	if err != nil {
		cl.Send("UpdateToDoDescription", UpdateToDoDescriptionResponse{Result: item, Error: err})
		return err
	}
	item.Description = param.Description
	_, err = store.UpdateToDo(item)
	if err != nil {
		cl.Send("UpdateToDoDescription", UpdateToDoDescriptionResponse{Result: item, Error: err})
		return err
	}
	cl.Send("UpdateToDoDescription", UpdateToDoDescriptionResponse{Result: item})
	return nil
}
