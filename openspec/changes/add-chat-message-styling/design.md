# Design: Chat Message Styling Sub-Models

## Overview

本设计描述如何为聊天面板创建消息样式子系统，通过边框颜色、对齐方式和标签来区分用户消息和AI消息。

## Visual Design

### Current State (Placeholder)

```
┌──────────────┬──────────────────────────┐
│ [Sessions]   │ [Chat Buffer]            │
│              │                          │
│              │ This is a placeholder.   │
│              │ Real implementation...   │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │
└─────────────────────────────────────────┘
```

### New State (With Message Styling)

```
┌──────────────┬──────────────────────────┐
│ [Sessions]   │                          │
│              │ AI: Hello! How can I help?│  ← 蓝色边框，左对齐
│              │ ┌──────────────────────┐ │
│              │ │                      │ │
│              │ └──────────────────────┘ │
│              │                          │
│              │    ┌──────────────────┐  │
│              │You:│ What is the time? │  │  ← 绿色边框，右对齐
│              │    └──────────────────┘  │
│              │                          │
├──────────────┴──────────────────────────┤
│ [Type your message...]                 │
└─────────────────────────────────────────┘
```

### Message Styling Details

| Attribute     | User Message               | AI Message                 |
| ------------- | -------------------------- | -------------------------- |
| Border Color  | Green (142)                | Blue (33)                  |
| Border Style  | Normal border              | Normal border              |
| Alignment     | Right-aligned              | Left-aligned               |
| Label         | "You:"                     | "AI:"                      |
| Label Color   | Green (142)                | Blue (33)                  |
| Content Color | Gray (244)                 | White (252)                |
| Background    | Dark gray (235)            | Dark gray (235)            |

## Implementation Architecture

### 1. ChatMessage Component (chat/message_view.go)

```go
package chat

import (
    "github.com/charmbracelet/lipgloss"
)

// ChatMessage renders a single chat message with role-specific styling.
type ChatMessage struct {
    styles *lipgloss.Style
}

// NewChatMessage creates a new chat message component.
func NewChatMessage() *ChatMessage {
    return &ChatMessage{}
}

// Render renders a message with appropriate styling based on role.
func (cm *ChatMessage) Render(msg Message, maxWidth int) string {
    switch msg.Role {
    case RoleUser:
        return cm.renderUserMessage(msg, maxWidth)
    case RoleAssistant:
        return cm.renderAssistantMessage(msg, maxWidth)
    default:
        return cm.renderAssistantMessage(msg, maxWidth)
    }
}

// renderUserMessage renders a user message with green border, right-aligned.
func (cm *ChatMessage) renderUserMessage(msg Message, maxWidth int) string {
    // Style for user message border
    borderStyle := lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("142")). // Green
        Padding(1).
        Width(maxWidth - 4) // Account for border and padding

    // Style for "You:" label
    labelStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("142")). // Green
        Bold(true)

    // Style for message content
    contentStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("244")). // Gray
        Width(maxWidth - 8)

    // Build message
    var content strings.Builder
    content.WriteString(labelStyle.Render("You:"))

    // Render blocks
    for _, block := range msg.Blocks {
        content.WriteString("\n")
        content.WriteString(contentStyle.Render(block.Text))
    }

    // Apply border and right-align
    messageContent := borderStyle.Render(content.String())
    return lipgloss.NewStyle().
        Align(lipgloss.Right).
        Width(maxWidth).
        Render(messageContent)
}

// renderAssistantMessage renders an AI message with blue border, left-aligned.
func (cm *ChatMessage) renderAssistantMessage(msg Message, maxWidth int) string {
    // Style for AI message border
    borderStyle := lipgloss.NewStyle().
        Border(lipgloss.NormalBorder()).
        BorderForeground(lipgloss.Color("33")). // Blue
        Padding(1).
        Width(maxWidth - 4)

    // Style for "AI:" label
    labelStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("33")). // Blue
        Bold(true)

    // Style for message content
    contentStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("252")). // White
        Width(maxWidth - 8)

    // Build message
    var content strings.Builder
    content.WriteString(labelStyle.Render("AI:"))

    // Render blocks
    for _, block := range msg.Blocks {
        content.WriteString("\n")
        content.WriteString(contentStyle.Render(block.Text))
    }

    // Apply border and left-align
    messageContent := borderStyle.Render(content.String())
    return lipgloss.NewStyle().
        Align(lipgloss.Left).
        Width(maxWidth).
        Render(messageContent)
}
```

### 2. Styles Update (ui/styles.go)

```go
type Styles struct {
    // ... existing fields ...

    // Message border styles
    UserMessageBorder      lipgloss.Style
    AssistantMessageBorder lipgloss.Style

    // Message labels
    UserLabel      lipgloss.Style
    AssistantLabel lipgloss.Style

    // Message content
    UserMessageContent      lipgloss.Style
    AssistantMessageContent lipgloss.Style
}

func DefaultStyles() *Styles {
    return &Styles{
        // ... existing styles ...

        // User message styles (green)
        UserMessageBorder: lipgloss.NewStyle().
            Border(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("142")), // Green

        UserLabel: lipgloss.NewStyle().
            Foreground(lipgloss.Color("142")). // Green
            Bold(true).

        UserMessageContent: lipgloss.NewStyle().
            Foreground(lipgloss.Color("244")), // Gray

        // Assistant message styles (blue)
        AssistantMessageBorder: lipgloss.NewStyle().
            Border(lipgloss.NormalBorder()).
            BorderForeground(lipgloss.Color("33")), // Blue

        AssistantLabel: lipgloss.NewStyle().
            Foreground(lipgloss.Color("33")). // Blue
            Bold(true).

        AssistantMessageContent: lipgloss.NewStyle().
            Foreground(lipgloss.Color("252")), // White
    }
}
```

### 3. Buffer Update (chat/buffer.go)

```go
type Model struct {
    // ... existing fields ...

    // NEW: Message renderer
    messageRenderer *ChatMessage
}

func NewModel() Model {
    return Model{
        Messages:        []Message{},
        ViewportOffset:  0,
        CursorLine:      0,
        Selection:       Selection{Active: false},
        messageRenderer: NewChatMessage(), // NEW
    }
}

// View renders all messages in the buffer.
func (m Model) View() string {
    if len(m.Messages) == 0 {
        return "  [Chat Buffer]\n\n" +
            "  No messages yet.\n" +
            "  Type a message to start chatting..."
    }

    var renderedMessages []string

    // Render each message
    for _, msg := range m.Messages {
        rendered := m.messageRenderer.Render(msg, m.Width)
        renderedMessages = append(renderedMessages, rendered)
    }

    // Join messages with spacing
    return strings.Join(renderedMessages, "\n\n")
}
```

## Layout Calculation

### Message Width Calculation

```
Max message width = ChatBuffer width - padding - border
                  = m.Width - 8 (4 for left border, 4 for right border)
```

### Message Stacking

Messages are stacked vertically with 2 empty lines between them for visual separation:

```
[AI Message]
              ← 2 empty lines
              ← 2 empty lines
          [User Message]
              ← 2 empty lines
              ← 2 empty lines
[AI Message]
```

## Color Scheme

| Element           | User Color | AI Color |
| ----------------- | ---------- | -------- |
| Border            | 142 (Green) | 33 (Blue) |
| Label             | 142 (Green) | 33 (Blue) |
| Content Text      | 244 (Gray)  | 252 (White) |
| Background        | 235 (Dark Gray) | 235 (Dark Gray) |

## Accessibility Considerations

1. **Color blindness**: Use both color AND label text ("You:"/ "AI:") for distinction
2. **High contrast**: Ensure text colors have sufficient contrast against background
3. **Text labels**: Always include role label, not just color

## Migration Path

### Phase 1: Create message renderer
1. Create `internal/chat/message_view.go`
2. Implement `ChatMessage` component
3. Implement `renderUserMessage()` and `renderAssistantMessage()`

### Phase 2: Update styles
1. Add message border styles to `ui/styles.go`
2. Add message label styles
3. Add message content styles

### Phase 3: Integrate into buffer
1. Add `messageRenderer` field to `chat.Model`
2. Update `Model.View()` to use message renderer
3. Handle empty state

### Phase 4: Test and refine
1. Test with single messages
2. Test with multiple messages
3. Test with long content
4. Test with code blocks
5. Verify responsive behavior

## Validation Criteria

The implementation should:
1. ✅ Display user messages with green border and right alignment
2. ✅ Display AI messages with blue border and left alignment
3. ✅ Show "You:" label on user messages in green
4. ✅ Show "AI:" label on AI messages in blue
5. ✅ Stack messages vertically with proper spacing
6. ✅ Handle empty message list gracefully
7. ✅ Support multi-line content
8. ✅ Adapt to terminal width changes
9. ✅ Maintain consistency with existing UI style
10. ✅ Not conflict with mode-based pane border colors

## Edge Cases

1. **Empty messages**: Display placeholder text in chat buffer
2. **Very long messages**: Wrap text to fit available width
3. **Empty content blocks**: Skip rendering of empty blocks
4. **Terminal too narrow**: Ensure minimum width for readability
5. **Rapid message addition**: Support real-time message updates

## Future Enhancements

1. **Focus indication**: Highlight currently selected message
2. **Timestamp display**: Show message creation time
3. **Message actions**: Add copy/delete buttons
4. **Code block syntax highlighting**: Enhanced code rendering
5. **Markdown support**: Rich text formatting
6. **Message grouping**: Group consecutive messages from same role
