# TaskUI

A TUI (and CLI) for Windows Task Scheduler.

## Architecture

Two thin front-ends over one UI-agnostic engine:

- `internal/scheduler` — the engine: the `Task` model, the `Scheduler`
  interface, and two backends: an in-memory fake (dev/tests) and the Windows
  backend (the real work).
- `cmd/taskui` — CLI front-end.
- `cmd/taskui-tui` — TUI front-end (Bubble Tea).

Neither front-end contains domain logic; both call the same interface, so the
engine is fully testable without a terminal or a Windows machine.

## Status

- [ ] Engine: `Task` model, `Scheduler` interface, in-memory fake, tests
- [ ] CLI front-end (`taskui list|run|enable|disable|delete`)
- [ ] TUI front-end
- [ ] Windows backend (`//go:build windows`)

## Usage

```sh
go run ./cmd/taskui list
go run ./cmd/taskui run '\MyTasks\DailyBackup'
go run ./cmd/taskui disable '\MyTasks\DailyBackup'
go run ./cmd/taskui delete '\MyTasks\DailyBackup'
go run ./cmd/taskui-tui      # interactive TUI
```

```sh
GOOS=windows go build ./...
```
