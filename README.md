# 🐹 Get Set Go

> A personal roadmap for learning Go — from zero to backend engineer & native Linux software developer.

---

## Why Go?

- **Single binary deployment** — no node_modules, no runtime, no 3GB Docker setups. Just one file that runs anywhere.
- **Built-in concurrency** — goroutines and channels make parallel programming intuitive, not painful.
- **Blazing fast** — compiles to native code in seconds.
- **Linux-native** — the language of Docker, Kubernetes, Terraform, and most modern DevOps tooling.
- **Minimal by design** — small surface area, explicit error handling, readable code.

---

## Roadmap Overview

| Phase | Focus | Duration |
|-------|-------|----------|
| [1 — Fundamentals](#phase-1--go-fundamentals) | Variables, structs, interfaces, error handling | 2–3 weeks |
| [2 — Concurrency & stdlib](#phase-2--concurrency--standard-library) | Goroutines, channels, net/http | 2–3 weeks |
| [3 — Backend Development](#phase-3--backend-development) | REST APIs, databases, auth | 4–6 weeks |
| [4 — Linux Systems](#phase-4--linux-systems-programming) | syscall, os/exec, signals, /proc | 3–4 weeks |
| [5 — CLI & TUI Apps](#phase-5--cli-tooling--tui-apps) | cobra, Bubbletea, GoReleaser | 3–4 weeks |
| [6 — Advanced Backend](#phase-6--advanced-backend--real-world-patterns) | WebSockets, gRPC, Redis, Docker | 4–6 weeks |
| [7 — Beast Mode](#phase-7--beast-mode) | Build something real & ship it | 6-8 weeks  |

---

## Phase 1 — Go Fundamentals

**Goal:** Get comfortable with the language and its opinions.

**Topics:**
- Variables, types, constants
- Functions, closures, methods
- Structs and interfaces
- Pointers
- Arrays, slices, maps
- Error handling (the Go way)
- Packages and modules (`go mod`)
- `defer`, `panic`, `recover`

**Projects:**
- `todo-cli` — File I/O, structs, `os.Args`, basic flags
- `unit-converter` — Functions, switch cases, packages
- `number-guessing-game` — Loops, random, user input, error handling

> **Tip:** Embrace `gofmt`. Don't fight Go's style — it exists for a reason and makes reading other people's code effortless.

---

## Phase 2 — Concurrency & Standard Library

**Goal:** Understand Go's superpower — goroutines and channels.

**Topics:**
- Goroutines and the Go scheduler
- Channels (buffered and unbuffered)
- `select` statement
- `sync.WaitGroup` and `sync.Mutex`
- `context` package
- `net/http` basics
- `encoding/json`
- `os`, `io`, `bufio`

**Projects:**
- 🔥 `port-scanner` — Goroutines + channels. Scan 10,000 ports in milliseconds. This is the "aha" moment.
- `weather-cli` — HTTP client, JSON parsing, API calls, `flags` package
- `concurrent-downloader` — Multiple goroutines, `WaitGroup`, progress tracking

> **Tip:** Build the port scanner before moving on. It will change how you think about concurrency forever.

---

## Phase 3 — Backend Development

**Goal:** Build real REST APIs and understand production backend patterns.

**Topics:**
- `net/http` deep dive (routing, middleware, handlers)
- Routing with Chi or stdlib mux
- Middleware chains
- `database/sql` + `sqlx`
- PostgreSQL and SQLite
- JWT authentication
- Table-driven tests
- Migrations with `goose`

**Projects:**
- `rest-api-stdlib` — Pure `net/http`, CRUD, JSON, proper error handling
- `url-shortener` — Database, migrations, redirect logic, basic analytics
- `auth-service` — JWT + refresh tokens, middleware, hashing, user sessions

> **Tip:** Start with the standard library only. Resist Gin and Echo until you feel the pain. You'll understand frameworks much better when you've seen what they save you from.

---

## Phase 4 — Linux Systems Programming

**Goal:** Use Go to interact directly with the Linux kernel and filesystem.

**Topics:**
- `syscall` package and `golang.org/x/sys/unix`
- `os/exec` — running and piping commands
- File system operations at depth
- Signal handling (`SIGTERM`, `SIGINT`, etc.)
- `stdin`/`stdout`/`stderr` pipes
- `fsnotify` for file watching
- Cross-compilation (`GOOS=linux GOARCH=amd64`)
- Embedding files with `//go:embed`

**Projects:**
- `my-grep` / `my-ls` — Flags, `os.Walk`, regex, stdin piping
- `file-organizer` — `fsnotify`, goroutines, auto-sorting on file change
- `system-monitor` — Read `/proc`, goroutines for live updates, terminal rendering

> **Tip:** Reading `/proc/meminfo` in Go with just `os.Open` feels like real systems hacking. Linux exposes everything as files and Go handles files beautifully.

---

## Phase 5 — CLI Tooling & TUI Apps

**Goal:** Build professional-grade terminal applications and ship them.

**Topics:**
- `cobra` for CLI framework
- `viper` for configuration
- `Bubbletea` for terminal UIs
- `Lipgloss` for TUI styling
- Terminal raw mode
- ANSI escape codes
- Shell completion generation
- Releasing with `GoReleaser`

**Projects:**
- 🔥 `tui-dashboard` — Full terminal app with panes, keybindings, live data (Bubbletea)
- `git-helper` — Subcommands, config file, shell completions, man page (cobra)
- `publish-a-tool` — GoReleaser + GitHub Actions → `.deb` binary, brew tap. Actually ship it.

> **Tip:** Bubbletea (by Charm) is what Docker and Kubernetes CLI teams use. It's the industry standard for Go TUIs — learn it and you can build anything terminal-based.

---

## Phase 6 — Advanced Backend & Real-World Patterns

**Goal:** Write production-grade Go services.

**Topics:**
- WebSockets (`gorilla/websocket` or `nhooyr.io/websocket`)
- gRPC + Protocol Buffers
- Message queues (NATS, Kafka basics)
- Redis caching
- Multi-stage Docker builds with `scratch` images
- Structured logging with `slog`
- Profiling with `pprof`
- OpenTelemetry tracing

**Projects:**
- 🔥 `realtime-chat` — WebSockets, goroutines per connection, broadcast hub, rooms
- `grpc-microservice` — Proto definitions, server + client, service discovery basics
- `deploy-go-api` — Scratch Docker image (<10MB), systemd service, nginx reverse proxy

> **Tip:** Deploy the chat app to a VPS with a 5MB Docker image and watch it handle 10,000 concurrent connections. That's the moment you become a Go believer for life.

---

## Phase 7 — Beast Mode

**Goal:** Build something real and complex. No tutorials — just you and the Go docs.

**Choose your endgame:**

- `mini-container` — Linux namespaces, cgroups, syscalls. Like a tiny Docker.
- `tcp-http-server` — Raw `net.Listen`, parse HTTP/1.1 manually, implement keep-alive.
- `kv-database` — File persistence, WAL logging, basic query language.
- **Contribute to a real Go project** — Hugo, Caddy, k9s, fzf, Bubbletea — all open source, all in Go.

---

## Resources

### Official
- [go.dev/tour](https://go.dev/tour) — The official interactive tour. Start here.
- [go.dev/doc/effective_go](https://go.dev/doc/effective_go) — Read this after the tour.
- [pkg.go.dev](https://pkg.go.dev) — The standard library docs. Your bible.

### Books
- *The Go Programming Language* — Donovan & Kernighan (the definitive book)
- *Let's Go* — Alex Edwards (best practical web dev book for Go)
- *Let's Go Further* — Alex Edwards (APIs, auth, production)

### Practice
- [exercism.org/tracks/go](https://exercism.org/tracks/go) — Excellent exercises with mentorship
- [gobyexample.com](https://gobyexample.com) — Annotated examples for every concept

### Community
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com)
- [golangweekly.com](https://golangweekly.com) — Weekly newsletter

---

## Repo Structure

```
get-set-go/
├── README.md
├── phase-1-fundamentals/
│   ├── todo-cli/
│   ├── unit-converter/
│   └── number-guessing-game/
├── phase-2-concurrency/
│   ├── port-scanner/
│   ├── weather-cli/
│   └── concurrent-downloader/
├── phase-3-backend/
│   ├── rest-api-stdlib/
│   ├── url-shortener/
│   └── auth-service/
├── phase-4-linux/
│   ├── my-grep/
│   ├── file-organizer/
│   └── system-monitor/
├── phase-5-cli-tui/
│   ├── tui-dashboard/
│   └── git-helper/
├── phase-6-advanced/
│   ├── realtime-chat/
│   └── grpc-microservice/
└── phase-7-beast-mode/
    └── (my magnum opus)
```

---

## Progress Tracker

- [ ] Phase 1 — Fundamentals
- [ ] Phase 2 — Concurrency & stdlib
- [ ] Phase 3 — Backend development
- [ ] Phase 4 — Linux systems
- [ ] Phase 5 — CLI & TUI apps
- [ ] Phase 6 — Advanced backend
- [ ] Phase 7 — Beast mode 🐹

---

*Built with Go. Documented with love. Started because node_modules got too heavy.*
