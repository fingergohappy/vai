# Change: Fix AI message width

## Why
AI message bubbles currently render at (or close to) the full chat pane width, even when the message text is short. This creates excessive horizontal whitespace and makes the chat panel feel visually unbalanced.

## What Changes
- Constrain assistant (AI) message bubble width so the bubble wraps closely around the rendered content, up to a maximum of 2/3 of the chat pane width (wrap long lines).
- Constrain user message bubble width similarly (max 2/3 of the chat pane width) while preserving right alignment.
- Keep overall chat pane layout unchanged; only adjust per-message rendering.
- Validate the visual result via tmux pane capture (non-interactive smoke check).

## Impact
- Affected specs: `basic-ui` (visual rendering details, message bubble presentation)
- Affected code:
  - `internal/chat/message_view.go`
  - Potentially `internal/chat/block.go` (if width calculation requires helpers)

## Non-Goals
- Changing the overall three-pane layout widths (session list / chat buffer / input area).
- Redesigning styles, colors, or adding new message features.
