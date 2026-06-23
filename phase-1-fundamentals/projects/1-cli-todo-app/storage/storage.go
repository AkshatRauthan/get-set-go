package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"cli-todo-app/todolist"
)

const storageDir = "storage/lists"
 
// ensureStorageDir creates the storage directory if it doesn't exist yet.
// Called before every write — safe to call repeatedly (os.MkdirAll is idempotent).
func ensureStorageDir() error {
	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ensureStorageDir: %w", err)
	}
	return nil
}

func listPath(name string) string {
	return filepath.Join(storageDir, name+".json")
}

func SaveList(list *todolist.TodoList) error {
	err := ensureStorageDir()
	if err != nil {
		return fmt.Errorf("SaveList: %w", err)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return fmt.Errorf("SaveList (marshal): %w", err)
	}

	err = os.WriteFile(listPath(list.Name), data, 0644)
	if err != nil {
		return fmt.Errorf("SaveList (write): %w", err)
	}

	return nil
}

func LoadList(name string) (*todolist.TodoList, error) {
	data, err := os.ReadFile(listPath(name))
	if err != nil {
		return nil, fmt.Errorf("LoadList (read, name=%s): %w", name, err)
	}

	var list todolist.TodoList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, fmt.Errorf("LoadList (unmarshal, name=%s): %w", name, err)
	}

	return &list, nil
}

func LoadAllLists() ([]*todolist.TodoList, error) {
	err := ensureStorageDir()
	if err != nil {
		return nil, fmt.Errorf("LoadAllLists: %w", err)
	}

	entries, err := os.ReadDir(storageDir)
	if err != nil {
		return nil, fmt.Errorf("LoadAllLists (readdir): %w", err)
	}

	lists := make([]*todolist.TodoList, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(storageDir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("LoadAllLists (read, file=%s): %w", entry.Name(), err)
		}

		var list todolist.TodoList
		err = json.Unmarshal(data, &list)
		if err != nil {
			return nil, fmt.Errorf("LoadAllLists (unmarshal, file=%s): %w", entry.Name(), err)
		}

		lists = append(lists, &list)
	}

	return lists, nil
}

func DeleteList(name string) error {
	err := os.Remove(listPath(name))
	if err != nil {
		return fmt.Errorf("DeleteList (name=%s): %w", name, err)
	}
	return nil
}