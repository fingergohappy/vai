# Spec: Basic UI

## Purpose

Basic UI rendering and layout for the vai application.

## Requirements

### Requirement: Three-pane layout

The UI MUST render three distinct panes: session list, chat buffer, and input area.

#### Scenario: Initial layout rendering

**Given** the vai application is started
**When** the initial UI renders
**Then** three panes should be visible: session list (left), chat buffer (right), input area (bottom)
**And** each pane should have a border

### Requirement: Focus indication

The UI MUST indicate which pane currently has focus.

#### Scenario: Focused pane has distinct border

**Given** any pane has focus
**When** viewing the UI
**Then** the focused pane should have a thick cyan border
**And** non-focused panes should have normal gray borders

### Requirement: Mode-based styling

The UI MUST indicate the current Vim mode through visual styling.

#### Scenario: Mode indication

**Given** the application is running
**When** the mode changes (NORMAL/INSERT/VISUAL)
**Then** the UI should visually indicate the current mode

### Requirement: Layout calculation

The UI MUST calculate pane dimensions based on terminal size.

#### Scenario: Terminal resize

**Given** the application is running
**When** the terminal is resized
**Then** pane dimensions should update to fit the new terminal size
**And** the layout should remain proportional

### Requirement: Focus switching

The UI MUST support switching focus between panes.

#### Scenario: Switch focus with Ctrl+w

**Given** the application is in NORMAL mode
**When** user presses Ctrl+w
**Then** focus should cycle to the next pane
**And** the border styling should update to reflect the new focus
