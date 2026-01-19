# Keybindings Reference

Complete reference of all keyboard shortcuts in vai.

## Mode Overview

| Mode | Purpose | Status Bar Color |
|------|---------|-----------------|
| NORMAL | Navigation, viewing, commands | White |
| INSERT | Text input in input area | Green |
| VISUAL | Text selection in chat buffer | Blue |

---

## NORMAL Mode

### Navigation

| Key | Action |
|-----|--------|
| `j` / `Ctrl+e` | Scroll down one line |
| `k` / `Ctrl+y` | Scroll up one line |
| `Ctrl+f` | Scroll down one page |
| `Ctrl+b` | Scroll up one page |
| `Ctrl+d` | Scroll down half screen |
| `Ctrl+u` | Scroll up half screen |
| `G` | Go to end of conversation |
| `gg` | Go to start of conversation |

### Word and Line Movement (chat buffer)

| Key | Action |
|-----|--------|
| `w` | Move to start of next word |
| `b` | Move to start of previous word |
| `e` | Move to end of current word |
| `0` | Move to start of current line |
| `$` | Move to end of current line |

### Code Block Operations (chat buffer only)

| Key | Action |
|-----|--------|
| `]c` | Jump to next code block |
| `[c` | Jump to previous code block |
| `N]c` | Jump N code blocks forward |
| `yc` | Copy current code block |
| `yNc` | Copy Nth code block |
| `ym` | Copy entire message |

### Pane Switching

| Key | Action |
|-----|--------|
| `Ctrl+w h` | Focus session list (left) |
| `Ctrl+w l` | Focus chat buffer (right) |
| `Ctrl+w j` | Focus input area (bottom) |
| `Ctrl+w k` | Focus upward in pane order |

### Session List (when focused)

| Key | Action |
|-----|--------|
| `j` / `k` | Move to next/previous session |
| `G` / `gg` | Go to last/first session |
| `Enter` | Open selected session |
| `r` | Rename selected session |
| `dd` | Delete selected session |
| `/` | Search sessions |
| `n` / `N` | Next/previous search result |

### Global Shortcuts

| Key | Action |
|-----|--------|
| `i` / `a` | Enter INSERT mode (move to input) |
| `v` | Enter VISUAL mode (chat buffer only) |
| `Ctrl+t` | Create new session |
| `Ctrl+q` | Quit (press twice to confirm) |
| `?` | Show help overlay |
| `Esc` / `Ctrl+c` | Return to NORMAL (from any mode) |

---

## INSERT Mode

INSERT mode is only active when focus is on the input area.

### Character Input

| Key | Action |
|-----|--------|
| Printable chars | Insert text |
| `Enter` | Send message |
| `Backspace` | Delete character before cursor |
| `Delete` / `Ctrl+d` | Delete character at cursor |
| `Arrow keys` | Move cursor (fallback) |

### Vim Movement

| Key | Action |
|-----|--------|
| `h` / `Ctrl+b` | Move cursor left |
| `l` / `Ctrl+f` | Move cursor right |
| `w` | Move to next word |
| `b` | Move to previous word |
| `e` | Move to end of current word |
| `0` / `Ctrl+a` | Move to line start |
| `$` / `Ctrl+e` | Move to line end |

### Editing

| Key | Action |
|-----|--------|
| `Ctrl+w` | Delete word before cursor |
| `Ctrl+u` | Delete to line start |
| `Ctrl+k` | Delete to line end |
| `Ctrl+h` | Delete character before cursor |

### Exit INSERT Mode

| Key | Action |
|-----|--------|
| `Esc` / `Ctrl+[` | Return to NORMAL mode |

---

## VISUAL Mode

VISUAL mode is only available in the chat buffer.

| Key | Action |
|-----|--------|
| `v` | Start selection (from NORMAL mode) |
| `h` / `j` / `k` / `l` | Expand selection |
| `w` / `b` / `0` / `$` | Expand selection by word/line |
| `y` | Copy selection to clipboard |
| `Esc` / `Ctrl+c` | Cancel selection |

---

## Mode-Focus Compatibility

Some modes are only valid in certain focus areas:

| Focus Area | NORMAL | INSERT | VISUAL |
|------------|--------|--------|--------|
| Session List | ✅ | ❌ | ❌ |
| Chat Buffer | ✅ | ❌ | ✅ |
| Input Area | ✅ | ✅ | ❌ |

**Note:** Attempting to enter an invalid mode will be silently ignored.

---

## Scope Restrictions

### Code Block Keys Only Work in Chat Buffer

These keys are **ignored** when focus is NOT on chat buffer:
- `]c`, `[c` - Jump to code blocks
- `yc`, `yNc`, `ym` - Copy operations

### VISUAL Mode Only Works in Chat Buffer

Pressing `v` when focus is on session list or input area will be ignored.

---

## Customization

Keybindings can be customized in `~/.config/vai/config.yaml`:

```yaml
keybindings:
  overrides:
    normal:
      next_line: "j"
      prev_line: "k"
      next_code_block: "]c"
      prev_code_block: "[c"
    visual:
      copy_selection: "y"
    insert:
      move_left: "h"
      move_right: "l"
```

---

## Help

Press `?` anytime in NORMAL mode to see context-sensitive help based on your current focus.
