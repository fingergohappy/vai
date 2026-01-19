# Tasks: Define Project Directory Structure

This document outlines the tasks for setting up the project directory structure for the vai TUI application.

## Task Summary

- Create all required directories following Go standard layout
- Add skeleton Go files with package declarations
- Set up configuration files (.gitignore, Makefile, etc.)
- Validate the structure builds correctly

---

## Phase 1: Create Directory Structure

### 1.1 Create root directories
- [x] Create `cmd/` directory
- [x] Create `internal/` directory
- [x] Create `pkg/` directory
- [x] Create `docs/` directory
- [x] Create `scripts/` directory
- [x] Create `assets/` directory

**Validation:** All directories exist with `ls -la`

### 1.2 Create cmd/vai structure
- [x] Create `cmd/vai/` directory
- [x] Create `cmd/vai/main.go` with package main declaration

**Validation:** `ls cmd/vai/` shows main.go

### 1.3 Create internal package directories
- [x] Create `internal/app/`
- [x] Create `internal/vim/`
- [x] Create `internal/ui/`
- [x] Create `internal/chat/`
- [x] Create `internal/session/`
- [x] Create `internal/input/`
- [x] Create `internal/clipboard/`
- [x] Create `internal/config/`

**Validation:** `ls internal/` shows all 8 package directories

### 1.4 Create pkg directory
- [x] Create `pkg/markdown/` (optional, for markdown parsing)

**Validation:** `ls pkg/` shows markdown directory

---

## Phase 2: Create Skeleton Go Files

### 2.1 Create cmd/vai/main.go
- [x] Create `cmd/vai/main.go`
- [x] Add package main declaration
- [x] Add empty main() function
- [x] Add TODO comment for Bubble Tea initialization

**Validation:** `go build ./cmd/vai` produces no syntax errors

### 2.2 Create internal/app skeleton
- [x] Create `internal/app/model.go` with package app
- [x] Define empty Model struct
- [x] Add Mode, Focus fields to Model
- [x] Add placeholders for sub-models

**Validation:** File compiles with `go build ./internal/app`

### 2.3 Create internal/vim skeleton
- [x] Create `internal/vim/mode.go` with package vim
- [x] Define Mode type (NORMAL, INSERT, VISUAL)
- [x] Create `internal/vim/keymap.go` for keybindings
- [x] Create `internal/vim/router.go` for key routing

**Validation:** Files compile with `go build ./internal/vim`

### 2.4 Create internal/ui skeleton
- [x] Create `internal/ui/focus.go` with package ui
- [x] Define Focus type (History, Buffer, Input)
- [x] Create `internal/ui/layout.go` for layout calculations
- [x] Create `internal/ui/styles.go` for Lipgloss styles
- [x] Create `internal/ui/statusbar.go` for status bar

**Validation:** Files compile with `go build ./internal/ui`

### 2.5 Create internal/chat skeleton
- [x] Create `internal/chat/buffer.go` with package chat
- [x] Create `internal/chat/message.go` for Message type
- [x] Create `internal/chat/block.go` for Block interface
- [x] Define TextBlock and CodeBlock structs

**Validation:** Files compile with `go build ./internal/chat`

### 2.6 Create internal/session skeleton
- [x] Create `internal/session/manager.go` with package session
- [x] Create `internal/session/session.go` for Session type
- [x] Create `internal/session/list.go` for session list component

**Validation:** Files compile with `go build ./internal/session`

### 2.7 Create internal/input skeleton
- [x] Create `internal/input/area.go` with package input
- [x] Create `internal/input/vim.go` for INSERT mode Vim movement

**Validation:** Files compile with `go build ./internal/input`

### 2.8 Create internal/clipboard skeleton
- [x] Create `internal/clipboard/clipboard.go` with package clipboard
- [x] Define Clipboard interface
- [x] Create `internal/clipboard/macos.go` for macOS
- [x] Create `internal/clipboard/linux.go` for Linux
- [x] Create `internal/clipboard/dummy.go` as fallback

**Validation:** Files compile with `go build ./internal/clipboard`

### 2.9 Create internal/config skeleton
- [x] Create `internal/config/config.go` with package config
- [x] Define Config struct
- [x] Create `internal/config/loader.go` for config loading
- [x] Create `internal/config/defaults.go` for default values

**Validation:** Files compile with `go build ./internal/config`

### 2.10 Create pkg/markdown skeleton (optional)
- [x] Create `pkg/markdown/parser.go` with package markdown
- [x] Create `pkg/markdown/ast.go` for AST types

**Validation:** Files compile with `go build ./pkg/markdown`

---

## Phase 3: Configuration Files

### 3.1 Create .gitignore
- [x] Create `.gitignore` file
- [x] Add `/vai` binary
- [x] Add `vai-*` platform binaries
- [x] Add `/dist/` directory
- [x] Add `*.log` files
- [x] Add Go-specific ignores (vendor/, .idea/, etc.)

**Validation:** `git status` shows no untracked Go artifacts

### 3.2 Create .editorconfig
- [x] Create `.editorconfig`
- [x] Set indent_style = space
- [x] Set indent_size = 4 for Go
- [x] Set charset = utf-8

**Validation:** Editor respects formatting rules

### 3.3 Create Makefile
- [x] Create `Makefile`
- [x] Add `build` target
- [x] Add `run` target
- [x] Add `test` target
- [x] Add `clean` target
- [x] Add `install` target

**Validation:** `make build` produces vai binary

### 3.4 Create build scripts
- [x] Create `scripts/build.sh`
- [x] Make script executable (`chmod +x`)
- [x] Create `scripts/install.sh`
- [x] Make script executable

**Validation:** `./scripts/build.sh` builds the project

---

## Phase 4: Documentation

### 4.1 Create README.md
- [x] Create `README.md` in project root
- [x] Add project description
- [x] Add installation instructions
- [x] Add usage overview
- [x] Add keybinding summary

**Validation:** README renders correctly in markdown viewer

### 4.2 Create docs/architecture.md
- [x] Create `docs/architecture.md`
- [x] Document the directory structure
- [x] Document package relationships
- [x] Include the import graph

**Validation:** Documentation matches the actual structure

### 4.3 Create docs/keybindings.md
- [x] Create `docs/keybindings.md`
- [x] Document all keybindings by mode
- [x] Include mode-focus compatibility matrix

**Validation:** Keybinding documentation is complete

---

## Phase 5: Dependencies and Build

### 5.1 Initialize Bubble Tea dependencies
- [x] Run `go get github.com/charmbracelet/bubbletea`
- [x] Run `go get github.com/charmbracelet/lipgloss`
- [x] Run `go get github.com/charmbracelet/bubbles`
- [x] Run `go mod tidy`

**Validation:** `go.mod` contains all dependencies

### 5.2 Verify build
- [x] Run `go build ./cmd/vai`
- [x] Verify `vai` binary is created
- [x] Run `./vai` (shows placeholder message)

**Validation:** Binary executes without errors

### 5.3 Initialize git repository
- [x] Run `git init` if not already initialized
- [x] Verify `.gitignore` is working
- [ ] Create initial commit with structure

**Validation:** Git repository is properly set up

---

## Phase 6: Validation

### 6.1 Verify package structure
- [x] Run `go build ./...` to build all packages
- [x] Verify no circular dependencies
- [x] Check that internal packages are not externally visible

**Validation:** All packages build successfully

### 6.2 Verify directory conventions
- [x] Check that all packages follow Go naming conventions
- [x] Verify file naming matches package names
- [x] Ensure no mixed-case directory names

**Validation:** Structure follows Go standards

### 6.3 Verify imports
- [x] Check that `internal/` packages only import within module
- [x] Verify `pkg/` packages have clean APIs
- [x] Ensure no unused imports

**Validation:** `go vet ./...` produces no warnings

---

## Dependencies

- **Phase 1** must be completed before Phase 2 (need directories for files)
- **Phase 2** must be completed before Phase 5 (need files to build)
- **Phase 3** can be done in parallel with Phase 2
- **Phase 4** can be done at any time
- **Phase 5** depends on Phase 2 (need Go files to build)

## Parallelizable Work

These phases can be worked on in parallel:
- **Phase 3** (Configuration) and **Phase 4** (Documentation)
- Different skeleton files in **Phase 2** (each package is independent)

## Estimated Completion

This is a foundational change that should be completed first, as all subsequent implementation work depends on this structure.

**Estimated time:** 2-3 hours for full setup

---

## Definition of Done

This change is complete when:
1. All directories are created following the design document
2. All skeleton Go files exist and compile
3. `.gitignore`, `Makefile`, and other config files are in place
4. `go build ./...` succeeds with no errors
5. `go vet ./...` produces no warnings
6. Initial README and architecture documentation exist
7. Git repository is initialized with the structure
