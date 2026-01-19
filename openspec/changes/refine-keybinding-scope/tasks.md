# Tasks: Refine Keybinding Scope and Remove Command Mode

This document outlines the implementation tasks for refining the keybinding system. This is a simplification change that removes complexity from the initial implementation.

## Task Summary

- **Remove:** Command mode (`:`) and all related functionality
- **Add:** Vim-style movement keys in INSERT mode for input area
- **Clarify:** Code block operations only work when focus is on chat buffer
- **Simplify:** Key routing logic by reducing mode-focus combinations

---

## Phase 1: Remove Command Mode

### 1.1 Remove command mode from vim-navigation spec
- [ ] Update vim-navigation spec to mark command mode as REMOVED
- [ ] Document alternative keybindings for former commands:
  - `:q` → `Ctrl+q` (quit)
  - `:new` → `Ctrl+t` (new session)
  - `:delete` → `dd` in session list
  - `:rename` → `r` in session list
  - `:help` → `?` key

**Validation:** Spec reflects that command mode is removed

### 1.2 Add global shortcut keybindings
- [ ] Implement `Ctrl+q` for quit application
  - First press: show confirmation "Press again to quit"
  - Second press within 2 seconds: quit immediately
- [ ] Implement `Ctrl+t` for create new session
- [ ] Implement `?` for show help overlay
- [ ] Ensure these work in NORMAL mode regardless of focus

**Validation:** Can quit, create session, and show help without command mode

---

## Phase 2: INSERT Mode Vim Movement

### 2.1 Add Vim movement keys to input area
- [ ] Implement `h` / `Ctrl+b` for move cursor left
- [ ] Implement `l` / `Ctrl+f` for move cursor right
- [ ] Implement `w` for move to next word
- [ ] Implement `b` for move to previous word
- [ ] Implement `e` for move to end of current word
- [ ] Implement `0` / `Ctrl+a` for move to line start
- [ ] Implement `$` / `Ctrl+e` for move to line end

**Validation:** Can move cursor using Vim keys while in INSERT mode

### 2.2 Add Vim editing keys to input area
- [ ] Implement `Ctrl+d` for delete character at cursor
- [ ] Implement `Ctrl+w` for delete word before cursor
- [ ] Implement `Ctrl+u` for delete to line start
- [ ] Implement `Ctrl+k` for delete to line end
- [ ] Implement `Ctrl+h` as alternative to backspace

**Validation:** Can edit text using Vim-style shortcuts

### 2.3 Update INSERT mode spec requirements
- [ ] Document all Vim movement and editing keys in vim-navigation spec
- [ ] Create scenarios for each keybinding
- [ ] Update mode-focus compatibility matrix

**Validation:** INSERT mode spec is complete with all Vim keys

---

## Phase 3: Scope Code Block Keys to Chat Buffer

### 3.1 Update chat-buffer spec with scope restrictions
- [ ] Add requirements clarifying code block keys only work in chat buffer focus
- [ ] Document "silently ignore" behavior for invalid focus areas
- [ ] Add scenarios for each code block operation showing both valid and invalid focus

**Validation:** chat-buffer spec clearly states scope restrictions

### 3.2 Update vim-navigation spec with scope clarifications
- [ ] Add focus-dependent routing scenarios for code block keys
- [ ] Document that `]c`, `[c`, `yc`, `yNc`, `ym` require chat buffer focus
- [ ] Update VISUAL mode entry to require chat buffer focus

**Validation:** vim-navigation spec reflects scope restrictions

### 3.3 Implement focus-dependent key routing
- [ ] Modify key routing logic to check focus before routing code block keys
- [ ] Implement silent ignore for invalid focus areas
- [ ] Add unit tests for routing with different focus states

**Validation:** Code block keys only work when focus is on chat buffer

---

## Phase 4: Update Task Lists

### 4.1 Remove command mode tasks from original add-tui-framework tasks
- [ ] Remove "Command Mode" section from tasks.md
- [ ] Remove "Enter command mode" task
- [ ] Remove "Execute command" task
- [ ] Remove "Supported commands" task
- [ ] Remove "Cancel command mode" task

**Validation:** Original tasks.md no longer references command mode

### 4.2 Add INSERT mode Vim movement tasks
- [ ] Add tasks for implementing Vim cursor movement in input area
- [ ] Add tasks for implementing Vim editing keys
- [ ] Add validation tasks for INSERT mode behavior

**Validation:** tasks.md includes INSERT mode Vim movement work

### 4.3 Add scope restriction tasks
- [ ] Add task for implementing focus-dependent key routing
- [ ] Add task for testing code block key scope
- [ ] Add task for testing VISUAL mode scope

**Validation:** tasks.md includes scope restriction work

---

## Phase 5: Update Help System

### 5.1 Update help overlay to reflect changes
- [ ] Remove command mode bindings from help
- [ ] Add INSERT mode Vim movement keys to help
- [ ] Add global shortcuts (`Ctrl+q`, `Ctrl+t`, `?`) to help
- [ ] Update code block key descriptions to mention scope restriction

**Validation:** Help shows current keybinding scheme without command mode

### 5.2 Update help for context-sensitive display
- [ ] Ensure help shows focus-specific keys
- [ ] When focus is on chat buffer, show code block keys
- [ ] When focus is on session list, don't show code block keys
- [ ] When in INSERT mode, show Vim movement keys

**Validation:** Help is context-aware based on focus and mode

---

## Phase 6: Validation and Testing

### 6.1 Test simplified key routing
- [ ] Verify code block keys don't work in session list
- [ ] Verify code block keys don't work in input area
- [ ] Verify code block keys work in chat buffer
- [ ] Verify VISUAL mode only works in chat buffer

**Validation:** All scope restrictions work correctly

### 6.2 Test INSERT mode Vim movement
- [ ] Test all movement keys in INSERT mode
- [ ] Test all editing keys in INSERT mode
- [ ] Verify `Esc` exits to NORMAL mode
- [ ] Verify traditional arrow keys still work as fallback

**Validation:** INSERT mode Vim movement works as expected

### 6.3 Test global shortcuts
- [ ] Test `Ctrl+q` for quit (with confirmation)
- [ ] Test `Ctrl+t` for new session
- [ ] Test `?` for help
- [ ] Verify these work regardless of focus

**Validation:** Global shortcuts work in NORMAL mode from any focus

### 6.4 Regression testing
- [ ] Ensure all NORMAL mode navigation still works
- [ ] Ensure all session list operations still work
- [ ] Ensure all chat buffer operations still work
- [ ] Ensure mode transitions still work

**Validation:** Existing functionality unaffected by changes

---

## Dependencies

- **Phase 1** must be completed before updating help (Phase 5)
- **Phase 2** can be done in parallel with Phase 3
- **Phase 3** must be completed before validation (Phase 6)
- **Phase 4** should be done alongside implementation phases

## Impact on Original Implementation Plan

### Removed from original plan:
- Command mode (~5-8 tasks removed)
- Command prompt component (~2 tasks removed)
- Command parser and executor (~3 tasks removed)
- Command help documentation (~1 task removed)

**Total reduction:** ~11-15 tasks

### Added to original plan:
- INSERT mode Vim movement keys (~5 tasks added)
- Focus-dependent routing tests (~3 tasks added)
- Updated help system (~2 tasks added)

**Total addition:** ~10 tasks

### Net effect:
- Slightly simpler implementation
- More predictable behavior (no command mode context)
- Better INSERT mode editing experience
- Clearer scope for code block operations

---

## Definition of Done

This refinement is complete when:
1. Command mode is fully removed from specs and tasks
2. INSERT mode supports Vim-style movement and editing
3. Code block keys only work when focus is on chat buffer
4. Global shortcuts replace common commands
5. All changes are validated with tests
6. Help system reflects the new keybinding scheme
