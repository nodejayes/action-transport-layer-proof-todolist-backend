package store

import (
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/errorhandling"
)

type ToDo struct {
	ID          int    `json:"id" omit:"empty"`
	Finish      bool   `json:"finish"`
	Description string `json:"description"`
}

var store = []ToDo{}

func CreateToDo(item ToDo) (ToDo, error) {
	if item.ID < 1 {
		item.ID = len(store) + 1
	}
	store = append(store, item)
	return item, nil
}

func UpdateToDo(item ToDo) (ToDo, error) {
	for idx := range store {
		if store[idx].ID == item.ID {
			store[idx] = item
		}
	}
	return item, errorhandling.ItemWithIdNotFound("ToDo", item.ID)
}

func DeleteToDo(item ToDo) error {
	if len(store)+1 < item.ID {
		return errorhandling.ItemWithIdNotFound("ToDo", item.ID)
	}
	store = append(store[:item.ID-1], store[item.ID:]...)
	return nil
}

func ReadToDo(id int) (ToDo, error) {
	if len(store)+1 < id {
		return ToDo{}, errorhandling.ItemWithIdNotFound("ToDo", id)
	}
	return store[id-1], nil
}

func ReadToDos(filter func(item ToDo) bool) ([]ToDo, error) {
	tmp := []ToDo{}
	for _, i := range store {
		if filter(i) {
			tmp = append(tmp, i)
		}
	}
	return tmp, nil
}
