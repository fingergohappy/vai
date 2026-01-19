# Spec: Basic UI Framework

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines the basic UI framework skeleton - the visual shell of the application with placeholder content. This is Phase 1 of the incremental implementation strategy.

## ADDED Requirements

### Requirement: Three-Pane Layout

The application MUST render a three-pane layout with placeholder content.

#### Scenario: Initial layout display

**Given** the application starts
**When** the UI is rendered
**Then** it SHOULD display:

```
┌─────────────────────────────────────────┐
│ NORMAL | Buffer                        │  ← Status bar
├──────────────┬──────────────────────────┤
│ [Sessions]   │ [Chat Buffer]            │  ← Session list
│              │                          │     and chat buffer
│ - Session 1  │ - Welcome to vai!        │
│ - Session 2  │ - This is a placeholder │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │  ← Input area
└─────────────────────────────────────────┘
```

#### Scenario: Layout proportions

**Given** the terminal is 80x24 or larger
**When** the layout is calculated
**Then** the dimensions SHOULD be:
- **Status bar:** Top 1 line
- **Session list:** Left 20-25% of width
- **Chat buffer:** Remaining width (75-80%)
- **Input area:** Bottom 3-5 lines
- **Content area:** Remaining height between panes

#### Scenario: Terminal resize

**Given** the application is running
**When** the terminal window is resized
**Then** the layout SHOULD:
- Recalculate pane proportions
- Redraw without distortion
- Maintain current focus

---

### Requirement: Status Bar

The application MUST display a status bar showing current mode and focus.

#### Scenario: Status bar content

**Given** the application is running
**When** the status bar is rendered
**Then** it MUST display:
- Current mode (NORMAL/INSERT)
- Current focus area (Session/Buffer/Input)
- Styled with distinct colors for different modes

#### Scenario: Mode colors

**Given** the status bar is displayed
**When** the mode changes
**Then** the mode text SHOULD use:
- White/neutral color for NORMAL mode
- Green color for INSERT mode
- Blue color for VISUAL mode (when implemented later)

#### Scenario: Status bar position

**Given** the layout is rendered
**When** measuring the status bar
**Then** it SHOULD:
- Be at the top of the screen (line 0)
- Span the full terminal width
- Have a height of 1 line

---

### Requirement: Focus Switching

The application MUST support switching focus between the three panes.

#### Scenario: Switch focus with Ctrl+w

**Given** the application is in NORMAL mode
**When** the user presses `Ctrl+w` followed by direction key:
- `Ctrl+w h` → Focus moves to session list
- `Ctrl+w j` → Focus moves to input area
- `Ctrl+w k` → Focus moves upward in pane order
- `Ctrl+w l` → Focus moves to chat buffer
**Then** the focused pane SHOULD be visually highlighted

#### Scenario: Visual focus indication

**Given** focus is on a specific pane
**When** that pane is rendered
**Then** it SHOULD display:
- A thicker or colored border
- A distinct background color
- Or other visual indicator

#### Scenario: Default focus

**Given** the application starts
**When** no user interaction has occurred
**Then** focus SHOULD default to the chat buffer pane

---

### Requirement: Mode Display

The application MUST maintain and display a Vim-style mode.

#### Scenario: Mode states

**Given** the application is running
**When** modes are implemented
**Then** the following modes MUST exist:
- **NORMAL** - Default mode for navigation
- **INSERT** - For typing in input area (placeholder in Phase 1)

#### Scenario: Mode transitions

**Given** the application is in NORMAL mode
**When** the user presses `i` or `a`
**Then** the mode SHOULD change to INSERT
**And** focus SHOULD move to the input area
**And** the status bar SHOULD update

**Given** the application is in INSERT mode
**When** the user presses `Esc`
**Then** the mode SHOULD return to NORMAL
**And** focus SHOULD return to the chat buffer
**And** the status bar SHOULD update

---

### Requirement: Placeholder Content

Each pane MUST display placeholder content indicating functionality to be implemented.

#### Scenario: Session list placeholder

**Given** the session list pane is rendered
**When** no real sessions exist yet
**Then** it SHOULD display:
- A title like "[Sessions]"
- 2-3 example session items
- A note like "(TODO: implement)"

#### Scenario: Chat buffer placeholder

**Given** the chat buffer pane is rendered
**When** no real messages exist yet
**Then** it SHOULD display:
- A title like "[Chat Buffer]"
- A welcome message
- A note like "This is a placeholder. Real implementation coming soon..."

#### Scenario: Input area placeholder

**Given** the input area is rendered
**When** the input component is initialized
**Then** it SHOULD display:
- A placeholder text like "[Type your message...]"
- The input should be a visual placeholder (not functional in Phase 1)

---

### Requirement: Quit Functionality

The application MUST support quitting with a keyboard shortcut.

#### Scenario: Quit with Ctrl+q

**Given** the application is running
**When** the user presses `Ctrl+q`
**Then** the application SHOULD:
- Display a confirmation message
- Quit if pressed again within 2 seconds
- OR quit immediately if a flag is set

#### Scenario: Quit with Ctrl+c

**Given** the application is running
**When** the user presses `Ctrl+c`
**Then** the application SHOULD quit immediately

---

### Requirement: Bubble Tea Integration

The application MUST use Bubble Tea framework for the TUI.

#### Scenario: Program initialization

**Given** the application starts
**When** the Bubble Tea program is created
**Then** it SHOULD:
- Create a Program with the top-level Model
- Use `WithAltScreen()` for full-screen mode
- Handle errors gracefully

#### Scenario: Model structure

**Given** the top-level Model is defined
**When** it is structured
**Then** it SHOULD contain:
```go
type Model struct {
    Mode   vim.Mode
    Focus  ui.Focus
    Quits  bool

    // Sub-models (all placeholders in Phase 1)
    Session session.Model
    Buffer  chat.Model
    Input   input.Model
}
```

---

### Requirement: Lipgloss Styling

The application MUST use Lipgloss for styling UI elements.

#### Scenario: Style definitions

**Given** styles are defined
**When** they are used
**Then** the application SHOULD define:
- Border styles for panes
- Color styles for modes
- Focused border styles
- Text styles

#### Scenario: Style application

**Given** a pane is rendered
**When** it is focused vs unfocused
**Then** the styling SHOULD differ:
- Focused: thicker/brighter border
- Unfocused: normal/dimmer border

---

## MODIFIED Requirements

### Requirement: Application Entry Point (Modified)

The application entry point MUST initialize the basic UI framework.

#### Scenario: Main function initialization

**Given** `main.go` is executed
**When** the application starts
**Then** it SHOULD:
- Load configuration (use defaults for now)
- Create the top-level Model with placeholder sub-models
- Start the Bubble Tea Program
- Run and exit cleanly

---

## Cross-References

- **project-structure**: Defines the directory structure
- **vim-navigation**: Will be implemented in Phase 3
- **chat-buffer**: Placeholder becomes real in Phase 2 and 5
- **session-manager**: Placeholder becomes real in Phase 4

---

## Implementation Notes

### Phase 1 Scope

This spec covers ONLY Phase 1 - the visual skeleton. Future phases will add real functionality.

### What Phase 1 DOES NOT include:

- ❌ Real text input (just visual placeholder)
- ❌ Real message storage
- ❌ Real session management
- ❌ Markdown rendering
- ❌ Code block operations
- ❌ File I/O

### What Phase 1 DOES include:

- ✅ Visual layout rendering
- ✅ Focus switching
- ✅ Mode display
- ✅ Quit functionality
- ✅ Basic Bubble Tea structure

### Success Criteria

Phase 1 is successful when:
1. Running `vai` shows the three-pane layout
2. User can switch focus with `Ctrl+w h/j/k/l`
3. Status bar shows mode and focus correctly
4. Pressing `i` switches to INSERT mode
5. Pressing `Esc` returns to NORMAL mode
6. Pressing `Ctrl+q` quits the application
