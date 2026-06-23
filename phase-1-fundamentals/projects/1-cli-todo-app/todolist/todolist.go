package todolist

import (
	"cli-todo-app/todo"
)

type TodoList struct {
	Counter int                    `json:"counter"`
	ListId  int                    `json:"list_id"`
	Name    string                 `json:"name"`
	Todos   map[int]*todo.TodoItem `json:"todos"`
}
