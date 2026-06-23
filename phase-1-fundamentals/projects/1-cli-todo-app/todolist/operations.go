package todolist

import (
	"fmt"

	"cli-todo-app/todo"
)

func CreateList(name string, id int) (*TodoList, error) {
	if name == "" {
		return nil, fmt.Errorf("CreateList: list name cannot be empty")
	}
	return &TodoList{
		Counter: 1,
		ListId:  id,
		Name:    name,
		Todos:   make(map[int]*todo.TodoItem),
	}, nil
}

func (t *TodoList) AddTask(name string, task string, priority todo.TodoPriorityType) error {
	newTodo, err := todo.CreateTodo(t.Counter, name, task, priority)
	if err != nil {
		// forward — caller (manager) will wrap this with further context
		return fmt.Errorf("AddTask: %w", err)
	}

	t.Todos[t.Counter] = newTodo
	t.Counter++
	return nil
}

func (t *TodoList) RemoveTask(taskId int) error {
	_, exists := t.Todos[taskId]
	if !exists {
		return fmt.Errorf("RemoveTask (id=%d): %w", taskId, todo.ErrTodoNotFound)
	}
	delete(t.Todos, taskId)
	return nil
}

func (t *TodoList) CompleteTask(taskId int) error {
	item, exists := t.Todos[taskId]
	if !exists {
		return fmt.Errorf("CompleteTask (id=%d): %w", taskId, todo.ErrTodoNotFound)
	}

	err := item.MarkCompleted()
	if err != nil {
		return fmt.Errorf("CompleteTask: %w", err)
	}
	return nil
}

func (t *TodoList) ChangePriority(taskId int, newPriority todo.TodoPriorityType) error {
	item, exists := t.Todos[taskId]
	if !exists {
		return fmt.Errorf("ChangePriority (id=%d): %w", taskId, todo.ErrTodoNotFound)
	}

	err := item.ChangePriority(newPriority)
	if err != nil {
		return fmt.Errorf("ChangePriority: %w", err)
	}
	return nil
}

func (t *TodoList) GetByPriority(priority todo.TodoPriorityType) []*todo.TodoItem {
	result := make([]*todo.TodoItem, 0)
	for _, item := range t.Todos {
		if item.Priority == priority {
			result = append(result, item)
		}
	}
	return result
}

func (t *TodoList) GetByStatus(status todo.TodoStatusType) []*todo.TodoItem {
	result := make([]*todo.TodoItem, 0)
	for _, item := range t.Todos {
		if item.Status == status {
			result = append(result, item)
		}
	}
	return result
}
