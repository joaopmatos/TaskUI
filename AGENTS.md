# AGENTS.md

Guidance for AI agents working in this repository.

## Agent guidelines

All agents must follow [AGENT_RULES.md](AGENT_RULES.md): mentor-not-ghostwriter, minimal examples with explanations, careful error handling, hint-first debugging, MUV architecture, testability, and security constraints.

## What this project is

TaskUI is a Go TUI and CLI for managing Windows Task Scheduler tasks. It targets
Windows only and talks to the Task Scheduler through COM (`ITaskService`), not
by parsing console output. The development environment is Linux (Fedora), so
cross-compilation matters.

Read [docs/Specification.md](docs/Specification.md) for scope and design.
[README.md](README.md) summarizes the same design.

## Current state

The repository is at the scaffolding stage. `main.go` currently only defines the
`Task` model, the `TaskState` enum, and its `String` method; `main()` is empty.
No CLI, TUI, scheduler engine, or Windows backend exists yet. The commands shown
in the README are the target interface, not working functionality.

Planned layout (one binary, two front-ends over one engine):

- `internal/scheduler` — `Task` model, `Scheduler` interface, in-memory fake,
  and the Windows COM backend.
- `main.go` — entry point; no args launches the TUI, subcommands run the CLI.

Keep domain logic in the engine. Front-ends must not contain scheduling logic,
and both front-ends must go through the same `Scheduler` interface so the engine
can be tested without a terminal or Windows.

## Commands

```sh
gofmt -l .                 # formatting check (should print nothing)
go build ./...
go vet ./...
go test ./...
go run .                   # entry point (currently does nothing)
CGO_ENABLED=0 GOOS=windows go build -o taskui.exe .
```

Cross-compiling for Windows **must** use `CGO_ENABLED=0`; the planned COM
library (`github.com/go-ole/go-ole`) is pure Go and depends on this.

There is no Makefile, lint config, or CI. Add tests as `_test.go` files next to
the code they test, using the standard library `testing` package and table-driven
tests where it fits. Tests must pass with `CGO_ENABLED=0`.

## Conventions

- Standard Go style; run `gofmt` before finishing.
- Use build tags for platform code: Windows backend under `//go:build windows`,
  with a non-Windows test fake available elsewhere.
- Terminal UI work uses Bubble Tea v2 (`charm.land/bubbletea/v2`) and Bubbles v2
  (`charm.land/bubbles/v2`). Bubbles is declared in `go.mod` but not imported by
  any code yet, so it is marked `// indirect`.
- Keep the README and specification accurate when behavior or layout changes.
- Documentation and code identifiers are in English.

## Known issues / gotchas

- `go.mod` includes Bubble Tea, Bubbles, and `github.com/go-ole/go-ole`.
- `go.mod` declares `go 1.26.8`; the local toolchain is newer. Keep both current.
- A prebuilt `TaskUI` binary exists at the repository root. It is ignored by
  `.gitignore` (`TaskUI*`) and untracked; do not commit binaries.
