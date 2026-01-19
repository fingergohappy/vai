## MODIFIED Requirements

### Requirement: User message styling

The chat panel MUST render user messages with a green border, right alignment, and a "You" label.

#### Scenario: User message display

**Given** a message with RoleUser exists in the chat buffer
**When** the chat panel renders
**Then** the message MUST be displayed with a green boxed border (color 142)
**And** the message MUST be right-aligned within the chat buffer
**And** the top border MUST include the label "You" centered within the border, with at least one space between the label and the border line on each side
**And** the message MUST include left/right padding between the border and the text
**And** the bubble width MUST shrink to fit the text when the text is narrower than 2/3 of the chat pane
**And** the bubble width MUST NOT exceed 2/3 of the chat pane width; when content exceeds the limit it MUST wrap until each rendered line fits
**And** wrapped content MUST be fully rendered (no truncation)

### Requirement: AI message styling

The chat panel MUST render AI messages with a blue border, left alignment, and an "AI" label.

#### Scenario: AI message display

**Given** a message with RoleAssistant exists in the chat buffer
**When** the chat panel renders
**Then** the message MUST be displayed with a blue boxed border (color 33)
**And** the message MUST be left-aligned within the chat buffer
**And** the top border MUST include the label "AI" centered within the border, with at least one space between the label and the border line on each side
**And** the message MUST include left/right padding between the border and the text
**And** the bubble width MUST shrink to fit the text when the text is narrower than 2/3 of the chat pane
**And** the bubble width MUST NOT exceed 2/3 of the chat pane width; when content exceeds the limit it MUST wrap until each rendered line fits
**And** wrapped content MUST be fully rendered (no truncation)

### Requirement: Message border style

Both user and AI messages MUST use the same boxed border style (lipgloss.NormalBorder) with only the color differing by role.

#### Scenario: Border style consistency

**Given** both user and AI messages exist in the chat buffer
**When** viewing the messages
**Then** both messages MUST use lipgloss.NormalBorder()
**And** only the border color MUST differ between user (green) and AI (blue) messages
