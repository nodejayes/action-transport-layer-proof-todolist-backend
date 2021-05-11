package main

import (
	"github.com/nodejayes/action-transport-layer-proof-todolist-backend/todo_actions"
	_ "github.com/nodejayes/action-transport-layer-proof-todolist-backend/todo_actions"
	atl "github.com/nodejayes/action-transport-layer/pkg"
)

func readMappings() map[string]func(cl *atl.Client, payload map[string]interface{}) error {
	tmp := make(map[string]func(cl *atl.Client, payload map[string]interface{}) error)
	for k, v := range todo_actions.GetMapping() {
		tmp[k] = v
	}
	return tmp
}

func main() {
	server := atl.NewActionEndpoint(atl.ActionEndpointConfig{
		Address: "localhost:3001",
		Actions: readMappings(),
	})
	server.Start()
}
