// Package vim provides Vim-style mode management for the vai application.
package vim

// Mode represents the current Vim mode.
type Mode int

const (
	// ModeNormal is the default mode for navigation and commands.
	ModeNormal Mode = iota

	// ModeInsert is for text input in the input area.
	ModeInsert

	// ModeVisual is for text selection in the chat buffer.
	ModeVisual
)

// String returns the string representation of the mode.
func (m Mode) String() string {
	switch m {
	case ModeNormal:
		return "NORMAL"
	case ModeInsert:
		return "INSERT"
	case ModeVisual:
		return "VISUAL"
	default:
		return "UNKNOWN"
	}
}

// IsValid returns true if the mode is valid.
func (m Mode) IsValid() bool {
	return m >= ModeNormal && m <= ModeVisual
}
