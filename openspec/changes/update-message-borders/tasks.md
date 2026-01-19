## 1. Investigation
- [x] 1.1 Capture current message style in tmux (vai:1) for baseline
- [x] 1.2 Confirm existing message width cap behavior remains correct

## 2. Implementation
- [x] 2.1 Update message border styles to a full box border (top/right/bottom/left) with left/right padding
- [x] 2.2 Implement a helper to render a centered title inside the top border ("AI" / "You")
- [x] 2.3 Implement dynamic bubble width:
  - bubbleWidth = min(paneWidth*2/3, contentWidth + padding + borderFrame)
  - if content exceeds the limit, wrap until each rendered line fits
  - else bubbleWidth shrinks to fit content
- [x] 2.4 Keep AI left-aligned and user right-aligned within the chat pane

## 3. Verification
- [x] 3.1 Run `make build`
- [x] 3.2 tmux smoke check (vai:1): run `./build/vai`, `capture-pane`
- [x] 3.3 Run `go test -v ./...`
