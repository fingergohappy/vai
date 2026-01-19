# Tasks: Implementation MVP - Phase 1 Basic UI

This document outlines Phase 1 tasks: creating the basic UI framework skeleton.

## Task Summary

- Implement three-pane layout with placeholder content
- Implement status bar with mode and focus display
- Implement focus switching between panes
- Implement basic mode transitions (NORMAL ↔ INSERT)
- Implement quit functionality
- Use Bubble Tea framework

---

## Phase 1.1: Foundation Setup

### 1.1.1 Update main.go
- [ ] Update `cmd/vai/main.go` to initialize Bubble Tea
- [ ] Create basic top-level Model structure
- [ ] Start the Program with alt screen
- [ ] Handle quit key (Ctrl+c)

**Validation:** `vai` command starts and shows a blank screen, exits with Ctrl+c

### 1.1.2 Create top-level Model
- [ ] Define `app.Model` struct with Mode, Focus, Quitting fields
- [ ] Add placeholder sub-models: Session, Buffer, Input
- [ ] Implement `Init()` function
- [ ] Implement basic `Update()` for quit handling
- [ ] Implement basic `View()` returning placeholder

**Validation:** Compiles and runs without errors

---

## Phase 1.2: Three-Pane Layout

### 1.2.1 Implement layout calculation
- [ ] Update `internal/ui/layout.go`
- [ ] Implement `CalculateLayout(tea.WindowSizeMsg) Layout`
- [ ] Calculate pane dimensions (20%/80% split)
- [ ] Calculate pane positions (x, y for each)

**Validation:** Layout function returns correct dimensions

### 1.2.2 Implement pane rendering in View
- [ ] Update `app.Model.View()` to render three panes
- [ ] Render session list pane (left)
- [ ] Render chat buffer pane (right)
- [ ] Render input area pane (bottom)
- [ ] Use Lipgloss for borders and styling

**Validation:** Three panes are visible with borders

### 1.2.3 Add pane placeholder content
- [ ] Add placeholder text to session list
- [ ] Add placeholder text to chat buffer
- [ ] Add placeholder text to input area
- [ ] Use styled boxes for visual clarity

**Validation:** All panes show placeholder content

---

## Phase 1.3: Status Bar

### 1.3.1 Create status bar component
- [ ] Update `internal/ui/statusbar.go`
- [ ] Implement `Render(mode, focus) string` method
- [ ] Add mode-colored text
- [ ] Add focus indicator
- [ ] Add separator and title

**Validation:** Status bar displays at top of screen

### 1.3.2 Integrate status bar into View
- [ ] Call status bar render in top-level View
- [ ] Pass current Mode and Focus
- [ ] Position at top (y=0)
- [ ] Full width styling

**Validation:** Status bar shows current state at top

---

## Phase 1.4: Focus Management

### 1.4.1 Implement Focus type and switching
- [ ] Verify `internal/ui/focus.go` has Focus type
- [ ] Implement `Next()` and `Prev()` methods
- [ ] Add visual focus indication

**Validation:** Focus type works correctly

### 1.4.2 Implement pane switching in Update
- [ ] Handle `Ctrl+w` followed by h/j/k/l in Update
- [ ] Update Model.Focus field
- [ ] Trigger re-render

**Validation:** Can switch focus between three panes

### 1.4.3 Add visual focus indication
- [ ] Apply different border styles based on focus
- [ ] Use thicker/brighter border for focused pane
- [ ] Use dimmer border for unfocused panes

**Validation:** Focused pane has distinct appearance

---

## Phase 1.5: Mode System

### 1.5.1 Implement Mode type
- [ ] Verify `internal/vim/mode.go` has Mode type
- [ ] Define NORMAL and INSERT constants
- [ ] Implement `String()` method

**Validation:** Mode type compiles and works

### 1.5.2 Implement mode transitions in Update
- [ ] Handle `i` and `a` keys to enter INSERT mode
- [ ] Move focus to input area when entering INSERT
- [ ] Handle `Esc` to return to NORMAL mode
- [ ] Move focus back to buffer when exiting INSERT

**Validation:** Can switch between NORMAL and INSERT modes

### 1.5.3 Update status bar with mode
- [ ] Pass mode to status bar render
- [ ] Display mode with appropriate color
- [ ] Update on mode change

**Validation:** Status bar shows current mode

---

## Phase 1.6: Quit Functionality

### 1.6.1 Implement Ctrl+q quit
- [ ] Handle `Ctrl+q` in Update function
- [ ] Set quitting flag or show confirmation
- [ ] Quit on second press or after timeout
- [ ] Return `tea.Quit` command

**Validation:** `Ctrl+q` (twice) quits application

### 1.6.2 Implement Ctrl+c quit
- [ ] Handle `Ctrl+q` in Update function
- [ ] Quit immediately
- [ ] Return `tea.Quit` command

**Validation:** `Ctrl+c` quits application immediately

---

## Phase 1.7: Styling

### 1.7.1 Define default styles
- [ ] Update `internal/ui/styles.go`
- [ ] Define border styles
- [ ] Define mode colors
- [ ] Define focus styles

**Validation:** Styles apply correctly to UI

### 1.7.2 Apply Lipgloss styling
- [ ] Use Lipgloss for all text styling
- [ ] Use Lipgloss for borders
- [ ] Use Lipgloss for colors
- [ ] Ensure ANSI codes work in terminal

**Validation:** Styled UI displays correctly

---

## Phase 1.8: Window Size Handling

### 1.8.1 Handle window resize messages
- [ ] Handle `tea.WindowSizeMsg` in Update
- [ ] Recalculate layout on resize
- [ ] Update stored dimensions
- [ ] Trigger re-render

**Validation:** Resizing terminal updates layout correctly

### 1.8.2 Set minimum terminal size
- [ ] Add check for minimum size (80x24)
- [ ] Show error if terminal too small
- [ ] Exit gracefully if can't fit content

**Validation:** Error message on small terminal

---

## Phase 1.9: Placeholder Components

### 1.9.1 Update session.Model placeholder
- [ ] Implement `session.Model.View()` with placeholder content
- [ ] Show example sessions
- [ ] Add "(TODO)" note

**Validation:** Session list shows placeholder

### 1.9.2 Update chat.Model placeholder
- [ ] Implement `chat.Model.View()` with placeholder content
- [ ] Show welcome message
- [ ] Add "(TODO)" note

**Validation:** Chat buffer shows placeholder

### 1.9.3 Update input.Model placeholder
- [ ] Implement `input.Model.View()` with placeholder
- [ ] Show "[Type your message...]"
- [ ] Add visual input box

**Validation:** Input area shows placeholder

---

## Phase 1.10: Integration and Testing

### 1.10.1 Update main.go with real Model
- [ ] Import all necessary packages
- [ ] Create Model with proper initialization
- [ ] Set up Program with alt screen
- [ ] Add error handling

**Validation:** Application starts without errors

### 1.10.2 Test all interactions
- [ ] Test startup: UI displays correctly
- [ ] Test focus switching: `Ctrl+w h/j/k/l`
- [ ] Test mode switching: `i`, `a`, `Esc`
- [ ] Test status bar updates
- [ ] Test quit: `Ctrl+c`, `Ctrl+q`
- [ ] Test resize: change terminal size

**Validation:** All interactions work as expected

### 1.10.3 Build and run verification
- [ ] Run `go build ./cmd/vai`
- [ ] Run `./vai` and verify UI
- [ ] Run `go vet ./...`
- [ ] Fix any warnings

**Validation:** Clean build and run

---

## Phase 1.11: Documentation

### 1.11.1 Update README.md
- [ ] Document Phase 1 capabilities
- [ ] Add screenshot description
- [ ] Document keybindings for Phase 1
- [ ] Note future phases

**Validation:** README reflects current state

### 1.11.2 Create Phase 1 completion note
- [ ] Document what Phase 1 delivers
- [ ] Document what's next
- [ ] Add to CHANGELOG or release notes

**Validation:** Progress is documented

---

## Dependencies

- **1.1** must be done first (foundation)
- **1.2-1.7** can be done in parallel (independent features)
- **1.8** depends on **1.2** (needs layout)
- **1.9** can be done alongside other phases
- **1.10** depends on all previous phases
- **1.11** can be done during or after implementation

## Estimated Time

- **Foundation (1.1):** 1 hour
- **Layout (1.2):** 2-3 hours
- **Status bar (1.3):** 1-2 hours
- **Focus (1.4):** 1-2 hours
- **Mode (1.5):** 1-2 hours
- **Quit (1.6):** 30 minutes
- **Styling (1.7):** 1-2 hours
- **Resize (1.8):** 1 hour
- **Placeholders (1.9):** 1-2 hours
- **Integration (1.10):** 2-3 hours
- **Documentation (1.11):** 1 hour

**Total Phase 1 Time:** ~15-20 hours

## Definition of Done

Phase 1 is complete when:
1. ✓ Application starts and shows three-pane layout
2. ✓ Status bar displays mode and focus
3. ✓ Focus switches with `Ctrl+w h/j/k/l`
4. ✓ Mode switches with `i`/`a`/`Esc`
5. ✓ All panes show placeholder content
6. ✓ `Ctrl+q` and `Ctrl+c` quit the application
7. ✓ Terminal resize updates layout
8. ✓ `go build` and `go vet` pass
9. ✓ README documents current capabilities
10. ✓ User can verify "the frame is drawn"

## Next Steps

After Phase 1 completion:
- **Phase 2:** Implement real input functionality
- **Phase 3:** Implement navigation and scrolling
- **Phase 4:** Implement session management
- **Phase 5:** Implement Markdown rendering
- **Phase 6:** Implement code block operations

Each phase will be a separate change proposal building on this foundation.
