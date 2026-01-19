# vai - Vim-style AI Chat TUI

A terminal-based AI chat tool designed for engineers who work primarily in the terminal. Built with a Vim-style interface for efficient keyboard-driven interaction.

## Features

- **Keyboard-first interface** - All operations accessible without a mouse
- **Vim-style navigation** - NORMAL, INSERT, and VISUAL modes
- **Code block focus** - Easy navigation and copying of code blocks
- **Session management** - Persistent chat history
- **Cross-platform** - Works on macOS and Linux

## Installation

### From source

```bash
git clone https://github.com/fingergohappy/vai.git
cd vai
make install
```

### Using go install

```bash
go install github.com/fingergohappy/vai@latest
```

## Quick Start

```bash
# Start vai
vai

# Common keybindings (NORMAL mode)
i           - Enter INSERT mode (type message)
Esc         - Return to NORMAL mode
j/k         - Scroll down/up
Ctrl+w h/l  - Switch between panes
?           - Show help
Ctrl+q      - Quit
```

## Keybindings

### NORMAL Mode

| Key | Action |
|-----|--------|
| `i` / `a` | Enter INSERT mode |
| `v` | Enter VISUAL mode (chat buffer) |
| `j` / `k` | Scroll down / up |
| `Ctrl+f` / `Ctrl+b` | Scroll down / up one page |
| `G` / `gg` | Go to end / start of conversation |
| `]c` / `[c` | Jump to next / previous code block |
| `yc` | Copy current code block |
| `ym` | Copy entire message |
| `Ctrl+w h/l/j/k` | Switch focus (history/buffer/input) |
| `Ctrl+t` | Create new session |
| `Ctrl+q` | Quit application |
| `?` | Show help |

### INSERT Mode

| Key | Action |
|-----|--------|
| `Enter` | Send message |
| `Esc` / `Ctrl+[` | Exit to NORMAL mode |
| `h` / `l` | Move cursor left / right |
| `w` / `b` | Move to next / previous word |
| `0` / `$` | Move to line start / end |
| `Ctrl+w` | Delete word before cursor |
| `Ctrl+u` | Delete to line start |

### VISUAL Mode

| Key | Action |
|-----|--------|
| `v` | Start selection |
| Movement keys | Expand selection |
| `y` | Copy selection |
| `Esc` | Cancel selection |

## Project Structure

```
vai/
├── cmd/vai/        # Application entry point
├── internal/       # Private packages
│   ├── app/        # Top-level Bubble Tea Model
│   ├── vim/        # Mode system
│   ├── ui/         # UI components
│   ├── chat/       # Chat buffer
│   ├── session/    # Session management
│   ├── input/      # Input area
│   ├── clipboard/  # Clipboard operations
│   └── config/     # Configuration
├── pkg/            # Public packages
└── docs/           # Documentation
```

## Configuration

Configuration is stored in `~/.config/vai/config.yaml`:

```yaml
editor:
  tab_width: 4
  word_wrap: true
  line_numbers: true

keybindings:
  # Custom keybindings override defaults
  overrides: {}

theme:
  name: default
  colors: {}
```

## Development

### Build

```bash
make build
```

### Run

```bash
make run
```

### Test

```bash
make test
```

### Lint

```bash
make lint
```

## License

MIT License - see LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
