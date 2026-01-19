# Spec: Vim Navigation (Refined)

**Status:** Active
**Version:** 1.1.0
**Owners:** vai project

## Overview

This spec defines the refined Vim-style navigation system. Key changes from v1.0:
- Removed command mode (`:`) for simpler initial implementation
- Added Vim-style movement keys in INSERT mode for the input area
- Clarified that code block operations only work in chat buffer focus

## MODIFIED Requirements

### Requirement: Mode System

The application MUST implement a Vim-inspired mode system for navigation and input.

#### Scenario: Mode types

**Given** the application is running
**When** modes are defined
**Then** the following modes MUST exist:

| Mode | Purpose | Icon/Label |
|------|---------|------------|
| NORMAL | Navigation, viewing, operations | NORMAL |
| INSERT | Text input with Vim-style editing | INSERT |
| VISUAL | Text selection in chat buffer only | VISUAL |

**And** the mode SHOULD be displayed in the status bar at all times

#### Scenario: Mode transitions

**Given** the user is in NORMAL mode
**When** the user performs specific actions:
- Press `i` or `a` → Enter INSERT mode (focus moves to input area)
- Press `v` in chat buffer → Enter VISUAL mode (only valid when focus is on chat buffer)
- Press `Esc` or `Ctrl+c` → Return to NORMAL mode (from any mode)

**Then** the mode SHOULD transition immediately
**And** the status bar SHOULD update to reflect the new mode

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
| `Ctrl+f` | Scroll down one page |
| `Ctrl+b` | Scroll up one page |
| `Ctrl+d` | Scroll down half a screen |
| `Ctrl+u` | Scroll up half a screen |
| `G` | Go to end of conversation |
| `gg` | Go to start of conversation |
| `w` | Move to start of next word |
| `b` | Move to start of previous word |
| `e` | Move to end of current word |
| `0` | Move to start of current line |
| `$` | Move to end of current line |

#### Scenario: Code block navigation (chat buffer only)

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses:
- `]c` - Jump to next code block
- `[c` - Jump to previous code block
- `N]c` - Jump N code blocks forward
**Then** the viewport SHOULD scroll to show the code block
**And** the code block SHOULD be highlighted briefly

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer (e.g., on session list)
**When** the user presses `]c` or `[c`
**Then** these keys SHOULD be ignored
**And** NO action SHOULD occur

#### Scenario: Code block copying (chat buffer only)

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses:
- `yc` - Copy the current/nearest code block
- `yNc` - Copy the Nth code block (1-indexed)
- `ym` - Copy the entire current message
**Then** the content SHOULD be copied to the system clipboard
**And** a confirmation SHOULD be displayed

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer
**When** the user presses `yc`, `yNc`, or `ym`
**Then** these keys SHOULD be ignored
**And** NO action SHOULD occur

#### Scenario: Pane switching

**Given** the user is in NORMAL mode
**When** the user presses `Ctrl+w` followed by:
- `h` - Move focus to session history pane
- `l` - Move focus to chat buffer pane
- `j` - Move focus to input area
- `k` - Move focus upward in pane order
**Then** focus SHOULD move to the specified pane
**And** the status bar SHOULD update to show the new focus

#### Scenario: Global shortcuts

**Given** the user is in NORMAL mode (any focus)
**When** the user presses:
- `Ctrl+t` - Create new session
- `Ctrl+q` - Quit application with confirmation
- `Ctrl+q` again (within 2s) - Quit without confirmation
- `?` - Show help overlay
**Then** the corresponding action SHOULD be executed

---

### Requirement: INSERT Mode Keybindings

The application MUST support text input with Vim-style movement when in INSERT mode.

**Note:** INSERT mode is only valid when focus is on the input area.

#### Scenario: Enter INSERT mode

**Given** the user is in NORMAL mode
**When** the user presses:
- `i` - Enter INSERT mode (focus moves to input area)
- `a` - Enter INSERT mode (focus moves to input area)
**Then** the mode SHOULD change to INSERT
**And** focus SHOULD move to the input area
**And** the cursor SHOULD be positioned in the input area

#### Scenario: Character input

**Given** the user is in INSERT mode
**When** the user types printable characters
**Then** the characters SHOULD be inserted into the input area
**And** the input area SHOULD scroll if content exceeds visible area

#### Scenario: Vim-style cursor movement in INSERT mode

**Given** the user is in INSERT mode
**When** the user presses Vim movement keys:
**Then** the cursor SHOULD move accordingly:

| Key | Action |
|-----|--------|
| `Esc` / `Ctrl+[` | Exit to NORMAL mode |
| `h` / `Ctrl+b` | Move cursor left |
| `l` / `Ctrl+f` | Move cursor right |
| `w` | Move to start of next word |
| `b` | Move to start of previous word |
| `e` | Move to end of current word |
| `0` / `Ctrl+a` | Move to start of line |
| `$` / `Ctrl+e` | Move to end of line |

#### Scenario: Editing keys in INSERT mode

**Given** the user is in INSERT mode
**When** the user presses editing keys:
**Then** the corresponding action SHOULD occur:

| Key | Action |
|-----|--------|
| `Enter` | Send the message |
| `Backspace` | Delete character before cursor |
| `Delete` / `Ctrl+d` | Delete character at cursor |
| `Ctrl+w` | Delete word before cursor |
| `Ctrl+u` | Delete to start of line |
| `Ctrl+k` | Delete to end of line |
| `Ctrl+h` | Delete character before cursor (same as Backspace) |
| `Arrow keys` | Move cursor (traditional fallback) |

#### Scenario: Send message

**Given** the user is in INSERT mode
**When** the user presses `Enter`
**Then** the message SHOULD be sent
**And** the mode SHOULD remain in INSERT (for continuous input)
**And** the input area SHOULD be cleared

#### Scenario: Multi-line input (optional for future)

**Given** the user is in INSERT mode
**When** the user presses a keybinding for multi-line (e.g., `Alt+Enter` or `Ctrl+j`)
**Then** a newline SHOULD be inserted instead of sending
**And** the message SHOULD NOT be sent

---

### Requirement: VISUAL Mode Keybindings

The application MUST support text selection in VISUAL mode (chat buffer only).

#### Scenario: VISUAL mode scope restriction

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer (e.g., on session list or input area)
**When** the user presses `v`
**Then** the key SHOULD be ignored
**And** the mode SHOULD remain NORMAL
**And** NO error SHOULD be displayed (silent ignore)

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses `v`
**Then** the mode SHOULD change to VISUAL
**And** selection SHOULD start at the current cursor position

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

---

### Requirement: Mode-Focus Compatibility Matrix

The application MUST enforce strict mode-focus compatibility.

#### Scenario: Valid mode-focus combinations

**Given** a key press or mode transition occurs
**When** the current mode and focus are checked
**Then** these combinations MUST be valid:

| Focus Area | NORMAL | INSERT | VISUAL |
|------------|--------|--------|--------|
| Session List | ✅ Valid | ❌ Invalid | ❌ Invalid |
| Chat Buffer | ✅ Valid | ❌ Invalid | ✅ Valid |
| Input Area | ✅ Valid | ✅ Valid | ❌ Invalid |

#### Scenario: Invalid mode-focus handling

**Given** the user attempts an invalid mode-focus transition
**When** the transition is detected
**Then** the application SHOULD:
- Silently ignore the invalid keypress
- OR automatically switch to a valid mode
- NOT display errors for common cases

---

### Requirement: Help System

The application MUST provide in-app help for keybindings.

#### Scenario: Help overlay

**Given** the user is in NORMAL mode
**When** the user presses `?`
**Then** a help overlay SHOULD appear showing:
- All available keybindings for the current mode
- Brief descriptions of each binding
- Mode-specific bindings organized by mode
- Focus-specific bindings (context-sensitive)

#### Scenario: Context-sensitive help

**Given** the user presses `?`
**When** focus is on a specific pane
**Then** the help SHOULD prioritize:
- Keybindings relevant to the current focus
- Cross-references to other modes
- "All bindings" option for complete reference

---

## REMOVED Requirements

### Requirement: Command Mode (REMOVED)

The Ex-style command mode (`:`) has been removed from the initial implementation to simplify the scope.

**Rationale:**
- Command mode adds significant complexity (parser, execution engine)
- All operations can be accomplished with direct keybindings
- Simpler initial implementation aligns with "minimal first" principle
- Can be added in future if needed

**Alternative keybindings for common operations:**

| Former Command | Replacement |
|----------------|-------------|
| `:q` / `:quit` | `Ctrl+q` (quit) |
| `:new` | `Ctrl+t` (new session) |
| `:delete` | `dd` in session list |
| `:rename` | `r` in session list |
| `:help` | `?` key |
| `:clear` | Can be added later if needed |
| `:e` / `:export` | Can be added later if needed |
| `:import` | Can be added later if needed |

---

## Cross-References

- **tui-framework**: Defines the top-level Model that manages mode state and key routing
- **chat-buffer**: Defines the code block navigation that is now explicitly scoped to chat buffer focus only
- **session-manager**: Defines the session list navigation operations
