# Tasks: Implement Vim-style TUI AI Chat Framework

This document outlines the ordered implementation tasks for delivering the TUI framework. Tasks are organized by priority and dependencies.

## Phase 1: Foundation Setup

### 1.1 Project scaffolding
- [ ] Initialize git repository
- [ ] Create Go module structure with proper package layout
- [ ] Set up `go.mod` with dependencies:
  - `charmbracelet/bubbletea`
  - `charmbracelet/lipgloss`
  - `charmbracelet/bubbles`
- [ ] Create directory structure:
  - `cmd/vai/` - Application entry point
  - `internal/ui/` - UI components
  - `internal/vim/` - Mode system
  - `internal/session/` - Session management
  - `internal/chat/` - Chat buffer
  - `internal/config/` - Configuration
  - `pkg/clipboard/` - Clipboard abstraction

**Validation:** `go build ./cmd/vai` produces no errors

### 1.2 Configuration system
- [ ] Implement config loading from `~/.config/vai/config.yaml`
- [ ] Define default configuration structure
- [ ] Implement XDG directory resolution for data storage
- [ ] Create `~/.local/share/vai/sessions/` directory on first run

**Validation:** Running vai creates config and data directories if missing

---

## Phase 2: TUI Framework Core

### 2.1 Basic Bubble Tea application
- [ ] Create `tea.Program` entry point in `cmd/vai/main.go`
- [ ] Implement minimal Model struct with Mode, Focus, and sub-models
- [ ] Implement basic Init(), Update(), and View() functions
- [ ] Create simple "Hello World" View to verify Bubble Tea works
- [ ] Implement graceful shutdown on Ctrl+C

**Validation:** `vai` command starts and exits cleanly

### 2.2 Three-pane layout
- [ ] Define Layout struct with viewport dimensions
- [ ] Implement pane dimension calculation based on terminal size
- [ ] Create basic View function rendering three panes:
  - Status bar (top)
  - Session list (left)
  - Chat buffer (right)
  - Input area (bottom)
- [ ] Implement terminal resize handling (tea.WindowSizeMsg)

**Validation:** All panes render correctly at 80x24 and resize proportionally

### 2.3 Focus management
- [ ] Define Focus enum (History, Buffer, Input)
- [ ] Implement focus state in top-level Model
- [ ] Create Update() routing based on current focus
- [ ] Implement `Ctrl+w h/j/k/l` for pane switching
- [ ] Add visual indication of current focus (border highlight)

**Validation:** Focus switches between panes with `Ctrl+w` + hjkl, visual indicator updates

### 2.4 Status bar
- [ ] Create status bar component with:
  - Current mode display (NORMAL/INSERT/VISUAL)
  - Current focus area
  - Optional: session name
- [ ] Implement mode-specific colors (lipgloss styling)
- [ ] Integrate status bar into top-level View

**Validation:** Status bar shows current mode and focus, colors change with mode

---

## Phase 3: Vim Mode System

### 3.1 Mode types and state
- [ ] Define Mode enum (NORMAL, INSERT, VISUAL)
- [ ] Add Mode field to top-level Model
- [ ] Implement mode transition logic
- [ ] Update status bar to display current mode

**Validation:** Mode state persists and updates status bar

### 3.2 Mode-aware key routing
- [ ] Implement key routing logic in Update()
- [ ] Create mode-focus compatibility matrix
- [ ] Route keys based on (Mode, Focus) combination
- [ ] Handle invalid mode-focus transitions

**Validation:** Keys route correctly based on mode and focus

### 3.3 NORMAL mode basics
- [ ] Implement `i`/`a` to enter INSERT mode
- [ ] Implement `v` to enter VISUAL mode (only in chat buffer)
- [ ] Implement `Esc`/`Ctrl+c` to return to NORMAL
- [ ] Add movement keys in chat buffer: `j`, `k`, `Ctrl+f`, `Ctrl+b`, `G`, `gg`

**Validation:** Can navigate chat buffer in NORMAL mode, mode transitions work

### 3.4 INSERT mode
- [ ] Implement bubbles/textarea for input area
- [ ] Route character input to textarea in INSERT mode
- [ ] Handle special keys: Enter, Backspace, Delete, Arrows, Ctrl+w, Ctrl+u, Ctrl+a, Ctrl+e
- [ ] Implement Enter to "send" (echo to buffer for now)

**Validation:** Can type and edit text in input area, Enter adds to buffer

### 3.5 VISUAL mode (basic)
- [ ] Implement selection state in chat buffer
- [ ] Handle `v` to start selection at cursor
- [ ] Expand selection with movement keys
- [ ] Implement `y` to copy selection (placeholder for now)
- [ ] Implement `Esc` to cancel selection

**Validation:** Can select text in chat buffer, selection highlights

---

## Phase 4: Session Manager

### 4.1 Session data model
- [ ] Define Session and Message structs
- [ ] Define Block interface and implementations (TextBlock, CodeBlock)
- [ ] Implement JSON serialization/deserialization
- [ ] Add UUID generation for session and message IDs

**Validation:** Can serialize and deserialize sessions to/from JSON

### 4.2 Session storage
- [ ] Implement session file I/O with atomic writes (.tmp + rename)
- [ ] Implement session directory scanning
- [ ] Load sessions on startup
- [ ] Save sessions on every message

**Validation:** Sessions persist across restarts, atomic writes prevent corruption

### 4.3 Session list display
- [ ] Render session list in left pane
- [ ] Display: title, time ago, message count, active indicator
- [ ] Implement relative time formatting (< 1m, 2h, 3d, Jan 15)
- [ ] Sort sessions by UpdatedAt (newest first)

**Validation:** Session list shows all sessions with correct info and sorting

### 4.4 Session switching
- [ ] Implement `Enter` to select session from list
- [ ] Load selected session messages into chat buffer
- [ ] Update visual indicator for active session
- [ ] Implement `Ctrl+t` or `:new` to create new session

**Validation:** Can switch between sessions, chat buffer updates

### 4.5 Session operations
- [ ] Implement `r` to rename session (with inline prompt)
- [ ] Implement `dd` to delete session with confirmation
- [ ] Implement `/` to search sessions
- [ ] Implement `n`/`N` for next/prev search result

**Validation:** Can rename, delete, and search sessions

---

## Phase 5: Chat Buffer

### 5.1 Message rendering
- [ ] Render messages in chat buffer viewport
- [ ] Implement role-based styling (user vs assistant)
- [ ] Add message separators
- [ ] Display timestamps

**Validation:** Messages render with correct styling and separators

### 5.2 Text block rendering
- [ ] Render TextBlock content with word wrap
- [ ] Implement paragraph preservation
- [ ] Apply subtle styling for user messages
- [ ] Apply bright styling for assistant messages

**Validation:** Text blocks wrap correctly, paragraphs preserved

### 5.3 Code block rendering
- [ ] Render CodeBlock with distinct background
- [ ] Display language identifier
- [ ] Display code block number: [1], [2], etc.
- [ ] Use monospace font for code
- [ ] Preserve indentation

**Validation:** Code blocks render with proper styling and numbering

### 5.4 Viewport and scrolling
- [ ] Implement viewport-based rendering (only visible lines)
- [ ] Implement scrolling: `j`, `k`, `Ctrl+f`, `Ctrl+b`, `G`, `gg`
- [ ] Handle scroll bounds (no overflow)
- [ ] Auto-scroll to new messages if at bottom

**Validation:** Scrolling is smooth, handles 1000+ messages without lag

### 5.5 Code block navigation
- [ ] Implement `]c` to jump to next code block
- [ ] Implement `[c` to jump to previous code block
- [ ] Support counted jumps: `N]c`
- [ ] Highlight target code block briefly
- [ ] Show "No code blocks" message if none exist

**Validation:** Can jump between code blocks, visual feedback works

---

## Phase 6: Clipboard Integration

### 6.1 Clipboard abstraction
- [ ] Create clipboard package with platform detection
- [ ] Implement `pbcopy` for macOS
- [ ] Implement `xclip -selection clipboard` for Linux X11
- [ ] Implement `wl-copy` for Linux Wayland
- [ ] Return error if no clipboard command available

**Validation:** Clipboard abstraction works on macOS and Linux

### 6.2 Copy operations
- [ ] Implement `yc` to copy current code block
- [ ] Implement `yNc` to copy Nth code block
- [ ] Implement `ym` to copy entire message
- [ ] Implement VISUAL mode `y` to copy selection
- [ ] Show confirmation messages

**Validation:** All copy operations work, confirmations display

---

## Phase 7: Command Mode

### 7.1 Command prompt
- [ ] Implement `:` to enter command mode
- [ ] Create command input overlay
- [ ] Route keystrokes to command prompt
- [ ] Handle `Esc` to cancel

**Validation:** Can enter and exit command mode

### 7.2 Basic commands
- [ ] Implement `:q`/`:quit` to quit
- [ ] Implement `:q!` to quit without confirmation
- [ ] Implement `:new` to create new session
- [ ] Implement `:help`/`?` to show help overlay

**Validation:** Basic commands work as expected

### 7.3 Advanced commands (optional)
- [ ] Implement `:e {name}` to export session
- [ ] Implement `:import {file}` to import session
- [ ] Implement `:rename {name}` to rename session
- [ ] Implement `:clear` to clear current chat

**Validation:** Advanced commands work

---

## Phase 8: Markdown Parsing

### 8.1 Markdown to blocks parser
- [ ] Implement markdown parser (use existing library or simple regex)
- [ ] Extract fenced code blocks (```language ... ```) into CodeBlocks
- [ ] Assign sequential numbers to CodeBlocks
- [ ] Preserve remaining content as TextBlocks
- [ ] Handle malformed markdown gracefully

**Validation:** AI responses parse into correct block structure

### 8.2 Simple text formatting (optional)
- [ ] Convert `**bold**` to bright text
- [ ] Convert `` `code` `` to different color
- [ ] Keep other markdown as plain text

**Validation:** Basic markdown formatting renders

---

## Phase 9: Polish and Testing

### 9.1 Error handling
- [ ] Handle clipboard errors gracefully
- [ ] Handle file I/O errors without crashing
- [ ] Show user-friendly error messages
- [ ] Log technical errors to debug.log

**Validation:** Errors don't crash the app, are user-friendly

### 9.2 Performance optimization
- [ ] Implement lazy rendering for long conversations
- [ ] Add message pagination (load 100 at a time)
- [ ] Test with 1000+ messages
- [ ] Profile and optimize hot paths

**Validation:** Smooth performance with 1000+ messages

### 9.3 Configuration customization
- [ ] Support custom keybindings in config
- [ ] Support theme customization (colors)
- [ ] Implement config validation on startup

**Validation:** Custom keybindings and themes work

### 9.4 Documentation
- [ ] Create README with installation and usage
- [ ] Document all keybindings
- [ ] Add `?` help overlay with all bindings
- [ ] Create man page or comprehensive help

**Validation:** User can discover and learn all features

---

## Dependencies

Tasks must be completed in order, respecting these dependencies:

- **Phase 2** depends on Phase 1 (needs project setup)
- **Phase 3** depends on Phase 2 (needs TUI framework)
- **Phase 4** can proceed in parallel with Phase 5 (independent features)
- **Phase 5** depends on Phase 2 (needs TUI framework)
- **Phase 6** depends on Phase 5 (needs code blocks to copy)
- **Phase 7** depends on Phase 3 and 4 (needs modes and sessions)
- **Phase 8** depends on Phase 5 (needs chat buffer)
- **Phase 9** depends on all previous phases

## Parallelizable Work

These tasks can be worked on in parallel by multiple contributors:

- **Phase 4 (Session Manager)** and **Phase 5 (Chat Buffer)** are independent
- **Phase 6 (Clipboard)** can be developed separately once Phase 5 has basic structure
- **Phase 7 (Command Mode)** can be developed in parallel with Phase 8
- **Documentation (9.4)** can be written at any time

## Definition of Done

Each task is complete when:
1. Code is written and follows Go best practices
2. Feature works as specified in the requirements
3. Edge cases are handled (bounds checking, errors, etc.)
4. Visual feedback is clear (mode changes, focus indicators, confirmations)
5. Manual testing confirms the feature works end-to-end
