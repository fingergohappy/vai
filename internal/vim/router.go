// Package vim provides Vim-style mode management for the vai application.
package vim

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Router handles mode-aware key routing based on mode and focus.
type Router struct {
	keymap *Keymap
}

// NewRouter creates a new key router.
func NewRouter() *Router {
	return &Router{
		keymap: NewKeymap(),
	}
}

// Focus is a type alias for the UI focus to avoid circular imports.
// This is defined here to allow vim package to accept focus parameters
// without importing ui package.
type Focus int

const (
	FocusHistory Focus = iota
	FocusBuffer
	FocusInput
)

// Route routes a key message based on the current mode and focus.
// Returns the resulting message and a boolean indicating if the key was handled.
func (r *Router) Route(msg tea.KeyMsg, mode Mode, focus Focus) (tea.Msg, bool) {
	// Get the key string representation
	key := msg.String()

	// Look up the keybinding in the current mode
	action, ok := r.keymap.Lookup(mode, key)
	if !ok {
		return nil, false
	}

	// Execute the action
	return action(), true
}

// SetKeymap sets the keymap for the router.
func (r *Router) SetKeymap(keymap *Keymap) {
	r.keymap = keymap
}

// Keymap returns the current keymap.
func (r *Router) Keymap() *Keymap {
	return r.keymap
}

// CanTransition returns true if a mode transition is valid for the given focus.
func CanTransition(from Mode, to Mode, focus Focus) bool {
	// Mode-focus compatibility matrix
	// FocusHistory: NORMAL only
	// FocusBuffer: NORMAL, VISUAL
	// FocusInput: INSERT only

	switch focus {
	case FocusHistory:
		return to == ModeNormal

	case FocusBuffer:
		return to == ModeNormal || to == ModeVisual

	case FocusInput:
		return to == ModeInsert

	default:
		return false
	}
}
