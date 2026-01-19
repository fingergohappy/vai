# Architecture

## Overview

vai is a terminal UI (TUI) application built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), following the Model-Update-View architecture pattern.

## Design Principles

1. **Keyboard First** - All operations accessible without mouse
2. **Vim Thinking** - Familiar modal editing for Vim users
3. **Explicit State** - Mode and focus always visible
4. **Code Blocks as First-Class Citizens** - Easy navigation and copying
5. **Clean Separation** - Packages organized by feature domain

## Architecture

### Top-Level Model

The application is driven by a top-level Model that combines sub-models:

```go
type Model struct {
    Mode   vim.Mode       // Current mode (NORMAL/INSERT/VISUAL)
    Focus  ui.Focus       // Current focus (History/Buffer/Input)

    Session session.Model // Session list sub-model
    Chat    chat.Model    // Chat buffer sub-model
    Input   input.Model   // Input area sub-model
}
```

### Message Flow

```
User Input → tea.KeyMsg → Router → Sub-models → tea.Cmd → Actions
                ↓
         Mode + Focus Check
```

### Package Organization

```
internal/
├── app/        # Top-level model, message routing
├── vim/        # Mode system, key routing
├── ui/         # Shared UI components (layout, styles)
├── chat/       # Chat buffer, message rendering
├── session/    # Session persistence, list
├── input/      # Input area with Vim movement
├── clipboard/  # Cross-platform clipboard
└── config/     # Configuration management
```

## Import Graph

```
        cmd/vai
           │
        internal/app
           │
    ┌──────┼──────┐
    │      │      │
  vim     ui    chat
           │      │
        input  clipboard
           │
        session
```

## Mode System

### Modes

- **NORMAL** - Navigation and commands
- **INSERT** - Text input in input area
- **VISUAL** - Text selection in chat buffer

### Mode-Focus Compatibility

| Focus    | NORMAL | INSERT | VISUAL |
|----------|--------|--------|--------|
| History  | ✅     | ❌     | ❌     |
| Buffer   | ✅     | ❌     | ✅     |
| Input    | ❌     | ✅     | ❌     |

## UI Layout

```
┌─────────────────────────────────────────┐
│ Status Bar: NORMAL | Buffer              │
├─────────────┬───────────────────────────┤
│  Session    │  Chat Buffer              │
│  List       │                           │
│             │                           │
├─────────────┴───────────────────────────┤
│  Input Area                            │
└─────────────────────────────────────────┘
```

## Key Components

### Chat Buffer

- Renders messages with structured blocks (text, code)
- Handles viewport scrolling for long conversations
- Supports code block navigation (`]c`, `[c`)
- Supports VISUAL mode selection

### Session Manager

- Manages chat session persistence
- Loads/sessions from `~/.local/share/vai/sessions/`
- Renders session list with metadata

### Input Area

- Wraps `bubbles.TextArea` for multi-line input
- Implements Vim-style movement in INSERT mode
- Sends messages on Enter

### Clipboard

- Cross-platform support (macOS `pbcopy`, Linux `xclip`/`wl-copy`)
- Graceful fallback when unavailable

## Data Models

### Message

```go
type Message struct {
    ID        string
    Role      Role      // "user" or "assistant"
    Blocks    []Block
    CreatedAt time.Time
}
```

### Block

```go
type Block interface {
    Kind() BlockType
    Render(width int) string
}

type TextBlock struct { Text string }
type CodeBlock struct { Lang string; Lines []string; Number int }
```

## Performance Considerations

- **Viewport-based rendering** - Only render visible messages
- **Lazy loading** - Load messages in pages
- **Efficient updates** - Bubble Tea's Elm architecture ensures minimal redraws

## Future Extensions

- Streaming responses
- Multiple AI providers
- Plugin system
- Custom themes
- Advanced markdown rendering
