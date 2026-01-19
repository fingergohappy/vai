# Change: Update Message Borders With Center Labels

## Why
Current message rendering uses a single right-side border and an underline separator. This reads like a divider rather than a "message bubble" and the role label (AI/You) is embedded as text, not part of the border.

## What Changes
- Replace the underline/side-border message style with a boxed border that wraps the message content.
- Render the role label ("AI" / "You") centered in the top border of the box.
- Add horizontal padding between the border and text (left/right).
- Apply dynamic bubble width:
  - If (text width + padding + border frame) exceeds 2/3 of the chat pane width, wrap content and set bubble width to exactly 2/3 of the pane.
  - If it does not exceed 2/3, set bubble width to (text width + padding + border frame).
- Preserve role-based color distinction (AI blue, user green) and alignment (AI left, user right).

## Impact
- Affected spec: `chat-message-styling`
- Affected code:
  - `internal/chat/message_view.go`

## Non-Goals
- Changing pane borders, layout, or focus styling.
- Changing keybindings or input behavior.
