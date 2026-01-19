# Spec: Chat Buffer

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines the chat buffer component that displays and enables interaction with AI conversation history, featuring structured message rendering, code block operations, and efficient viewing for long conversations.

## ADDED Requirements

### Requirement: Structured Message Data Model

The chat buffer MUST use a structured data model for messages rather than raw text strings.

#### Scenario: Message structure

**Given** an AI conversation message
**When** the message is represented in memory
**Then** it MUST follow this structure:
```go
type Message struct {
    ID      string   // Unique message identifier
    Role    string   // "user" or "assistant" (or "ai")
    Blocks  []Block  // Ordered content blocks
    Created time.Time // Timestamp
}

type Block interface {
    Kind() BlockType
    Render(width int) string
}

type BlockType int
const (
    TextBlock BlockType = iota
    CodeBlock
)

type TextBlock struct {
    Text string // Plain text or markdown content
}

type CodeBlock struct {
    Lang    string   // Language identifier (go, python, bash, etc.)
    Lines   []string // Code content split by lines
    Number  int      // Sequential code block number in message
}
```

#### Scenario: Multiple blocks per message

**Given** an AI assistant response
**When** the response contains both explanatory text and multiple code examples
**Then** the Message SHOULD contain:
- One or more TextBlocks for explanations
- Zero or more CodeBlocks for code examples
- Blocks in the order they appeared in the original response

#### Scenario: Message with single code block

**Given** an AI response contains only a code block
**When** the message is rendered
**Then** the message SHOULD contain exactly one CodeBlock
**And** the CodeBlock's Number field SHOULD be set to 1

---

### Requirement: Message Rendering

The chat buffer MUST render messages in a readable format with appropriate styling for different block types.

#### Scenario: Text block rendering

**Given** a message containing a TextBlock
**When** the buffer is rendered to the viewport
**Then** the TextBlock SHOULD:
- Wrap text at the viewport width
- Preserve paragraph breaks (double newlines)
- Apply subtle styling (gray color for user, white for assistant)
- NOT render full markdown formatting (keep it simple)

#### Scenario: Code block rendering

**Given** a message containing a CodeBlock
**When** the buffer is rendered to the viewport
**Then** the CodeBlock SHOULD:
- Display with a distinct background color (darker than text)
- Show language identifier in top-right or top-left corner
- Show code block number in brackets: `[1]`, `[2]`, etc.
- Preserve original indentation and formatting
- Use a monospace font

#### Scenario: Message separator

**Given** two consecutive messages
**When** the buffer is rendered
**Then** messages SHOULD be separated by:
- A visible border or line
- Different background colors for alternating messages (optional)
- Role indicator (User/Assistant) at the start of each message

#### Scenario: Role-based styling

**Given** a message is being rendered
**When** the message Role is "user"
**Then** the message SHOULD use:
- Gray color for text
- Right-alignment or distinct marker
- NO code block highlighting (user messages rarely contain code)

**When** the message Role is "assistant" or "ai"
**Then** the message SHOULD use:
- White/bright color for text
- Full code block styling
- Code block numbers

---

### Requirement: Code Block Navigation

The chat buffer MUST provide Vim-style navigation for jumping between code blocks.

#### Scenario: Jump to next code block

**Given** the user is in NORMAL mode
**When** the user presses `]c`
**Then** the buffer SHOULD:
- Move cursor/viewport to the start of the next code block in the current message
- Scroll the viewport if necessary to make the code block visible
- Highlight the code block briefly or show visual indicator

#### Scenario: Jump to previous code block

**Given** the user is in NORMAL mode
**When** the user presses `[c`
**Then** the buffer SHOULD:
- Move cursor/viewport to the start of the previous code block in the current message
- Scroll the viewport if necessary
- Wrap to the last code block if at the first one

#### Scenario: Jump to specific code block

**Given** the user is in NORMAL mode
**When** the user presses `N]c` (where N is a count)
**Then** the buffer SHOULD:
- Jump forward count code blocks forward (e.g., `2]c` jumps 2 code blocks ahead)
- Clamp to the valid range of code blocks in the current message

#### Scenario: No code blocks in message

**Given** the current message has no code blocks
**When** the user presses `]c` or `[c`
**Then** the buffer SHOULD:
- Display a brief message: "No code blocks in this message"
- Do nothing else

---

### Requirement: Code Block Copying

The chat buffer MUST provide keyboard-driven code block copying operations.

#### Scenario: Copy current code block

**Given** the user has navigated to a code block
**When** the user presses `yc` in NORMAL mode
**Then** the buffer SHOULD:
- Copy the entire code block content to system clipboard
- Display confirmation: "Copied code block [N] to clipboard"
- Use the appropriate clipboard command:
  - macOS: `pbcopy`
  - Linux X11: `xclip -selection clipboard`
  - Linux Wayland: `wl-copy`

#### Scenario: Copy specific code block by number

**Given** the user is in NORMAL mode
**When** the user presses `yNc` (where N is a number)
**Then** the buffer SHOULD:
- Copy the Nth code block (1-indexed) from the current message
- Display confirmation with block number
- Show error if N exceeds the number of code blocks

#### Scenario: Copy entire message

**Given** the user is viewing a message
**When** the user presses `ym` in NORMAL mode
**Then** the buffer SHOULD:
- Copy the entire message content (all blocks) to clipboard
- Preserve code block formatting
- Display confirmation

#### Scenario: Clipboard command availability

**Given** the user tries to copy a code block
**When** no clipboard command is available on the system
**Then** the buffer SHOULD:
- Display error: "Clipboard not available. Please install pbcopy/xclip/wl-copy"
- NOT crash or throw an unhandled error

---

### Requirement: Viewport and Scrolling

The chat buffer MUST implement efficient viewport-based rendering for long conversations.

#### Scenario: Initial viewport position

**Given** a session is loaded
**When** the buffer is first displayed
**Then** the viewport SHOULD:
- Position at the end of the conversation (most recent messages)
- Show the last N messages that fit in the viewport
- Ensure the input area below remains visible

#### Scenario: Scrolling in NORMAL mode

**Given** the user is in NORMAL mode and focus is on the buffer
**When** the user presses Vim navigation keys:
- `j` / `Ctrl+e`: Scroll down one line
- `k` / `Ctrl+y`: Scroll up one line
- `Ctrl+f`: Scroll down one page
- `Ctrl+b`: Scroll up one page
- `G`: Go to end of conversation
- `gg`: Go to start of conversation
**Then** the viewport SHOULD scroll accordingly
**And** scrolling SHOULD be smooth (not jumpy)

#### Scenario: Scrolling beyond bounds

**Given** the user is at the top of the conversation
**When** the user presses `k` or `Ctrl+b`
**Then** the viewport SHOULD stay at the top
**And** no overflow error SHOULD occur

**Given** the user is at the bottom of the conversation
**When** the user presses `j` or `Ctrl+f`
**Then** the viewport SHOULD stay at the bottom
**And** no overflow error SHOULD occur

#### Scenario: Auto-scroll on new message

**Given** the user is not at the bottom of the conversation
**When** a new message arrives
**Then** the buffer SHOULD:
- Display a "New message" indicator
- NOT auto-scroll to the new message
- Allow user to press `G` to jump to the latest

**Given** the user is already at the bottom
**When** a new message arrives
**Then** the buffer SHOULD auto-scroll to show the new message

---

### Requirement: VISUAL Mode Text Selection

The chat buffer MUST support VISUAL mode for selecting and copying arbitrary text ranges.

#### Scenario: Enter VISUAL mode

**Given** the user is in NORMAL mode
**When** the user presses `v`
**Then** the buffer SHOULD:
- Enter VISUAL mode
- Start selection at the current cursor position
- Highlight the selected character
- Update status bar to show VISUAL mode

#### Scenario: Expand selection

**Given** the user is in VISUAL mode
**When** the user presses movement keys (`h`/`j`/`k`/`l`/`w`/`$`/etc.)
**Then** the buffer SHOULD:
- Expand the selection to the new position
- Highlight the selected text range
- Update the selection in real-time as the cursor moves

#### Scenario: Copy selection

**Given** the user has a text selection in VISUAL mode
**When** the user presses `y`
**Then** the buffer SHOULD:
- Copy the selected text to clipboard
- Return to NORMAL mode
- Display confirmation: "Copied selection to clipboard"

#### Scenario: Cancel selection

**Given** the user is in VISUAL mode
**When** the user presses `Esc` or `Ctrl+c`
**Then** the buffer SHOULD:
- Clear the selection
- Return to NORMAL mode
- Display no confirmation (silent cancel)

---

### Requirement: Long Conversation Performance

The chat buffer MUST handle long conversations (1000+ messages) without performance degradation.

#### Scenario: Lazy rendering

**Given** a conversation with 1000+ messages
**When** the buffer is rendered
**Then** the buffer SHOULD:
- Only render messages visible in the viewport
- Use a viewport or windowing technique
- NOT attempt to render all 1000+ messages at once

#### Scenario: Scroll performance

**Given** a conversation with 1000+ messages
**When** the user scrolls through the conversation
**Then** scrolling SHOULD:
- Remain responsive (no noticeable lag)
- Render at 60fps or better
- NOT cause memory leaks

#### Scenario: Message pagination

**Given** a conversation with 1000+ messages
**When** the session is loaded
**Then** the buffer SHOULD:
- Load messages in pages (e.g., 100 messages at a time)
- Load additional pages as the user scrolls near the boundary
- Cache previously loaded pages in memory

---

### Requirement: Message Metadata

The chat buffer MUST display and handle message metadata.

#### Scenario: Timestamp display

**Given** a message is rendered
**When** the message is displayed
**Then** the message SHOULD show:
- A timestamp in a subtle format (e.g., "2:30 PM" or "14:30")
- The timestamp positioned near the role indicator
- Older messages showing date if not today (e.g., "Yesterday 2:30 PM")

#### Scenario: Message length indicator

**Given** a message is very long
**When** the message is rendered
**Then** the buffer SHOULD:
- Show an ellipsis or indicator if the message is truncated
- Allow full viewing by scrolling
- NOT hard truncate message content

---

### Requirement: Markdown Handling

The chat buffer MUST handle basic markdown in text blocks without full rendering.

#### Scenario: Code block extraction

**Given** an AI response contains markdown code blocks
**When** the response is parsed into Blocks
**Then** the parser SHOULD:
- Extract fenced code blocks (```language ... ```) into CodeBlocks
- Assign sequential numbers to each CodeBlock
- Include the language identifier
- Preserve exact indentation and formatting

#### Scenario: Simple text formatting

**Given** a TextBlock contains markdown
**When** the TextBlock is rendered
**Then** the renderer MAY:
- Convert `**bold**` to bright/bold text
- Convert `` `code` `` to a different color
- Ignore other markdown (headings, links, tables, etc.)

#### Scenario: Malformed markdown handling

**Given** an AI response contains malformed markdown
**When** the response is parsed
**Then** the parser SHOULD:
- Treat malformed code blocks as plain text
- NOT crash or throw errors
- Preserve the original content as much as possible

---

## Cross-References

- **tui-framework**: Defines the top-level Model that hosts this chat-buffer sub-model
- **vim-navigation**: Defines the NORMAL/VISUAL modes used by chat-buffer operations
- **session-manager**: Provides the message data that chat-buffer renders
