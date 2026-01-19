// Package vim provides Vim-style mode management for the vai application.
package vim

import "github.com/charmbracelet/bubbletea"

// Keymap defines keybindings for each mode.
type Keymap struct {
	// NORMAL mode keybindings
	Normal map[string]func() tea.Msg

	// INSERT mode keybindings
	Insert map[string]func() tea.Msg

	// VISUAL mode keybindings
	Visual map[string]func() tea.Msg
}

// NewKeymap creates a new keymap with default bindings.
func NewKeymap() *Keymap {
	return &Keymap{
		Normal: make(map[string]func() tea.Msg),
		Insert: make(map[string]func() tea.Msg),
		Visual: make(map[string]func() tea.Msg),
	}
}

// Bind adds a keybinding in the specified mode.
func (k *Keymap) Bind(mode Mode, key string, action func() tea.Msg) {
	switch mode {
	case ModeNormal:
		k.Normal[key] = action
	case ModeInsert:
		k.Insert[key] = action
	case ModeVisual:
		k.Visual[key] = action
	}
}

// Lookup looks up a keybinding in the specified mode.
// Returns the action function and true if found, nil and false otherwise.
func (k *Keymap) Lookup(mode Mode, key string) (func() tea.Msg, bool) {
	var bindings map[string]func() tea.Msg
	switch mode {
	case ModeNormal:
		bindings = k.Normal
	case ModeInsert:
		bindings = k.Insert
	case ModeVisual:
		bindings = k.Visual
	default:
		return nil, false
	}

	action, ok := bindings[key]
	return action, ok
}
