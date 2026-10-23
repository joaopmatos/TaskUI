# TaskUI

A TUI (and CLI) for Windows Task Scheduler.

## Architecture

One `taskui` binary with two front-ends over a single UI-agnostic engine:

- `internal/scheduler` — the engine: the `Task` model, the `Scheduler`
  interface, and two backends: an in-memory fake (dev/tests) and the Windows
  backend (the real work).
- `main.go` — the entry point. With no arguments it launches the TUI; with a
  subcommand (`list`, `run`, `enable`, `disable`, `delete`) it behaves as a CLI.

Neither front-end contains domain logic; both call the same interface, so the
engine is fully testable without a terminal or a Windows machine.

## Status

- [x] `Task` model and `TaskState` enum/stringer (currently in `main.go`)
- [ ] Move the model into `internal/scheduler`
- [ ] `Scheduler` interface, in-memory fake, tests
- [ ] CLI subcommands (`list|run|enable|disable|delete`)
- [ ] TUI front-end (Bubble Tea), launched when no subcommand is given
- [ ] Windows backend using `ITaskService` (`//go:build windows`)

## Usage

```sh
go run . list
go run . run '\MyTasks\DailyBackup'
go run . disable '\MyTasks\DailyBackup'
go run . delete '\MyTasks\DailyBackup'
go run .                 # interactive TUI
```

Cross-compile for Windows (pure-Go OLE requires cgo to be disabled):

```sh
CGO_ENABLED=0 GOOS=windows go build -o taskui.exe .
```

See [docs/Specification.md](docs/Specification.md) for the project scope and design.
