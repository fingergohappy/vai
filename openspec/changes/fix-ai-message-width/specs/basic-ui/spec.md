## MODIFIED Requirements

### Requirement: Three-pane layout

The UI MUST render three distinct panes: session list, chat buffer, and input area.

#### Scenario: Initial layout rendering

**Given** the vai application is started
**When** the initial UI renders
**Then** three panes should be visible: session list (left), chat buffer (right), input area (bottom)
**And** each pane should have a border
**And** chat messages within the chat buffer SHOULD avoid excessive horizontal whitespace for short content
**And** message bubbles MUST NOT exceed 2/3 of the chat buffer pane width; content beyond the limit MUST wrap

### Requirement: Focus indication

The UI MUST indicate which pane currently has focus.

#### Scenario: Focused pane has distinct border

**Given** any pane has focus
**When** viewing the UI
**Then** the focused pane should have a thick cyan border
**And** non-focused panes should have normal gray borders
