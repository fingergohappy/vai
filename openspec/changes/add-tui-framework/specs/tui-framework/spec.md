# Spec: TUI Framework

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines the core TUI framework architecture for the vai AI chat terminal application, built on Bubble Tea with a three-pane layout and mode-based navigation system.

## ADDED Requirements

### Requirement: Application Entry Point

The application MUST provide a CLI entry point that initializes the Bubble Tea program with proper configuration.

#### Scenario: Launch application

**Given** the user has installed vai
**When** the user runs `vai` command
**Then** the TUI application SHOULD start with:
- Full-screen terminal mode
- Default three-pane layout visible
- Status bar showing current mode (NORMAL) and focus area
- Left pane showing session history (empty or loaded from storage)
- Center pane showing current chat buffer (empty or loaded)
- Bottom pane showing input area ready for input

#### Scenario: Clean shutdown

**Given** the application is running
**When** the user presses `Ctrl+C` or `:q`
**Then** the application SHOULD:
- Exit gracefully without errors
- Preserve any unsaved state to temporary storage
- Restore terminal state

---

### Requirement: Three-Pane Layout

The application MUST implement a three-pane layout matching the design document specifications.

#### Scenario: Layout dimensions

**Given** the application has started
**When** the terminal is sized at 80x24 or larger
**Then** the layout SHOULD render:
- **Top**: Status bar spanning full width (1 line height)
- **Left**: Session history pane (20-25% width, remaining height minus input)
- **Right**: Chat buffer pane (remaining width, remaining height minus input)
- **Bottom**: Input area (full width, 3-5 lines height)

#### Scenario: Terminal resize handling

**Given** the application is running
**When** the terminal window is resized
**Then** the layout SHOULD:
- Recalculate pane dimensions proportionally
- Redraw all panes without distortion
- Preserve current scroll positions in each pane
- Maintain visible content within new bounds

#### Scenario: Minimum terminal size

**Given** the user tries to start vai
**When** the terminal is smaller than 80x24
**Then** the application SHOULD:
- Display an error message indicating minimum size requirement
- Exit gracefully without starting TUI

---

### Requirement: Focus Management

The application MUST implement focus management for routing keyboard events to the appropriate pane.

#### Scenario: Default focus

**Given** the application has just started
**When** no user interaction has occurred
**Then** focus SHOULD default to the **Chat Buffer** pane
**And** the status bar SHOULD indicate the current focus

#### Scenario: Pane switching

**Given** the application is running
**When** the user presses `Ctrl-w` followed by `h`/`j`/`k`/`l`
**Then** focus SHOULD move:
- `h`: Move focus to **Session History** pane (if available)
- `l`: Move focus to **Chat Buffer** pane
- `j`: Move focus to **Input Area** pane
- `k`: Move focus upwards in pane order (Buffer → History)

#### Scenario: Visual focus indication

**Given** focus is on a specific pane
**When** that pane is rendered
**Then** the pane SHOULD display a visible border or highlight
**And** other panes SHOULD render without focus indication

---

### Requirement: Mode System Integration

The application MUST integrate with the Vim-style mode system defined in `vim-navigation` spec.

#### Scenario: Mode display

**Given** the application is running
**When** the current mode changes (NORMAL/INSERT/VISUAL/COPY)
**Then** the status bar SHOULD display the current mode prominently
**And** the display SHOULD update immediately on mode change

#### Scenario: Mode-aware focus constraints

**Given** the application is running
**When** the current mode is **INSERT**
**And** focus is NOT on the Input Area
**Then** keyboard input SHOULD NOT enter text into any pane
**And** the user SHOULD be prompted to switch to Input Area or change mode

#### Scenario: Mode-dependent key routing

**Given** the application is running
**When** a key press occurs
**Then** the key event SHOULD be routed based on:
1. Current focus pane
2. Current mode
3. Pane-mode compatibility matrix (see `vim-navigation` spec)

---

### Requirement: Status Bar

The application MUST provide a status bar displaying application state information.

#### Scenario: Status bar content

**Given** the application is running
**When** the status bar is rendered
**Then** it SHOULD display:
- Current Vim mode (NORMAL/INSERT/VISUAL/COPY)
- Current focus area (History/Buffer/Input)
- Optional: session name or indicator

#### Scenario: Status bar styling

**Given** the status bar is being rendered
**When** different modes are active
**Then** the mode display SHOULD use distinct colors:
- NORMAL: neutral/white
- INSERT: green
- VISUAL: blue
- COPY: yellow

---

### Requirement: Bubble Tea Model Architecture

The application MUST implement a proper Bubble Tea Model hierarchy with autonomous sub-models.

#### Scenario: Top-level model structure

**Given** the application architecture
**When** the top-level Model is defined
**Then** it MUST contain:
```go
type Model struct {
    Mode   vim.Mode       // Current Vim mode
    Focus  ui.Focus       // Current focus area
    History history.Model // Session history sub-model
    Buffer  chatbuffer.Model // Chat buffer sub-model
    Input   input.Model   // Input area sub-model
}
```

#### Scenario: Message dispatching

**Given** the top-level Model receives a Bubble Tea Msg
**When** the Update function processes the message
**Then** it SHOULD:
- Determine which sub-model(s) should receive the message
- Dispatch to sub-models based on focus and mode
- Aggregate returned Cmds from sub-models
- Return updated Model and combined Cmd(s)

#### Scenario: Sub-model autonomy

**Given** a sub-model (e.g., Buffer)
**When** it processes a message and needs to notify other components
**Then** it SHOULD emit a Bubble Tea Msg
**And** the top-level Model SHOULD receive and handle that message

---

### Requirement: Dependency Integration

The application MUST properly integrate required Bubble Tea ecosystem dependencies.

#### Scenario: Bubble Tea framework

**Given** the application go.mod
**When** dependencies are resolved
**Then** `charmbracelet/bubbletea` MUST be included at latest stable version
**And** the application MUST use `tea.Program` for the main loop

#### Scenario: Lipgloss styling

**Given** the application needs styling
**When** UI elements are rendered
**Then** `charmbracelet/lipgloss` MUST be used for all styling
**And** styles SHOULD be defined in a centralized `ui/styles.go` package

#### Scenario: Bubbles components

**Given** the application needs standard components
**When** the Input Area is implemented
**Then** `charmbracelet/bubbles/textarea` MUST be used for multi-line input
**And** `charmbracelet/bubbles/viewport` MAY be used for scrollable panes

---

### Requirement: Configuration Management

The application MUST support configuration for user preferences.

#### Scenario: Configuration file location

**Given** the user wants to customize vai
**When** vai starts
**Then** it SHOULD look for configuration in:
1. `~/.config/vai/config.yaml` (or platform equivalent)
2. Current directory `.vai.yaml` (for project-specific config)
3. Default built-in configuration if none found

#### Scenario: Theme configuration

**Given** the configuration file exists
**When** theme settings are defined
**Then** the application SHOULD apply custom colors and styles
**And** fall back to default theme for undefined values

#### Scenario: Keybinding customization

**Given** the configuration file exists
**When** custom keybindings are defined
**Then** the application SHOULD use custom bindings
**And** fall back to default Vim-style bindings for undefined keys

---

### Requirement: Error Handling and Logging

The application MUST implement proper error handling and optional logging.

#### Scenario: Graceful error recovery

**Given** the application encounters a non-fatal error
**When** the error occurs during normal operation
**Then** the application SHOULD:
- Display a user-friendly error message in the status bar or overlay
- Log technical details to log file if logging is enabled
- Continue operation if possible

#### Scenario: Logging configuration

**Given** the user enables debug logging
**When** `--debug` flag is set or config enables logging
**Then** the application SHOULD:
- Write logs to `~/.config/vai/debug.log`
- Include timestamp, log level, and context
- Respect log level configuration (ERROR, WARN, INFO, DEBUG)

---

## Cross-References

- **vim-navigation**: Defines mode system and keybindings that this framework integrates
- **chat-buffer**: Defines the chat buffer sub-model this framework hosts
- **session-manager**: Defines the session history sub-model this framework hosts
