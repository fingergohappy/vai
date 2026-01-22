<!-- OPENSPEC:START -->
# OpenSpec Instructions

These instructions are for AI assistants working in this project.

Always open `@/openspec/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/openspec/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

Keep this managed block so 'openspec update' can refresh the instructions.

<!-- OPENSPEC:END -->

# Agent Guide (vai)

This repo is a Go TUI app (Bubble Tea). This file is for agentic coding tools.

## Commands

Primary entrypoints:

```bash
# Build (outputs: build/vai)
make build

# Run (builds first)
make run

# Install (to GOPATH/bin or $HOME/go/bin)
make install

# Tests (none exist yet; uses `go test`)
make test

# Coverage
make test-coverage

# Lint (optional; requires golangci-lint)
make lint

# Format and basic static checks
make fmt
make vet

# Dependencies
make deps

# Cleanup
make clean
```

Direct Go equivalents:

```bash
go build -o build/vai ./cmd/vai

go test -v ./...

go fmt ./...
go vet ./...
```

### Run a single test / package

```bash
# Run all tests in one package
go test -v ./internal/chat

# Run a single test function (exact name)
go test -v ./internal/chat -run '^TestName$'

# Run tests matching a pattern
go test -v ./internal/chat -run 'TestModel'

# Avoid cached results while iterating
go test -v ./internal/chat -run '^TestName$' -count=1
```



## Code Style (Go)

Formatting:
- `gofmt` is the source of truth. Run `make fmt` (or `go fmt ./...`) before committing.
- `.editorconfig` exists; if it conflicts with `gofmt`, follow `gofmt` for Go files.

Imports:
- Use standard Go grouping: stdlib, blank line, third-party, blank line, local.
- Aliases only when needed (example already used: `ui ".../internal/ui"`).

- Packages: short, lower-case (`app`, `chat`, `vim`, `ui`, `session`).
- Exported types/functions: PascalCase (`Model`, `NewModel`, `DefaultStyles`).
- Constants: PascalCase with a clear prefix (`ModeNormal`, `RoleAssistant`).

Comments:
- Keep GoDoc on exported symbols (this repo uses `// Package ... provides ...`).
- Keep comments factual; prefer code clarity over commentary.

Error handling:
- Do not ignore errors. Use early returns:
  `if err != nil { return ..., err }`
- Wrap errors with context using `%w` when propagating.
- Avoid `panic` in app logic; `main` can print error and `os.Exit(1)`.

Architecture expectations (Bubble Tea):
- Follow Model-Update-View: keep state in models; route messages predictably.
- Keep `internal/app` as the composition/root router; keep domain logic in feature packages.

## Project Layout (high-level)

- `cmd/vai/`: program entrypoint
- `internal/app/`: top-level Bubble Tea model + routing
- `internal/ui/`: layout/styles/shared UI bits
- `internal/vim/`: mode + key routing
- `internal/chat/`, `internal/session/`, `internal/input/`: feature modules
- `internal/clipboard/`, `internal/config/`: platform + configuration
- `pkg/`: public packages (currently `pkg/markdown`)

# TUI Verification Workflow (tmux-based)

This project is a **TUI (Terminal User Interface) program written in Go**.

Any change that affects **TUI behavior, layout, or interaction** MUST follow this workflow.
The purpose is to ensure that all user-visible effects are verified in a **real terminal environment** using `tmux`.

---

## Core Principles

1. **Plan first**
   - Describe the expected user-visible effect before coding.
2. **Implement**
   - Modify the TUI-related code.
3. **Verify**
   - Run the program inside `tmux`.
4. **Assert**
   - Capture terminal output and confirm it matches the plan.

Skipping any step invalidates the change.

---

## tmux Conventions

| Item    | Value |
|--------|-------|
| Session | `vai` |
| Window  | `vai` |
| Program | `vai` |

All verification MUST happen inside this tmux session and window.

---

## Step 1: Check tmux session

Check whether a tmux session named `vai` exists.

```bash
tmux has-session -t vai 2>/dev/null
```

### If the session does not exist, create it

```bash
tmux new-session -d -s vai
```

---

## Step 2: Check tmux window

Check whether the `vai` session has a window named `vai`.

```bash
tmux list-windows -t vai
```

### If the window does not exist, create it

```bash
tmux new-window -t vai -n vai
```

---

## Step 3: Run the TUI program

The TUI program **MUST** be started from the **current project root directory**  
and **MUST** use `go run`.

Running a precompiled binary is **NOT allowed**.

### 3.1 Enter the project root directory

The project root SHOULD be determined explicitly, for example:

```bash
git rev-parse --show-toplevel
```

Then enter the directory inside tmux:

```bash
tmux send-keys -t vai:vai 'cd /path/to/current/project' C-m
```

> Agents MUST NOT assume a fixed directory path.

---

### 3.2 Start the TUI program using `go run`

If the project root is the entry point:

```bash
tmux send-keys -t vai:vai 'go run .' C-m
```

If the entry point is under `cmd/vai`:

```bash
tmux send-keys -t vai:vai 'go run ./cmd/vai' C-m
```

---

### Enforcement rules (Step 3)

- MUST run from project root
- MUST use `go run`
- MUST reflect latest source code
- MUST NOT execute a prebuilt `vai` binary

---

## Step 4: Capture and verify output

After the TUI program is running, capture the pane output:

```bash
tmux capture-pane -t vai:vai -p
```

### Verification checklist

Captured output MUST match the planned effect:

- Layout is correct
- UI elements appear as expected
- No unexpected errors or visual artifacts
- Interaction flow matches the plan

If any check fails, the change is **invalid** and must be revised.

---

## Mandatory Rules for Agents

- TUI changes MUST be verified using tmux
- Do NOT rely on screenshots, assumptions, or imagination
- Every TUI change MUST have:
  - a plan
  - a tmux run
  - a captured output
  - a verification decision

---

## Example Verification Loop

1. Describe expected UI behavior
2. Modify TUI code
3. Run via tmux using `go run`
4. Capture pane output
5. Compare output with the plan

Only after all steps succeed is the change considered **complete**.
