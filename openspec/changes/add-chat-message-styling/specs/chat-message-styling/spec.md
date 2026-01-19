# Spec: Chat Message Styling

Visual distinction between user and AI messages in the chat panel through border colors, alignment, and labels.

## ADDED Requirements

### Requirement: User message styling

The chat panel MUST render user messages with a green border, right alignment, and a "You:" label.

#### Scenario: User message display

**Given** a message with RoleUser exists in the chat buffer
**When** the chat panel renders
**Then** the message MUST be displayed with a green border (color 142)
**And** the message MUST be right-aligned within the chat buffer
**And** a "You:" label MUST be displayed at the top of the message in green (color 142)
**And** the message content MUST be displayed in gray (color 244)

### Requirement: AI message styling

The chat panel MUST render AI messages with a blue border, left alignment, and an "AI:" label.

#### Scenario: AI message display

**Given** a message with RoleAssistant exists in the chat buffer
**When** the chat panel renders
**Then** the message MUST be displayed with a blue border (color 33)
**And** the message MUST be left-aligned within the chat buffer
**And** an "AI:" label MUST be displayed at the top of the message in blue (color 33)
**And** the message content MUST be displayed in white (color 252)

### Requirement: Message border style

Both user and AI messages MUST use the same border style (normal border) with only the color differing by role.

#### Scenario: Border style consistency

**Given** both user and AI messages exist in the chat buffer
**When** viewing the messages
**Then** both messages MUST use lipgloss.NormalBorder()
**And** only the border color MUST differ between user (green) and AI (blue) messages

### Requirement: Message content rendering

Messages MUST render their content blocks within the styled border.

#### Scenario: Multi-block message

**Given** a message contains multiple content blocks
**When** the message is rendered
**Then** all blocks MUST be displayed within the message border
**And** blocks MUST be stacked vertically
**And** each block's text MUST be wrapped to fit the available width

### Requirement: Message spacing

Multiple messages MUST be stacked vertically with consistent spacing between them.

#### Scenario: Multiple messages display

**Given** multiple messages exist in the chat buffer
**When** the chat panel renders
**Then** messages MUST be stacked vertically
**And** there MUST be 2 empty lines between consecutive messages
**And** the order MUST match the message creation time

### Requirement: Empty chat buffer

The chat panel MUST display a helpful placeholder when no messages exist.

#### Scenario: No messages

**Given** the chat buffer has no messages
**When** the chat panel renders
**Then** a placeholder text MUST be displayed
**And** the placeholder MUST say "No messages yet."
**And** the placeholder MUST suggest "Type a message to start chatting..."

### Requirement: Responsive message width

Message rendering MUST adapt to the chat buffer width.

#### Scenario: Terminal resize

**Given** messages are displayed in the chat buffer
**When** the terminal is resized
**Then** message width MUST recalculate based on new chat buffer width
**And** text MUST re-wrap to fit the new width
**And** borders MUST remain intact

### Requirement: Message renderer component

A ChatMessage component MUST encapsulate the rendering logic for individual messages.

#### Scenario: Message renderer usage

**Given** the chat buffer needs to render a message
**When** the ChatMessage.Render() method is called with a Message and maxWidth
**Then** the method MUST return a formatted string with appropriate styling
**And** the styling MUST be determined by the message's Role field
