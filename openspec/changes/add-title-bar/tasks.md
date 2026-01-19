# Tasks: Add Title Bar for Session Name Display

## Implementation Tasks

### 1. Create Title Bar Component
- [ ] Create `internal/ui/titlebar.go`
- [ ] Implement `TitleBar` struct with `styles` field
- [ ] Implement `NewTitleBar(styles *Styles) *TitleBar` constructor
- [ ] Implement `Render(sessionTitle string) string` method
- [ ] Implement `SetWidth(width int)` method

### 2. Add Title Bar Styles
- [ ] Add `TitleBar lipgloss.Style` field to `Styles` struct in `internal/ui/styles.go`
- [ ] Add title bar style initialization in `DefaultStyles()`:
  - Bold text
  - Center alignment
  - Foreground color 252 (white)
  - Background color 235 (dark gray)

### 3. Update Layout Calculation
- [ ] Add `TitleBar PaneLayout` field to `Layout` struct in `internal/ui/layout.go`
- [ ] Add `titleBarHeight := 1` variable in `CalculateLayout()`
- [ ] Update `contentHeight` calculation: `contentHeight = height - titleBarHeight - inputHeight`
- [ ] Set title bar position: `X: 0, Y: 0, Width: width, Height: titleBarHeight`
- [ ] Update `SessionList.Y` to `titleBarHeight`
- [ ] Update `ChatBuffer.Y` to `titleBarHeight`
- [ ] Update `InputArea.Y` to `titleBarHeight + contentHeight`

### 4. Integrate Title Bar into Model
- [ ] Add `TitleBar *ui.TitleBar` field to `Model` struct in `internal/app/model.go`
- [ ] Initialize title bar in `NewModel()`: `titleBar := ui.NewTitleBar(styles)`
- [ ] Add `TitleBar: titleBar` to Model initialization
- [ ] Add `m.TitleBar.SetWidth(msg.Width)` in `Update()` method's `WindowSizeMsg` case

### 5. Add Title Bar Rendering
- [ ] Add `renderTitleBar() string` method to `Model` in `internal/app/model.go`
- [ ] Implement logic to get current session title
- [ ] Call `m.TitleBar.Render(currentTitle)` in the method
- [ ] Add title bar to `View()` method's vertical join at the top

### 6. Add Session Title Access
- [ ] Add `GetCurrentTitle() string` method to `session.Model` in `internal/session/`
- [ ] Handle empty sessions list case (return "New Chat")
- [ ] Handle valid index case (return `Sessions[currentIndex].Title`)
- [ ] Handle out of bounds case (return "New Chat")

### 7. Update Session List Pane Height
- [ ] Verify `renderSessionPane()` uses correct height from layout
- [ ] Ensure pane content fits within calculated height

### 8. Update Chat Buffer Pane Height
- [ ] Verify `renderChatPane()` uses correct height from layout
- [ ] Ensure pane content fits within calculated height

### 9. Verification and Testing
- [ ] Run `make build` to verify compilation
- [ ] Run application to verify title bar displays
- [ ] Verify title bar shows "Sessions - New Chat" by default
- [ ] Verify title bar centers text properly
- [ ] Test with long session names to ensure no truncation
- [ ] Test terminal resize to verify title bar width adapts
- [ ] Verify panes are positioned correctly below title bar
- [ ] Verify mode border colors still work correctly

## Dependencies

- Requires `remove-statusbar-mode-border-colors` change to be completed
- All tasks should be completed in numerical order

## Validation

After completing all tasks:
1. Build succeeds without errors
2. Title bar is visible at top of screen
3. Title bar text is centered
4. Title bar shows "Sessions - [session name]"
5. Long session names display fully
6. Panes are positioned below title bar
7. Terminal resize works correctly
8. Mode border colors function as before
