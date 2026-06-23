package manager

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"cli-todo-app/storage"
	"cli-todo-app/todo"
	"cli-todo-app/todolist"
)

func CreateList(name string) error {
	_, err := storage.LoadList(name)
	if err == nil {
		return fmt.Errorf("CreateList: list '%s' already exists", name)
	}

	list, err := todolist.CreateList(name, 1)
	if err != nil {
		return fmt.Errorf("CreateList: %w", err)
	}

	err = storage.SaveList(list)
	if err != nil {
		return fmt.Errorf("CreateList: %w", err)
	}

	fmt.Printf("List '%s' created.\n", name)
	return nil
}

func DeleteList(name string) error {
	err := storage.DeleteList(name)
	if err != nil {
		return fmt.Errorf("DeleteList: %w", err)
	}
	fmt.Printf("List '%s' deleted.\n", name)
	return nil
}

// ---- Task Operations ----
// Every task operation follows the same pattern:
// 1. Load list from disk
// 2. Mutate in memory
// 3. Save back to disk
// This ensures every change is persisted immediately.

func AddTask(listName string, name string, task string, priority todo.TodoPriorityType) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("AddTask: %w", err)
	}

	err = list.AddTask(name, task, priority)
	if err != nil {
		return fmt.Errorf("AddTask: %w", err)
	}

	err = storage.SaveList(list)
	if err != nil {
		return fmt.Errorf("AddTask (save): %w", err)
	}

	fmt.Printf("Task '%s' added to list '%s' with priority %s.\n",
		name, listName, todo.PriorityLabel[priority])
	return nil
}

func RemoveTask(listName string, taskId int) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("RemoveTask: %w", err)
	}

	err = list.RemoveTask(taskId)
	if err != nil {
		return fmt.Errorf("RemoveTask: %w", err)
	}

	err = storage.SaveList(list)
	if err != nil {
		return fmt.Errorf("RemoveTask (save): %w", err)
	}

	fmt.Printf("Task #%d removed from list '%s'.\n", taskId, listName)
	return nil
}

func CompleteTask(listName string, taskId int) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("CompleteTask: %w", err)
	}

	err = list.CompleteTask(taskId)
	if err != nil {
		// Use errors.Is to give a more specific message for already-completed todos
		if errors.Is(err, todo.ErrAlreadyCompleted) {
			return fmt.Errorf("CompleteTask: task #%d is already completed", taskId)
		}
		return fmt.Errorf("CompleteTask: %w", err)
	}

	err = storage.SaveList(list)
	if err != nil {
		return fmt.Errorf("CompleteTask (save): %w", err)
	}

	fmt.Printf("Task #%d marked as completed in list '%s'.\n", taskId, listName)
	return nil
}

func ChangePriority(listName string, taskId int, newPriority todo.TodoPriorityType) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("ChangePriority: %w", err)
	}

	err = list.ChangePriority(taskId, newPriority)
	if err != nil {
		return fmt.Errorf("ChangePriority: %w", err)
	}

	err = storage.SaveList(list)
	if err != nil {
		return fmt.Errorf("ChangePriority (save): %w", err)
	}

	fmt.Printf("Task #%d priority updated to %s.\n", taskId, todo.PriorityLabel[newPriority])
	return nil
}

// ---- Display Functions ----
func ListAllTasks(listName string) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("ListAllTasks: %w", err)
	}

	printTaskTable(list.Name, list.Todos)
	return nil
}

func ListByPriority(listName string, priority todo.TodoPriorityType) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("ListByPriority: %w", err)
	}

	filtered := list.GetByPriority(priority)
	if len(filtered) == 0 {
		fmt.Printf("No %s priority tasks in list '%s'.\n",
			todo.PriorityLabel[priority], listName)
		return nil
	}

	// Convert slice back to map for printTaskTable
	m := make(map[int]*todo.TodoItem)
	for _, item := range filtered {
		m[item.Id] = item
	}
	printTaskTable(fmt.Sprintf("%s (priority: %s)", listName, todo.PriorityLabel[priority]), m)
	return nil
}

func ListByStatus(listName string, status todo.TodoStatusType) error {
	list, err := storage.LoadList(listName)
	if err != nil {
		return fmt.Errorf("ListByStatus: %w", err)
	}

	filtered := list.GetByStatus(status)
	if len(filtered) == 0 {
		fmt.Printf("No %s tasks in list '%s'.\n", status, listName)
		return nil
	}

	m := make(map[int]*todo.TodoItem)
	for _, item := range filtered {
		m[item.Id] = item
	}
	printTaskTable(fmt.Sprintf("%s (status: %s)", listName, status), m)
	return nil
}

func ListAllLists() error {
	lists, err := storage.LoadAllLists()
	if err != nil {
		return fmt.Errorf("ListAllLists: %w", err)
	}

	if len(lists) == 0 {
		fmt.Println("No lists found. Create one with: add-list --name <listname>")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "LIST NAME\tTOTAL TASKS\tPENDING\tCOMPLETED\tEXPIRED")
	fmt.Fprintln(w, "---------\t-----------\t-------\t---------\t-------")

	for _, list := range lists {
		pending, completed, expired := 0, 0, 0
		for _, item := range list.Todos {
			switch item.Status {
			case todo.PENDING:
				pending++
			case todo.COMPLETED:
				completed++
			case todo.EXPIRED:
				expired++
			}
		}
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\n",
			list.Name, len(list.Todos), pending, completed, expired)
	}

	return nil
}

// ---- Internal Helpers ----

// printTaskTable renders a map of todos as an aligned table using tabwriter.
// Sorted by ID for consistent output order — maps in Go have no guaranteed order.
func printTaskTable(title string, todos map[int]*todo.TodoItem) {
	if len(todos) == 0 {
		fmt.Printf("List '%s' has no tasks.\n", title)
		return
	}

	// Sort IDs for deterministic output
	ids := make([]int, 0, len(todos))
	for id := range todos {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\nList: %s\n", title)
	fmt.Fprintln(w, "ID\tNAME\tTASK\tPRIORITY\tSTATUS\tCREATED")
	fmt.Fprintln(w, "--\t----\t----\t--------\t------\t-------")

	for _, id := range ids {
		item := todos[id]
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\n",
			item.Id,
			item.Name,
			item.Task,
			todo.PriorityLabel[item.Priority],
			item.Status,
			item.CreatedAt.Format("2006-01-02 15:04"),
		)
	}
}
