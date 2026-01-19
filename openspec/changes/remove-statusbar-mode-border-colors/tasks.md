# Tasks: Remove Status Bar, Add Mode Border Colors

## Task Summary

- Remove status bar component and rendering
- Add mode-based border colors for non-focused panes
- Update layout calculation to not reserve status bar space
- Clean up unused code

---

## Phase 1: Update Styles

### 1.1 Add mode border styles
- [ ] Add `NormalModeBorder` style to `internal/ui/styles.go`
- [ ] Add `InsertModeBorder` style to `internal/ui/styles.go`
- [ ] Add `VisualModeBorder` style to `internal/ui/styles.go`
- [ ] Configure: NormalBorder = gray (240), NormalBorder
- [ ] Configure: InsertBorder = green (142), NormalBorder
- [ ] Configure: VisualBorder = blue (33), NormalBorder

**Validation:** Styles compile without errors

### 1.2 Remove StatusBar style (deferred to Phase 4)
- [ ] Mark for removal: `StatusBar` style in `internal/ui/styles.go`

---

## Phase 2: Update Model Rendering

### 2.1 Add getPaneStyle helper method
- [ ] Add `getPaneStyle(isFocused bool) lipgloss.Style` method to Model
- [ ] Return `FocusedBorder` if `isFocused == true`
- [ ] Return mode border based on `m.Mode` if not focused
- [ ] Handle all three modes: Normal, Insert, Visual

**Validation:** Method compiles and returns correct styles

### 2.2 Update renderSessionPane
- [ ] Change from `if m.Focus == ui.FocusHistory` logic
- [ ] Use `m.getPaneStyle(m.Focus == ui.FocusHistory)`
- [ ] Keep width and height calculations

**Validation:** Session pane renders with correct border

### 2.3 Update renderChatPane
- [ ] Change from `if m.Focus == ui.FocusBuffer` logic
- [ ] Use `m.getPaneStyle(m.Focus == ui.FocusBuffer)`
- [ ] Keep width and height calculations

**Validation:** Chat pane renders with correct border

### 2.4 Update renderInputPane
- [ ] Change from `if m.Focus == ui.FocusInput` logic
- [ ] Use `m.getPaneStyle(m.Focus == ui.FocusInput)`
- [ ] Keep width and height calculations

**Validation:** Input pane renders with correct border

---

## Phase 3: Remove Status Bar from View

### 3.1 Remove status bar rendering
- [ ] Remove `statusBar := m.StatusBar.Render(m.Mode, m.Focus)` from View()
- [ ] Remove `statusBar` from `lipgloss.JoinVertical` call
- [ ] Update JoinVertical to only include: topSection, inputPane

**Validation:** View compiles without errors

### 3.2 Update Model struct (deferred to Phase 4)
- [ ] Mark for removal: `StatusBar *ui.StatusBar` field

---

## Phase 4: Update Layout Calculation

### 4.1 Remove statusBarHeight from layout
- [ ] Remove `statusBarHeight := 1` from `CalculateLayout`
- [ ] Update `contentHeight` calculation: `height - inputHeight` (no longer subtract statusBarHeight)
- [ ] Remove `StatusBarHeight` from Layout struct if present

**Validation:** Layout compiles and content area fills available space

### 4.2 Update pane Y positions
- [ ] Ensure SessionList and ChatBuffer start at `Y: 0` (no offset for status bar)
- [ ] Ensure InputArea starts at `Y: contentHeight`

**Validation:** Panes position correctly without gaps

---

## Phase 5: Cleanup

### 5.1 Delete status bar component
- [ ] Delete `internal/ui/statusbar.go`

### 5.2 Remove StatusBar from Styles
- [ ] Remove `StatusBar lipgloss.Style` from Styles struct
- [ ] Remove StatusBar initialization from `DefaultStyles()`

### 5.3 Remove StatusBar from Model
- [ ] Remove `StatusBar *ui.StatusBar` field from Model struct
- [ ] Remove `m.StatusBar.SetWidth(msg.Width)` call from Update()
- [ ] Remove `ui.NewStatusBar(styles)` call from `NewModel()`

### 5.4 Remove imports
- [ ] Check and remove any unused imports in modified files

---

## Phase 6: Testing and Verification

### 6.1 Build verification
- [ ] Run `make build`
- [ ] Fix any compilation errors

### 6.2 Manual testing
- [ ] Run `./build/vai`
- [ ] Verify: No status bar is displayed
- [ ] Verify: NORMAL mode shows gray borders on non-focused panes
- [ ] Press `i` to enter INSERT mode
- [ ] Verify: INSERT mode shows green borders on non-focused panes
- [ ] Verify: Input area has thick cyan border (focused)
- [ ] Press `Esc` to return to NORMAL mode
- [ ] Verify: Borders return to gray
- [ ] Press `Ctrl+w` to switch focus
- [ ] Verify: Focused pane has thick cyan border
- [ ] Verify: Non-focused panes show mode color

### 6.3 Code quality
- [ ] Run `go vet ./...`
- [ ] Fix any warnings
- [ ] Verify no unused code remains

---

## Dependencies

- **Phase 1** must be done first (add styles before using them)
- **Phase 2** depends on **Phase 1** (needs new styles)
- **Phase 3** can be done in parallel with **Phase 1-2**
- **Phase 4** depends on **Phase 3** (remove status bar before updating layout)
- **Phase 5** depends on **Phase 1-4** (cleanup after all changes)
- **Phase 6** depends on all previous phases

## Estimated Time

- **Phase 1 (Styles):** 15 minutes
- **Phase 2 (Model Rendering):** 30 minutes
- **Phase 3 (Remove Status Bar):** 15 minutes
- **Phase 4 (Layout):** 15 minutes
- **Phase 5 (Cleanup):** 15 minutes
- **Phase 6 (Testing):** 30 minutes

**Total Time:** ~2 hours

## Definition of Done

This change is complete when:
1. ✓ Status bar is completely removed from the UI
2. ✓ Mode is indicated by border colors (gray/green/blue)
3. ✓ Focused pane has thick cyan border
4. ✓ Content area uses full available height
5. ✓ All code compiles without errors
6. ✓ `go vet` passes without warnings
7. ✓ Manual testing confirms all scenarios work
8. ✓ No unused code or imports remain
