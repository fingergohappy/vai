## 1. Investigation
- [x] 1.1 Reproduce current rendering in tmux and capture pane output (`tmux capture-pane`) to document the problem
- [x] 1.2 Identify where assistant bubble width is decided (`internal/chat/message_view.go`) and how block wrapping contributes

## 2. Implementation
- [x] 2.1 Introduce a shared “message bubble max width” rule (cap at 2/3 of chat pane width)
- [x] 2.2 Update assistant message rendering so bubble width is content-driven but never exceeds the cap; long content wraps
- [x] 2.3 Update user message rendering similarly (cap at 2/3) while preserving right alignment

## 3. Verification
- [x] 3.1 Run `make build`
- [x] 3.2 Smoke check in tmux session `vai` window 1: run `./build/vai`, wait, `capture-pane`
- [x] 3.3 Run `go test -v ./...` (note: may be no tests; ensure build passes)
