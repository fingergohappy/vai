# Spec: Chat Buffer (Refined)

**Status:** Active
**Version:** 1.1.0
**Owners:** vai project

## Overview

This refines the chat buffer spec to explicitly clarify that code block operations are only available when focus is on the chat buffer.

## MODIFIED Requirements

### Requirement: Code Block Navigation Scope

The chat buffer's code block navigation MUST only work when focus is on the chat buffer.

#### Scenario: Code block navigation requires chat buffer focus

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses:
- `]c` - Jump to next code block in current message
- `[c` - Jump to previous code block in current message
**Then** the viewport SHOULD scroll to show the code block
**And** the code block SHOULD be highlighted briefly

**Given** the user is in NORMAL mode
**And** focus is on the session list or input area
**When** the user presses `]c` or `[c`
**Then** these keys SHOULD be silently ignored
**And** NO error message SHOULD be displayed
**And** NO navigation SHOULD occur

#### Scenario: Counted code block jumps

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses `N]c` (where N is a count)
**Then** the viewport SHOULD jump N code blocks forward
**And** if N exceeds the number of code blocks, it SHOULD wrap or clamp

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer
**When** the user presses `N]c`
**Then** the key combination SHOULD be ignored
**And** NO action SHOULD occur

---

### Requirement: Code Block Copying Scope

The chat buffer's code block copying MUST only work when focus is on the chat buffer.

#### Scenario: Copy current code block requires chat buffer focus

**Given** the user has navigated to a code block
**And** focus is on the chat buffer
**When** the user presses `yc` in NORMAL mode
**Then** the buffer SHOULD:
- Copy the entire code block content to system clipboard
- Display confirmation: "Copied code block [N] to clipboard"
- Use the appropriate clipboard command

**Given** the user is in NORMAL mode
**And** focus is on the session list or input area
**When** the user presses `yc`
**Then** the key SHOULD be silently ignored
**And** NO copy operation SHOULD occur
**And** NO message SHOULD be displayed

#### Scenario: Copy specific code block by number requires chat buffer focus

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses `yNc` (where N is a number)
**Then** the buffer SHOULD:
- Copy the Nth code block (1-indexed) from the current message
- Display confirmation with block number
- Show error if N exceeds the number of code blocks

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer
**When** the user presses `yNc`
**Then** the key combination SHOULD be ignored
**And** NO action SHOULD occur

#### Scenario: Copy entire message requires chat buffer focus

**Given** the user is viewing a message
**And** focus is on the chat buffer
**When** the user presses `ym` in NORMAL mode
**Then** the buffer SHOULD:
- Copy the entire message content (all blocks) to clipboard
- Preserve code block formatting
- Display confirmation

**Given** the user is in NORMAL mode
**And** focus is NOT on the chat buffer
**When** the user presses `ym`
**Then** the key SHOULD be ignored
**And** NO action SHOULD occur

---

### Requirement: VISUAL Mode Scope

VISUAL mode MUST only be available when focus is on the chat buffer.

#### Scenario: Enter VISUAL mode only in chat buffer

**Given** the user is in NORMAL mode
**And** focus is on the chat buffer
**When** the user presses `v`
**Then** the buffer SHOULD:
- Enter VISUAL mode
- Start selection at the current cursor position
- Highlight the selected character
- Update status bar to show VISUAL mode

**Given** the user is in NORMAL mode
**And** focus is on the session list
**When** the user presses `v`
**Then** the key SHOULD be silently ignored
**And** the mode SHOULD remain NORMAL
**And** NO selection SHOULD be created

**Given** the user is in NORMAL mode
**And** focus is on the input area
**When** the user presses `v`
**Then** the key SHOULD be silently ignored
**And** the mode SHOULD remain NORMAL

---

## ADDED Requirements

### Requirement: Focus-Dependent Key Routing

The application MUST route code block and VISUAL mode keys only when focus is on the chat buffer.

#### Scenario: Key routing for code block operations

**Given** a key press occurs (`]c`, `[c`, `yc`, `yNc`, `ym`)
**When** the key event is processed
**Then** the routing SHOULD:

1. Check if focus is on chat buffer
2. If YES: Process the key as code block operation
3. If NO: Silently ignore the key

#### Scenario: Key routing for VISUAL mode entry

**Given** a `v` key press occurs
**When** the key event is processed
**Then** the routing SHOULD:

1. Check if focus is on chat buffer
2. If YES: Enter VISUAL mode and start selection
3. If NO: Silently ignore the key

#### Scenario: Visual feedback for ignored keys

**Given** the user presses a key that is invalid for the current focus
**When** the key is ignored
**Then** the application SHOULD:
- NOT display any error message (silent ignore)
- NOT change modes
- NOT perform any action
- Maintain current state unchanged

---

## Summary of Scope Clarifications

This change clarifies the following:

| Operation | Valid Focus Areas | Invalid Focus Areas |
|-----------|------------------|-------------------|
| `]c` / `[c` (jump to code block) | Chat Buffer only | Session List, Input Area |
| `yc` / `yNc` (copy code block) | Chat Buffer only | Session List, Input Area |
| `ym` (copy message) | Chat Buffer only | Session List, Input Area |
| `v` (enter VISUAL mode) | Chat Buffer only | Session List, Input Area |

These keys will be **silently ignored** when focus is not on the chat buffer, preventing accidental operations and keeping behavior predictable.

---

## Cross-References

- **vim-navigation**: Updated to remove command mode and clarify focus-dependent key routing
- **tui-framework**: Defines the focus management that this spec depends on
