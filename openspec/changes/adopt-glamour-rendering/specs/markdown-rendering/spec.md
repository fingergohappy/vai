# Spec: Markdown Rendering

**Status:** Active
**Version:** 1.0.0
**Owners:** vai project

## Overview

This spec defines Markdown rendering using Glamour for display and regex-based code block extraction for copy operations.

## ADDED Requirements

### Requirement: Glamour Integration

The application MUST use Glamour for rendering Markdown content in the chat buffer.

#### Scenario: Render Markdown with Glamour

**Given** an AI assistant response containing Markdown
**When** the message is rendered
**Then** the application SHOULD:
- Use Glamour to render the Markdown to ANSI text
- Apply syntax highlighting to code blocks
- Support common Markdown formatting (headers, bold, lists, etc.)
- Preserve the original Markdown source for searching

#### Scenario: Glamour initialization

**Given** the application starts
**When** the Markdown renderer is initialized
**Then** it SHOULD:
- Create a Glamour TermRenderer instance
- Apply a default style (dark or light based on config)
- Set word wrap based on terminal width
- Fall back to plain text if Glamour fails

#### Scenario: Custom styles

**Given** the user has configured a custom theme
**When** Glamour renders Markdown
**Then** it SHOULD:
- Map the theme colors to Glamour stylesheet
- Apply code block colors from theme
- Apply text colors from theme
- Use default Glamour styles for undefined elements

---

### Requirement: Code Block Extraction

The application MUST extract code blocks from Markdown using regex for copy operations.

#### Scenario: Extract fenced code blocks

**Given** an AI response contains fenced code blocks
**When** the message is processed
**Then** the extractor SHOULD:
- Find all fenced code blocks (````lang ... ```)
- Extract the language identifier
- Extract the code content
- Assign sequential numbers starting from 1
- Store the position in the original text

#### Scenario: Handle code blocks without language

**Given** an AI response contains code blocks without language identifier
**When** the code block is extracted
**Then** the extractor SHOULD:
- Treat it as a valid code block
- Set language to empty string or "text"
- Still assign it a number for navigation

#### Scenario: Handle malformed code blocks

**Given** an AI response contains unclosed or malformed code blocks
**When** the extraction runs
**Then** the extractor SHOULD:
- Ignore incomplete code blocks
- Not throw errors
- Continue processing other blocks
- Return a valid (possibly empty) index

#### Scenario: Get code block by number

**Given** a message has been processed
**When** `GetBlock(n)` is called
**Then** it SHOULD return:
- The nth code block (1-indexed)
- The language identifier
- The raw code content (without backticks)
- Nil if n is out of range

---

### Requirement: Code Block Navigation

The application MUST support jumping to code blocks using the extracted index.

#### Scenario: Jump to next code block

**Given** the user is in NORMAL mode viewing a message
**When** the user presses `]c`
**Then** the buffer SHOULD:
- Find the next code block using the index
- Calculate its line position in the rendered content
- Scroll the viewport to show the code block
- Highlight the code block briefly

#### Scenario: Jump to previous code block

**Given** the user is in NORMAL mode viewing a message
**When** the user presses `[c`
**Then** the buffer SHOULD:
- Find the previous code block using the index
- Calculate its line position in the rendered content
- Scroll the viewport to show the code block
- Wrap to the last block if at the first

#### Scenario: Jump to specific code block

**Given** the user presses `N]c` (count)
**When** the buffer processes the command
**Then** it SHOULD:
- Jump to the Nth code block (1-indexed)
- Clamp to the valid range
- Show an error if no code blocks exist

#### Scenario: Track current code block

**Given** the user has navigated to a code block
**When** the current position changes
**Then** the buffer SHOULD:
- Track which code block is currently focused
- Display the code block number in UI
- Use this for copy operations

---

### Requirement: Code Block Copying

The application MUST copy code blocks using the extracted (non-rendered) content.

#### Scenario: Copy current code block

**Given** the user has navigated to a code block
**When** the user presses `yc`
**Then** the application SHOULD:
- Get the current code block from the index
- Copy the raw code content (without backticks or syntax highlighting)
- Display confirmation with block number
- Use the system clipboard (pbcopy/xclip/wl-copy)

#### Scenario: Copy specific code block by number

**Given** the user presses `yNc` (count)
**When** the copy operation executes
**Then** it SHOULD:
- Get the Nth code block from the index
- Copy the raw code content
- Display confirmation
- Show error if N is out of range

#### Scenario: Copy with formatting removal

**Given** a code block is being copied
**When** the copy operation executes
**Then** the content SHOULD:
- NOT include the ``` delimiters
- NOT include the language identifier line
- NOT include ANSI color codes
- Be plain text suitable for pasting into code

---

### Requirement: Message Structure Integration

The application MUST integrate Glamour rendering and code block extraction into the Message structure.

#### Scenario: Message rendering pipeline

**Given** a new AI response arrives
**When** the message is added to the chat buffer
**Then** the application SHOULD:
- Store the original Markdown in `RawMarkdown`
- Render with Glamour and store in `Rendered`
- Extract code blocks and store in `BlockIndex`
- Store all three for different purposes

#### Scenario: Display uses rendered content

**Given** a message is being displayed in the chat buffer
**When** the View function renders the message
**Then** it SHOULD use:
- The Glamour-rendered ANSI text
- NOT the raw Markdown
- This ensures syntax highlighting and formatting

#### Scenario: Copy uses extracted content

**Given** the user copies a code block
**When** the copy operation executes
**Then** it SHOULD use:
- The plain text from `BlockIndex`
- NOT the rendered content
- This avoids copying ANSI codes

---

### Requirement: Performance

The application MUST render and extract efficiently.

#### Scenario: Lazy rendering

**Given** a conversation with many messages
**When** the chat buffer is displayed
**Then** only visible messages SHOULD be rendered

#### Scenario: Rendering cache

**Given** a message has been rendered once
**When** it needs to be displayed again
**Then** the cached render SHOULD be used
**And** the message SHOULD NOT be re-rendered

#### Scenario: Extraction performance

**Given** a message with many code blocks
**When** the extractor processes it
**Then** extraction SHOULD:
- Complete in O(n) time where n is text length
- Use a single regex pass
- Not build a full AST

---

### Requirement: Error Handling

The application MUST handle rendering and extraction errors gracefully.

#### Scenario: Glamour rendering failure

**Given** Glamour fails to render Markdown
**When** the error occurs
**Then** the application SHOULD:
- Fall back to displaying the raw Markdown
- Log the error for debugging
- Not crash or display an error to the user

#### Scenario: Extraction with invalid regex

**Given** the regex pattern fails to compile
**When** the extractor is initialized
**Then** the application SHOULD:
- Use a pre-validated regex pattern
- Never encounter runtime regex errors
- Return empty index if extraction fails

---

## MODIFIED Requirements

### Requirement: Chat Buffer Display (Modified)

The chat buffer MUST display Glamour-rendered Markdown.

#### Scenario: Display rendered messages

**Given** a message with Glamour-rendered content
**When** the buffer is rendered
**Then** it SHOULD display:
- The ANSI-formatted text from Glamour
- Syntax-highlighted code blocks
- Formatted headers, lists, etc.
- NOT the raw Markdown source

---

## Cross-References

- **chat-buffer**: This spec updates chat buffer to use Glamour rendering
- **vim-navigation**: Code block navigation (']c', '[c') uses the extracted index
- **design.md**: See design document for architecture details

---

## Implementation Notes

### Regex Pattern

```go
var codeBlockRE = regexp.MustCompile(
    "^```([a-zA-Z0-9+_-]*)?\n" + // Language (optional)
    "([\\s\\S]*?)" +              // Content (non-greedy)
    "\n```$"                       // Closing
)
```

### Glamour Basic Usage

```go
import "github.com/charmbracelet/glamour"

renderer, _ := glamour.NewTermRenderer(
    glamour.WithAutoStyle(),
    glamour.WithWordWrap(width),
)

out, _ := renderer.Render(markdown)
```
