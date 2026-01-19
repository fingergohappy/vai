# project-structure Specification

## Purpose
TBD - created by archiving change define-project-structure. Update Purpose after archive.
## Requirements
### Requirement: Standard Go Project Layout

The project MUST follow standard Go project layout conventions.

#### Scenario: Root directory structure

**Given** the vai project root
**When** listing the directory contents
**Then** the following top-level directories MUST exist:

| Directory | Purpose |
|-----------|---------|
| `cmd/` | Application entry points |
| `internal/` | Private application code |
| `pkg/` | Public library code |
| `docs/` | Documentation |
| `scripts/` | Build and install scripts |
| `assets/` | Static assets (help text, etc.) |
| `openspec/` | OpenSpec change proposals |

#### Scenario: Go module files

**Given** the project root
**When** examining Go-related files
**Then** these files MUST exist:
- `go.mod` - Module definition
- `go.sum` - Dependency checksums (after go mod tidy)
- `.gitignore` - Git ignore patterns (including `vai` binary)

---

### Requirement: Internal Package Structure

The project MUST organize internal code by feature domain.

#### Scenario: Internal package directories

**Given** the `internal/` directory
**When** listing its contents
**Then** the following package directories MUST exist:

| Package | Purpose |
|---------|---------|
| `app/` | Top-level Bubble Tea Model and application logic |
| `vim/` | Mode system and key routing |
| `ui/` | Shared UI components (layout, styles, status bar) |
| `chat/` | Chat buffer display and interaction |
| `session/` | Session management and persistence |
| `input/` | Input area handling |
| `clipboard/` | Cross-platform clipboard operations |
| `config/` | Configuration management |

#### Scenario: Package dependencies

**Given** the internal packages
**When** examining import relationships
**Then** the dependency rules MUST be:

- `internal/app` MAY import all other internal packages
- `internal/vim` MAY import `internal/ui`
- `internal/ui` SHOULD NOT import other internal packages (leaf package)
- `internal/chat` MAY import `internal/ui`, `internal/clipboard`
- `internal/session` MAY import `internal/ui`
- `internal/input` MAY import `internal/vim`
- `internal/clipboard` SHOULD NOT import other internal packages (leaf package)
- `internal/config` SHOULD NOT import other internal packages (leaf package)

---

### Requirement: Command Entry Point

The project MUST have a clean entry point in `cmd/vai/`.

#### Scenario: Main package structure

**Given** the `cmd/vai/` directory
**When** examining its contents
**Then** it MUST contain:
- `main.go` - Application entry point with `main()` function

#### Scenario: Main responsibilities

**Given** the `main.go` file
**When** the application starts
**Then** the `main()` function MUST:
- Parse command-line flags
- Load configuration
- Initialize the Bubble Tea program
- Run the program with the top-level Model
- Handle any startup errors gracefully

---

### Requirement: Package File Organization

Each package MUST follow consistent file organization patterns.

#### Scenario: Core package files

**Given** an internal package directory (e.g., `internal/chat/`)
**When** listing its Go files
**Then** it SHOULD contain:

| File | Purpose |
|------|---------|
| `{package}.go` | Core type definitions (e.g., `buffer.go`, `model.go`) |
| `update.go` | Bubble Tea Update logic (if complex) |
| `view.go` | Bubble Tea View logic (if complex) |
| `{feature}.go` | Feature-specific implementations |

#### Scenario: Test file organization

**Given** a package directory
**When** test files are present
**Then** they MUST be named:
- `{package}_test.go` - Main test file
- `{feature}_test.go` - Feature-specific tests
- `example_{name}_test.go` - Example tests

---

### Requirement: Configuration File Locations

The application MUST use standard XDG base directory paths for configuration and data.

#### Scenario: Configuration directory

**Given** the application needs to store configuration
**When** determining the config location
**Then** it SHOULD use:
- `~/.config/vai/config.yaml` on Linux/macOS
- `%APPDATA%\vai\config.yaml` on Windows

#### Scenario: Data directory

**Given** the application needs to store session data
**When** determining the data location
**Then** it SHOULD use:
- `~/.local/share/vai/sessions/` on Linux (XDG_DATA_HOME)
- `~/Library/Application Support/vai/sessions/` on macOS
- `%LOCALAPPDATA%\vai\sessions\` on Windows

---

### Requirement: Build Output

The project MUST define clear build output conventions.

#### Scenario: Binary naming

**Given** the project is built
**When** the build completes
**Then** the output binary SHOULD be named:
- `vai` for the default build
- `vai-{os}-{arch}` for platform-specific builds

#### Scenario: Git ignore patterns

**Given** the `.gitignore` file
**When** it is configured
**Then** it MUST ignore:
- `/vai` - The compiled binary
- `vai-*` - Platform-specific binaries
- `/dist/` - Release packages
- `*.log` - Log files

---

### Requirement: Documentation Structure

The project MUST maintain documentation in the `docs/` directory.

#### Scenario: Documentation files

**Given** the `docs/` directory
**When** listing its contents
**Then** it SHOULD contain:
- `architecture.md` - Architecture overview
- `keybindings.md` - Keybinding reference
- Additional markdown files as needed

---

### Requirement: Script Organization

Build and utility scripts MUST be organized in the `scripts/` directory.

#### Scenario: Build scripts

**Given** the `scripts/` directory
**When** listing its contents
**Then** it SHOULD contain:
- `build.sh` - Build script for local development
- `install.sh` - Install script for system-wide installation
- Additional utility scripts as needed

#### Scenario: Script executability

**Given** a shell script in `scripts/`
**When** the repository is cloned
**Then** the script SHOULD have execute permissions (`chmod +x`)

---

### Requirement: Package Visibility

The project MUST properly separate internal and public APIs.

#### Scenario: Internal package visibility

**Given** a package in `internal/`
**When** it is imported
**Then** it MUST only be importable by code within the module (Go's `internal/` restriction)

#### Scenario: Public package visibility

**Given** a package in `pkg/`
**When** it is imported
**Then** it MAY be imported by external projects
**And** it MUST have a stable API

---

### Requirement: Asset Management

Static assets MUST be properly organized.

#### Scenario: Asset directory

**Given** the project has static assets
**When** they are organized
**Then** they SHOULD be placed in `assets/`

#### Scenario: Help content

**Given** the application displays help text
**When** the help content is stored
**Then** it MAY be:
- Embedded in Go code (using `//go:embed`)
- Stored as text files in `assets/`

---

