package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"cli-todo-app/manager"
	"cli-todo-app/todo"
)

// Build Command: go build -o mytodolist .
// Usage: ./mytodolist [args]

// ---- Subcommand registry ----
type cmdEntry struct {
	name    string
	usage   string
	flags   string
	example string
}

var commands = []cmdEntry{
	{
		name:    "add-list",
		usage:   "Create a new todo list",
		flags:   "--name <listname>",
		example: "todo add-list --name work",
	},
	{
		name:    "delete-list",
		usage:   "Delete an existing todo list",
		flags:   "--name <listname>",
		example: "todo delete-list --name work",
	},
	{
		name:    "show-lists",
		usage:   "Show all todo lists with summary counts",
		flags:   "(no flags)",
		example: "todo show-lists",
	},
	{
		name:    "add-task",
		usage:   "Add a task to a list",
		flags:   "--list <listname> --name <taskname> --task <description> --priority <low|medium|high>",
		example: `todo add-task --list work --name "Fix bug" --task "Resolve nil pointer" --priority high`,
	},
	{
		name:    "remove-task",
		usage:   "Remove a task from a list by its ID",
		flags:   "--list <listname> --id <taskid>",
		example: "todo remove-task --list work --id 3",
	},
	{
		name:    "complete-task",
		usage:   "Mark a task as completed",
		flags:   "--list <listname> --id <taskid>",
		example: "todo complete-task --list work --id 3",
	},
	{
		name:    "set-priority",
		usage:   "Change the priority of a task",
		flags:   "--list <listname> --id <taskid> --priority <low|medium|high>",
		example: "todo set-priority --list work --id 3 --priority low",
	},
	{
		name:    "show-tasks",
		usage:   "Show tasks in a list (optionally filter by status or priority)",
		flags:   "--list <listname> [--status <pending|completed|expired>] [--priority <low|medium|high>]",
		example: "todo show-tasks --list work --status pending",
	},
}

func printHelp() {
	fmt.Println(`
cli-todo-app — a simple command-line todo manager

USAGE:
  todo <command> [flags]
  todo --help

CONDITIONS / RULES:
  • List names must be non-empty strings (spaces not recommended).
  • Task --name and --task must both be non-empty.
  • --priority must be one of: low, medium, high  (default: medium for add-task).
  • --status  must be one of: pending, completed, expired.
  • --id must be a positive integer (shown in show-tasks output).
  • A task cannot be completed if it is already completed or expired.
  • You cannot create a list that already exists.
  • --status and --priority filters in show-tasks are mutually exclusive.

COMMANDS:`)

	for _, c := range commands {
		fmt.Printf("\n  %-14s  %s\n", c.name, c.usage)
		fmt.Printf("  %-14s  Flags   : %s\n", "", c.flags)
		fmt.Printf("  %-14s  Example : %s\n", "", c.example)
	}
	fmt.Println()
}

func requireFlag(val, name, cmd string) {
	if val == "" {
		fmt.Fprintf(os.Stderr, "error: --%s is required for '%s'\n", name, cmd)
		os.Exit(1)
	}
}

func requireID(id int, cmd string) {
	if id <= 0 {
		fmt.Fprintf(os.Stderr, "error: --id must be a positive integer for '%s'\n", cmd)
		os.Exit(1)
	}
}

func parsePriority(s, cmd string) todo.TodoPriorityType {
	p, ok := todo.PriorityFromString[strings.ToLower(s)]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: --priority must be low, medium, or high (got %q) for '%s'\n", s, cmd)
		os.Exit(1)
	}
	return p
}

func parseStatus(s, cmd string) todo.TodoStatusType {
	switch strings.ToLower(s) {
	case "pending":
		return todo.PENDING
	case "completed":
		return todo.COMPLETED
	case "expired":
		return todo.EXPIRED
	default:
		fmt.Fprintf(os.Stderr, "error: --status must be pending, completed, or expired (got %q) for '%s'\n", s, cmd)
		os.Exit(1)
	}
	return "" // unreachable
}

// ---- subcommand handlers ----

func handleAddList(args []string) {
	fs := flag.NewFlagSet("add-list", flag.ExitOnError)
	name := fs.String("name", "", "Name of the new list")
	fs.Parse(args)

	requireFlag(*name, "name", "add-list")
	if err := manager.CreateList(*name); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleDeleteList(args []string) {
	fs := flag.NewFlagSet("delete-list", flag.ExitOnError)
	name := fs.String("name", "", "Name of the list to delete")
	fs.Parse(args)

	requireFlag(*name, "name", "delete-list")
	if err := manager.DeleteList(*name); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleShowLists(_ []string) {
	if err := manager.ListAllLists(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleAddTask(args []string) {
	fs := flag.NewFlagSet("add-task", flag.ExitOnError)
	list := fs.String("list", "", "Target list name")
	name := fs.String("name", "", "Short name/title for the task")
	task := fs.String("task", "", "Task description")
	priorityStr := fs.String("priority", "medium", "Priority: low | medium | high")
	fs.Parse(args)

	requireFlag(*list, "list", "add-task")
	requireFlag(*name, "name", "add-task")
	requireFlag(*task, "task", "add-task")
	priority := parsePriority(*priorityStr, "add-task")

	if err := manager.AddTask(*list, *name, *task, priority); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleRemoveTask(args []string) {
	fs := flag.NewFlagSet("remove-task", flag.ExitOnError)
	list := fs.String("list", "", "List name")
	id := fs.Int("id", 0, "Task ID to remove")
	fs.Parse(args)

	requireFlag(*list, "list", "remove-task")
	requireID(*id, "remove-task")

	if err := manager.RemoveTask(*list, *id); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleCompleteTask(args []string) {
	fs := flag.NewFlagSet("complete-task", flag.ExitOnError)
	list := fs.String("list", "", "List name")
	id := fs.Int("id", 0, "Task ID to mark complete")
	fs.Parse(args)

	requireFlag(*list, "list", "complete-task")
	requireID(*id, "complete-task")

	if err := manager.CompleteTask(*list, *id); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleSetPriority(args []string) {
	fs := flag.NewFlagSet("set-priority", flag.ExitOnError)
	list := fs.String("list", "", "List name")
	id := fs.Int("id", 0, "Task ID")
	priorityStr := fs.String("priority", "", "New priority: low | medium | high")
	fs.Parse(args)

	requireFlag(*list, "list", "set-priority")
	requireID(*id, "set-priority")
	requireFlag(*priorityStr, "priority", "set-priority")
	priority := parsePriority(*priorityStr, "set-priority")

	if err := manager.ChangePriority(*list, *id, priority); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func handleShowTasks(args []string) {
	fs := flag.NewFlagSet("show-tasks", flag.ExitOnError)
	list := fs.String("list", "", "List name")
	statusStr := fs.String("status", "", "Filter by status: pending | completed | expired")
	priorityStr := fs.String("priority", "", "Filter by priority: low | medium | high")
	fs.Parse(args)

	requireFlag(*list, "list", "show-tasks")

	if *statusStr != "" && *priorityStr != "" {
		fmt.Fprintln(os.Stderr, "error: --status and --priority are mutually exclusive in 'show-tasks'")
		os.Exit(1)
	}

	var err error
	switch {
	case *statusStr != "":
		status := parseStatus(*statusStr, "show-tasks")
		err = manager.ListByStatus(*list, status)
	case *priorityStr != "":
		priority := parsePriority(*priorityStr, "show-tasks")
		err = manager.ListByPriority(*list, priority)
	default:
		err = manager.ListAllTasks(*list)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// ---- entry point ----

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	// Support bare --help / -help / -h / help as first arg
	switch os.Args[1] {
	case "--help", "-help", "-h", "help":
		printHelp()
		os.Exit(0)
	}

	subcommand := os.Args[1]
	rest := os.Args[2:]

	switch subcommand {
	case "add-list":
		handleAddList(rest)
	case "delete-list":
		handleDeleteList(rest)
	case "show-lists":
		handleShowLists(rest)
	case "add-task":
		handleAddTask(rest)
	case "remove-task":
		handleRemoveTask(rest)
	case "complete-task":
		handleCompleteTask(rest)
	case "set-priority":
		handleSetPriority(rest)
	case "show-tasks":
		handleShowTasks(rest)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %q\nRun 'todo --help' to see available commands.\n", subcommand)
		os.Exit(1)
	}
}
