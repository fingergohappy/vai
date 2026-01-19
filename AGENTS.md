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

## Testing Notes (tmux)

This is a TUI program; when you add integration/e2e tests that require a real terminal multiplexer,
use `tmux` to create an isolated session for the test, then delete it when the test finishes.

Guidelines:
- Use a unique session name per test run (include PID + timestamp).
- Always cleanup even on failure (Go: `t.Cleanup(...)`; shell: `trap ... EXIT`).
- Prefer `tmux has-session -t <name>` to detect and avoid collisions.

Example (shell):

```bash
SESSION="vai-test-$$-$(date +%s)"
trap 'tmux kill-session -t "$SESSION" 2>/dev/null || true' EXIT

tmux new-session -d -s "$SESSION" "./build/vai"
# ... assertions / captures ...
```

Example (Go test helper):

```go
// Create session, ensure cleanup.
cmd := exec.Command("tmux", "new-session", "-d", "-s", sessionName, "./build/vai")
if err := cmd.Run(); err != nil { t.Fatal(err) }

t.Cleanup(func() {
	_ = exec.Command("tmux", "kill-session", "-t", sessionName).Run()
})
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

## TUI Smoke Check via tmux (agent workflow)


当修改完关于ui方面的代码,检查task,如果task中设计ui的改动,使用以下方式测试,ui是否符合task改动的预期
When someone says the UI "looks wrong", reproduce the screen inside tmux and capture the pane output.
This is intentionally non-interactive: start the program, wait briefly, then `capture-pane`.

Prereqs:
- `tmux` installed
- A tmux session named `vai` exists (create it manually if needed)

Commands (run from repo root):

```bash
# Build first
make build

# Pick window 1 pane id
PANE=$(tmux list-panes -t vai:1 -F '#{pane_id}' | head -n 1)

# Stop whatever is currently running in that pane (best-effort)
tmux send-keys -t "$PANE" C-c

# Start the TUI
tmux send-keys -t "$PANE" "cd /Users/finger/code/my-code/vai && ./build/vai" C-m

# Allow the UI to paint
sleep 1.2

# Capture visible output (use -J to join wrapped lines)
tmux capture-pane -pJ -t "$PANE" -S -200
```

Notes:
- For Bubble Tea apps using alt-screen, capture while the program is running.
- Prefer using `#{pane_id}` rather than window indices to avoid off-by-one issues.

## AI/Tooling Rules

- If the request mentions "proposal", "spec", "change", or "plan": read `openspec/AGENTS.md` first.
- Cursor rules: none found in `.cursor/rules/` or `.cursorrules`.
- Copilot rules: none found in `.github/copilot-instructions.md`.
