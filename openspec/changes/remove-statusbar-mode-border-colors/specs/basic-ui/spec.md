# Spec: Basic UI

## REMOVED Requirements

### Requirement: Status bar display

The UI MUST NOT display a status bar.

**Rationale:** Status bar consumes valuable screen space without providing significant value. Mode indication is moved to border colors.

#### Scenario: Application starts without status bar

**Given** the vai application is started
**When** the initial UI renders
**Then** no status bar should be displayed
**And** the content area should extend to the top of the screen

## ADDED Requirements

### Requirement: Mode indication via border colors

The UI MUST indicate the current Vim mode (NORMAL, INSERT, VISUAL) through border colors of non-focused panes.

#### Scenario: NORMAL mode border color

**Given** the application is in NORMAL mode
**When** viewing any non-focused pane
**Then** the pane border should be gray (color 240)

#### Scenario: INSERT mode border color

**Given** the application is in INSERT mode
**When** viewing any non-focused pane
**Then** the pane border should be green (color 142)

#### Scenario: VISUAL mode border color

**Given** the application is in VISUAL mode
**When** viewing any non-focused pane
**Then** the pane border should be blue (color 33)

### Requirement: Focused pane indication

The UI MUST indicate the currently focused pane with a distinct border style.

#### Scenario: Focused pane has thick cyan border

**Given** any pane has focus
**When** viewing the focused pane
**Then** the border should be thick (ThickBorder)
**And** the border color should be cyan (color 151)
**And** this should be true regardless of current mode

#### Scenario: Focus switches between panes

**Given** the chat buffer has focus
**When** user presses Ctrl+w to switch focus
**Then** the previous pane should lose the thick cyan border
**And** the newly focused pane should gain the thick cyan border

## MODIFIED Requirements

### Requirement: Layout calculation without status bar

The UI MUST calculate pane dimensions without reserving space for a status bar.

#### Scenario: Full height content area

**Given** a terminal with height H
**When** calculating layout
**Then** the content area height should be H - inputHeight
**And** no space should be reserved for status bar

#### Scenario: Pane positioning from top

**Given** the layout is calculated
**When** positioning panes
**Then** the session list and chat buffer should start at Y=0 (top of screen)
