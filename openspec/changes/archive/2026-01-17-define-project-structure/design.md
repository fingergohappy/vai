# Design: Project Directory Structure

## Overview

This document defines the directory structure for the vai TUI application, following Go standard project layout and Bubble Tea best practices.

## Directory Tree

```
vai/
├── cmd/
│   └── vai/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── app/
│   │   ├── model.go               # Top-level Bubble Tea Model
│   │   ├── update.go              # Update logic and message routing
│   │   ├── view.go                # View rendering logic
│   │   └── init.go                # Initialization logic
│   │
│   ├── vim/
│   │   ├── mode.go                # Mode type and state (NORMAL/INSERT/VISUAL)
│   │   ├── keymap.go              # Keybinding definitions
│   │   └── router.go              # Mode-aware key routing
│   │
│   ├── ui/
│   │   ├── focus.go               # Focus type and management
│   │   ├── layout.go              # Layout calculation (pane dimensions)
│   │   ├── styles.go              # Lipgloss style definitions
│   │   └── statusbar.go           # Status bar component
│   │
│   ├── chat/
│   │   ├── buffer.go              # Chat buffer Model
│   │   ├── message.go             # Message and Block types
│   │   ├── block.go               # Block interface and implementations
│   │   ├── viewport.go            # Viewport rendering
│   │   ├── cursor.go              # Cursor position tracking
│   │   └── markdown.go            # Markdown to block parser
│   │
│   ├── session/
│   │   ├── manager.go             # Session manager Model
│   │   ├── session.go             # Session type and storage
│   │   ├── list.go                # Session list component
│   │   └── storage.go             # File I/O operations
│   │
│   ├── input/
│   │   ├── area.go                # Input area Model (wraps bubbles.TextArea)
│   │   ├── vim.go                 # Vim-style movement in INSERT mode
│   │   └── history.go             # Input history (optional)
│   │
│   ├── clipboard/
│   │   ├── clipboard.go           # Clipboard interface
│   │   ├── macos.go               # macOS implementation (pbcopy)
│   │   ├── linux.go               # Linux implementation (xclip/wl-copy)
│   │   └── dummy.go               # Fallback implementation
│   │
│   └── config/
│       ├── config.go              # Configuration type
│       ├── loader.go              # Config file loading
│       └── defaults.go            # Default values
│
├── pkg/
│   └── markdown/
│       ├── parser.go              # Markdown parser (optional library wrapper)
│       └── ast.go                 # AST types
│
├── assets/
│   └── help.txt                   # Help content (optional)
│
├── docs/
│   ├── architecture.md            # Architecture documentation
│   └── keybindings.md             # Keybinding reference
│
├── scripts/
│   ├── build.sh                   # Build script
│   └── install.sh                 # Install script
│
├── .editorconfig                  # Editor configuration
├── .gitignore                     # Git ignore rules
├── go.mod                         # Go module definition
├── go.sum                         # Go checksums
├── LICENSE                        # License file
├── README.md                      # Project README
└── Makefile                       # Build automation
```

## Package Responsibilities

### `cmd/vai/` - Application Entry

**Purpose:** Main application entry point

**Responsibilities:**
- Parse command-line flags
- Initialize Bubble Tea program
- Set up error handling
- Wire dependencies

**Exports:** None (this is the binary)

---

### `internal/app/` - Top-Level Application

**Purpose:** Bubble Tea top-level Model

**Responsibilities:**
- Combine sub-models (History, Buffer, Input)
- Route messages to sub-models
- Aggregate commands from sub-models
- Manage global state (Mode, Focus)

**Key Types:**
- `Model` - Top-level Bubble Tea Model
- `Msg` - Application-specific messages

**Dependencies:** All other internal packages

---

### `internal/vim/` - Mode System

**Purpose:** Vim-style mode management

**Responsibilities:**
- Define Mode type (NORMAL, INSERT, VISUAL)
- Define keybinding mappings
- Route keys based on mode and focus
- Handle mode transitions

**Key Types:**
- `Mode` - Mode enum
- `Keymap` - Keybinding configuration
- `Router` - Key routing logic

**Dependencies:** `internal/ui` (for Focus)

---

### `internal/ui/` - UI Components

**Purpose:** Shared UI utilities and components

**Responsibilities:**
- Define Focus type and management
- Calculate layout dimensions
- Define Lipgloss styles
- Render status bar

**Key Types:**
- `Focus` - Focus enum (History, Buffer, Input)
- `Layout` - Pane dimensions
- `Styles` - Style definitions
- `StatusBar` - Status bar component

**Dependencies:** `charmbracelet/lipgloss`

---

### `internal/chat/` - Chat Buffer

**Purpose:** Chat content display and interaction

**Responsibilities:**
- Render messages (text and code blocks)
- Handle viewport scrolling
- Track cursor position
- Support code block navigation
- Support VISUAL mode selection

**Key Types:**
- `Model` - Chat buffer Bubble Tea Model
- `Message` - Message type
- `Block` - Block interface
- `TextBlock`, `CodeBlock` - Block implementations
- `Viewport` - Viewport state
- `Cursor` - Cursor position
- `Selection` - VISUAL mode selection

**Dependencies:** `internal/ui`, `internal/clipboard`, `charmbracelet/bubbles/viewport`

---

### `internal/session/` - Session Management

**Purpose:** Session persistence and list management

**Responsibilities:**
- Manage session data
- Load/save sessions to disk
- Render session list
- Handle session operations (create, delete, rename, switch)

**Key Types:**
- `Model` - Session manager Bubble Tea Model
- `Session` - Session type
- `List` - Session list component
- `Storage` - File I/O operations

**Dependencies:** `internal/ui`

---

### `internal/input/` - Input Area

**Purpose:** User input handling

**Responsibilities:**
- Wrap bubbles.TextArea
- Handle Vim-style movement in INSERT mode
- Send messages on Enter
- Maintain input state

**Key Types:**
- `Model` - Input area Bubble Tea Model
- `VimMotion` - Vim movement handler

**Dependencies:** `internal/vim`, `charmbracelet/bubbles/textarea`

---

### `internal/clipboard/` - Clipboard Operations

**Purpose:** Cross-platform clipboard access

**Responsibilities:**
- Provide clipboard interface
- Platform-specific implementations
- Error handling when unavailable

**Key Types:**
- `Clipboard` - Clipboard interface
- Platform implementations (macOS, Linux)

**Dependencies:** `os/exec`

---

### `internal/config/` - Configuration

**Purpose:** Application configuration

**Responsibilities:**
- Define configuration structure
- Load config from file
- Provide default values
- Validate configuration

**Key Types:**
- `Config` - Configuration type
- `Loader` - Config loading logic

**Dependencies:** None

---

### `pkg/markdown/` - Markdown Parser (Optional)

**Purpose:** Parse markdown into structured blocks

**Responsibilities:**
- Parse markdown text
- Extract code blocks
- Return structured AST

**Key Types:**
- `Parser` - Parser interface
- `AST` - Abstract syntax tree

**Dependencies:** External markdown library (if used)

---

## Design Principles

### 1. Standard Go Layout

Follow [golang-standards/project-layout](https://github.com/golang-standards/project-layout):
- `cmd/` - Main applications
- `internal/` - Private application code
- `pkg/` - Public library code
- `docs/` - Documentation

### 2. Package by Feature

Organize by feature domain, not by layer:
- ✅ `internal/chat/` (feature)
- ❌ `internal/model/`, `internal/view/` (layer)

### 3. Minimal Dependencies

Each package should have minimal dependencies:
- Leaf packages (config, clipboard) have no internal dependencies
- Core packages (app) depend on many but are dependency-free for others

### 4. Testability

Structure for easy testing:
- Clear interfaces between packages
- Dependency injection through constructors
- Mock-friendly design

### 5. Bubble Tea Patterns

Follow Bubble Tea conventions:
- One Model per major component
- Clear Msg types for communication
- Init/Update/View separation

---

## Import Graph

```
                    ┌─────────────┐
                    │   cmd/vai   │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │ internal/app │ ◄──── Top-level Model
                    └──────┬──────┘
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
┌───────▼───────┐  ┌───────▼───────┐  ┌──────▼──────┐
│ internal/vim  │  │  internal/ui  │  │internal/chat │
└───────────────┘  └───────┬───────┘  └──────┬──────┘
                           │                  │
                    ┌──────▼───────┐          │
                    │internal/input│          │
                    └──────┬───────┘          │
                           │                  │
        ┌──────────────────┼──────────────────┼──────────┐
        │                  │                  │          │
┌───────▼────────┐ ┌──────▼──────┐  ┌───────▼─────┐ ┌▼──────────┐
│internal/session│ │internal/config│ │internal/clip│ │pkg/markdown│
└────────────────┘ └─────────────┘  └─────────────┘ └───────────┘
```

---

## File Naming Conventions

### Go Files

- `{package}.go` - Main package file (e.g., `model.go`)
- `{feature}.go` - Feature-specific file (e.g., `vim.go`)
- `{type}_test.go` - Tests for a type
- `example_{name}_test.go` - Example tests

### Documentation Files

- `{package}.md` - Package documentation (optional)
- `README.md` - Project overview
- `DESIGN.md` - Design decisions (already exists)

---

## Build Artifacts

```
vai/
├── vai                    # Compiled binary (gitignored)
├── vai-*                  # Platform-specific binaries
│   ├── vai-linux-amd64
│   ├── vai-darwin-arm64
│   └── vai-windows-amd64.exe
└── dist/                  # Release packages (gitignored)
    ├── vai-linux-amd64.tar.gz
    └── vai-darwin-arm64.tar.gz
```

---

## Configuration and Data Files

```
~/.config/vai/
├── config.yaml            # User configuration
└── keybindings.yaml       # Custom keybindings (optional)

~/.local/share/vai/
├── sessions/              # Session data
│   ├── {session-id}.json
│   └── ...
└── debug.log              # Debug logs (if enabled)
```

---

## Next Steps

After this structure is defined:
1. Create all directories
2. Add skeleton Go files with package declarations
3. Define key types and interfaces
4. Add basic tests for package structure
5. Set up build tooling (Makefile, scripts)
