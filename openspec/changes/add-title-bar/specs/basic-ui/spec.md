# Spec: Basic UI

## ADDED Requirements

### Requirement: Title bar display

The UI MUST display a title bar at the top of the screen showing session information.

**Rationale:** Provides clear visual indication of the current session and improves UI hierarchy.

#### Scenario: Title bar displays session name

**Given** the vai application is running
**When** the UI renders
**Then** a title bar should be displayed at the top of the screen
**And** the title bar should show "Sessions - [current session title]"
**And** the text should be centered horizontally

#### Scenario: Title bar with default session name

**Given** the vai application is started
**When** no session title has been set
**Then** the title bar should display "Sessions - New Chat"

#### Scenario: Title bar with custom session name

**Given** the vai application is running
**When** the current session title is "My AI Conversation"
**Then** the title bar should display "Sessions - My AI Conversation"

#### Scenario: Title bar width adapts to terminal

**Given** the vai application is running
**When** the terminal is resized
**Then** the title bar width should match the terminal width
**And** the text should remain centered

### Requirement: Long session name display

The title bar MUST display long session names without truncation.

#### Scenario: Long session name displays fully

**Given** the current session has a very long title
**When** the title bar renders
**Then** the full session title should be visible
**And** no text should be truncated or abbreviated

### Requirement: Title bar styling

The title bar MUST use consistent styling to distinguish it from other UI elements.

#### Scenario: Title bar visual appearance

**Given** the title bar is displayed
**When** viewing the UI
**Then** the title bar should have a dark background (color 235)
**And** the text should be white (color 252)
**And** the text should be bold
**And** the title bar should be 1 line high

## MODIFIED Requirements

### Requirement: Layout calculation with title bar

The UI MUST calculate pane dimensions reserving space for the title bar.

#### Scenario: Content area below title bar

**Given** a terminal with height H
**When** calculating layout
**Then** the title bar should occupy Y=0 with height 1
**And** the content area should start at Y=1
**And** the content height should be H - titleBarHeight - inputHeight

#### Scenario: Pane positioning with title bar

**Given** the layout is calculated
**When** positioning panes
**Then** the session list and chat buffer should start at Y=1 (below title bar)
**And** the input area should start at Y=1 + contentHeight

### Requirement: Three-pane layout

The UI MUST render three distinct panes below the title bar: session list, chat buffer, and input area.

#### Scenario: Layout with title bar

**Given** the vai application is started
**When** the initial UI renders
**Then** a title bar should be visible at the top
**And** three panes should be visible below the title bar: session list (left), chat buffer (right), input area (bottom)
**And** each pane should have a border
