# Spec: Vim Navigation

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines the Vim-style navigation system including modes, keybindings, and the mode-aware key routing that enables keyboard-first interaction with the TUI application.

## ADDED Requirements

### Requirement: Mode System

The application MUST implement a Vim-inspired mode system for navigation and input.

#### Scenario: Mode types

**Given** the application is running
**When** modes are defined
**Then** the following modes MUST exist:

| Mode | Purpose | Icon/Label |
|------|---------|------------|
| NORMAL | Navigation, viewing, commands | NORMAL |
| INSERT | Text input in the input area | INSERT |
| VISUAL | Text selection in chat buffer | VISUAL |

**And** the mode SHOULD be displayed in the status bar at all times

#### Scenario: Mode transitions

**Given** the user is in NORMAL mode
**When** the user performs specific actions:
- Press `i` or `a` in chat buffer → Enter INSERT mode (focus moves to input)
- Press `v` in chat buffer → Enter VISUAL mode
- Press `Esc` or `Ctrl+c` → Return to NORMAL mode (from any mode)

**Then** the mode SHOULD transition immediately
**And** the status bar SHOULD update to reflect the new mode

#### Scenario: Mode-specific behaviors

**Given** the application is in a specific mode
**When** the user presses keys
**Then** the behavior SHOULD be:

| Mode | Key Input Behavior |
|------|-------------------|
| NORMAL | No text is entered; keys trigger commands/navigation |
| INSERT | Text is entered into the input area |
| VISUAL | Movement expands selection; no text entry |

---

### Requirement: NORMAL Mode Keybindings

The application MUST support Vim-style navigation keys when in NORMAL mode.

#### Scenario: Basic movement in chat buffer

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses movement keys:
**Then** the behavior SHOULD be:

| Key | Action |
|-----|--------|
| `j` or `Ctrl+e` | Scroll down one line |
| `k` or `Ctrl+y` | Scroll up one line |
| `Ctrl+f` | Scroll down one page (half screen) |
| `Ctrl+b` | Scroll up one page (half screen) |
| `Ctrl+d` | Scroll down half a screen |
| `Ctrl+u` | Scroll up half a screen |
| `G` | Go to end of conversation (latest message) |
| `gg` | Go to start of conversation (oldest message) |

#### Scenario: Word and line movement

**Given** the user is in NORMAL mode in the chat buffer
**When** the user presses:
- `w` - Move to start of next word
- `b` - Move to start of previous word
- `e` - Move to end of current word
- `0` - Move to start of current line
- `$` - Move to end of current line
**Then** the cursor/viewport SHOULD move accordingly

#### Scenario: Count prefixes

**Given** the user is in NORMAL mode
**When** the user presses a count before a command (e.g., `5j`, `3w`)
**Then** the command SHOULD execute the specified number of times
**And** if the count exceeds the buffer bounds, it SHOULD clamp to valid range

#### Scenario: Pane switching

**Given** the user is in NORMAL mode
**When** the user presses `Ctrl+w` followed by:
- `h` - Move focus to session history pane
- `l` - Move focus to chat buffer pane
- `j` - Move focus to input area
- `k` - Move focus upward in pane order
**Then** focus SHOULD move to the specified pane
**And** the status bar SHOULD update to show the new focus

---

### Requirement: INSERT Mode Keybindings

The application MUST support text input when in INSERT mode.

#### Scenario: Enter INSERT mode

**Given** the user is in NORMAL mode
**When** the user presses:
- `i` - Enter INSERT mode at current position (moves to input area)
- `a` - Enter INSERT mode after current position (moves to input area)
**Then** the mode SHOULD change to INSERT
**And** focus SHOULD move to the input area
**And** the cursor SHOULD be positioned in the input area

#### Scenario: Text input

**Given** the user is in INSERT mode
**When** the user types printable characters
**Then** the characters SHOULD be inserted into the input area
**And** the input area SHOULD scroll if content exceeds visible area

#### Scenario: Special keys in INSERT mode

**Given** the user is in INSERT mode
**When** the user presses:
- `Enter` - Send the message / add newline (if multi-line)
- `Backspace` - Delete character before cursor
- `Delete` - Delete character at cursor
- `Arrow keys` - Move cursor within input
- `Ctrl+w` - Delete word before cursor
- `Ctrl+u` - Delete to start of line
- `Ctrl+a` - Move to start of line
- `Ctrl+e` - Move to end of line
**Then** the corresponding action SHOULD occur in the input area

#### Scenario: Exit INSERT mode

**Given** the user is in INSERT mode
**When** the user presses `Esc` or `Ctrl+[`
**Then** the mode SHOULD change to NORMAL
**And** focus SHOULD return to the chat buffer
**And** the input content SHOULD be preserved (not cleared)

---

### Requirement: VISUAL Mode Keybindings

The application MUST support text selection in VISUAL mode.

#### Scenario: Enter VISUAL mode

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses `v`
**Then** the mode SHOULD change to VISUAL
**And** selection SHOULD start at the current cursor position
**And** the current character SHOULD be highlighted

#### Scenario: Expand selection

**Given** the user is in VISUAL mode
**When** the user presses movement keys (`h`/`j`/`k`/`l`/`w`/`b`/`0`/`$`/etc.)
**Then** the selection SHOULD expand to include the new position
**And** the selected text SHOULD be highlighted

#### Scenario: Copy selection

**Given** the user has a selection in VISUAL mode
**When** the user presses `y`
**Then** the selected text SHOULD be copied to the system clipboard
**And** the mode SHOULD return to NORMAL
**And** a confirmation SHOULD be displayed: "Copied selection to clipboard"

#### Scenario: Cancel selection

**Given** the user is in VISUAL mode
**When** the user presses `Esc` or `Ctrl+c`
**Then** the selection SHOULD be cleared
**And** the mode SHOULD return to NORMAL
**And** no confirmation SHOULD be displayed (silent cancel)

---

### Requirement: Code Block Navigation Keybindings

The application MUST support special keybindings for code block operations.

#### Scenario: Jump to code blocks

**Given** the user is in NORMAL mode in the chat buffer
**When** the user presses:
- `]c` - Jump to next code block in current message
- `[c` - Jump to previous code block in current message
**Then** the viewport SHOULD scroll to show the code block
**And** the code block SHOULD be highlighted briefly

#### Scenario: Counted code block jumps

**Given** the user is in NORMAL mode
**When** the user presses `N]c` (where N is a count)
**Then** the viewport SHOULD jump N code blocks forward
**And** if N exceeds the number of code blocks, it SHOULD wrap or clamp

#### Scenario: Copy code blocks

**Given** the user is in NORMAL mode in the chat buffer
**When** the user presses:
- `yc` - Copy the current/nearest code block
- `yNc` - Copy the Nth code block (1-indexed)
- `ym` - Copy the entire current message
**Then** the content SHOULD be copied to the system clipboard
**And** a confirmation SHOULD be displayed

#### Scenario: No code block handling

**Given** the current message has no code blocks
**When** the user presses `]c`, `[c`, or `yc`
**Then** a message SHOULD be displayed: "No code blocks in this message"
**And** no other action SHOULD occur

---

### Requirement: Session List Navigation

The application MUST support Vim-style navigation in the session list.

#### Scenario: Session list movement

**Given** the user is in NORMAL mode
**And** focus is on the session list
**When** the user presses:
- `j` or `Ctrl+e` - Move to next session
- `k` or `Ctrl+y` - Move to previous session
- `G` - Go to last session (most recent)
- `gg` - Go to first session (oldest)
**Then** the selection SHOULD move accordingly
**And** the viewport SHOULD scroll if needed

#### Scenario: Session list actions

**Given** the user is in NORMAL mode in the session list
**When** the user presses:
- `Enter` - Open the selected session
- `r` - Rename the selected session
- `dd` - Delete the selected session
- `/` - Search sessions
- `n` - Next search result
- `N` - Previous search result
**Then** the corresponding action SHOULD be performed

#### Scenario: Create new session

**Given** the user is in NORMAL mode
**When** the user presses:
- `Ctrl+t` - Create new session
- Or types `:new` and presses Enter
**Then** a new empty session SHOULD be created
**And** it SHOULD become the active session

---

### Requirement: Command Mode

The application MUST support an Ex-style command mode for advanced operations.

#### Scenario: Enter command mode

**Given** the user is in NORMAL mode
**When** the user presses `:`
**Then** a command prompt SHOULD appear
**And** focus SHOULD move to the command prompt
**And** subsequent keystrokes SHOULD be captured as the command

#### Scenario: Execute command

**Given** the user is in command mode
**When** the user types a command and presses `Enter`
**Then** the command SHOULD be executed
**And** the prompt SHOULD close
**And** focus SHOULD return to the previous location

#### Scenario: Supported commands

**Given** the user is in command mode
**When** the user types:
**Then** these commands SHOULD be supported:

| Command | Action |
|---------|--------|
| `:q` or `:quit` | Quit the application |
| `:q!` | Quit without confirmation |
| `:w` | Save current session (manual save) |
| `:new` | Create new session |
| `:e {name}` or `:export {name}` | Export current session |
| `:import {file}` | Import session from file |
| `:delete` | Delete current session |
| `:rename {name}` | Rename current session |
| `:help` or `:?` | Show help overlay |
| `:clear` | Clear current chat (keep session) |

#### Scenario: Cancel command mode

**Given** the user is in command mode
**When** the user presses `Esc`
**Then** the command prompt SHOULD close
**And** focus SHOULD return to the previous location
**And** no command SHOULD be executed

---

### Requirement: Mode-Aware Key Routing

The application MUST route key events based on mode and focus.

#### Scenario: Key routing logic

**Given** a key press occurs
**When** the key event is processed
**Then** the routing SHOULD follow this logic:

1. Check current mode
2. Check current focus area
3. Check if the focus area supports the current mode
4. Route to appropriate handler based on mode + focus combination

#### Scenario: Mode-focus compatibility matrix

**Given** a key press occurs
**When** the current mode and focus are checked
**Then** key routing SHOULD respect this compatibility:

| Focus Area | NORMAL | INSERT | VISUAL |
|------------|--------|--------|--------|
| Session List | ✅ Enabled | ❌ Invalid | ❌ Invalid |
| Chat Buffer | ✅ Enabled | ❌ Invalid | ✅ Enabled |
| Input Area | ❌ Invalid | ✅ Enabled | ❌ Invalid |

**And** if a key press results in an invalid mode-focus combination:
- The mode SHOULD automatically switch to a valid mode
- Or the key SHOULD be ignored with a brief indicator

---

### Requirement: Keybinding Customization

The application MUST support user customization of keybindings.

#### Scenario: Keybinding configuration

**Given** the user wants to customize keybindings
**When** keybindings are defined in the config file
**Then** the application SHOULD:
- Load custom keybindings from config
- Fall back to defaults for undefined bindings
- Validate bindings on startup

#### Scenario: Config file format

**Given** the user edits the config file
**When** defining custom keybindings
**Then** the format SHOULD allow:

```yaml
keybindings:
  normal:
    next_line: "j"
    prev_line: "k"
    next_code_block: "]c"
    prev_code_block: "[c"
  visual:
    copy_selection: "y"
  command:
    quit: ":q"
```

---

### Requirement: Help and Documentation

The application MUST provide in-app help for keybindings.

#### Scenario: Help overlay

**Given** the user is in NORMAL mode
**When** the user presses `?` or types `:help`
**Then** a help overlay SHOULD appear showing:
- All available keybindings for the current mode
- Brief descriptions of each binding
- Mode-specific bindings organized by mode

#### Scenario: Context-sensitive help

**Given** the user presses `?`
**When** focus is on a specific pane
**Then** the help SHOULD prioritize:
- Keybindings relevant to the current focus
- Cross-references to other modes
- "All bindings" option for complete reference

---

## Cross-References

- **tui-framework**: Defines the top-level Model that manages mode state and key routing
- **chat-buffer**: Defines the code block navigation (']c', '[c', 'yc') and VISUAL mode operations
- **session-manager**: Defines the session list navigation operations
