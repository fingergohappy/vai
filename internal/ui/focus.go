// Package ui provides shared UI components and utilities for the vai application.
package ui

// Focus represents the currently focused UI area.
type Focus int

const (
	// FocusHistory is the session list pane (left).
	FocusHistory Focus = iota

	// FocusBuffer is the chat buffer pane (right).
	FocusBuffer

	// FocusInput is the input area (bottom).
	FocusInput
)

// String returns the string representation of the focus.
func (f Focus) String() string {
	switch f {
	case FocusHistory:
		return "History"
	case FocusBuffer:
		return "Buffer"
	case FocusInput:
		return "Input"
	default:
		return "Unknown"
	}
}

// IsValid returns true if the focus value is valid.
func (f Focus) IsValid() bool {
	return f >= FocusHistory && f <= FocusInput
}

// Next returns the next focus area in cyclic order.
func (f Focus) Next() Focus {
	switch f {
	case FocusHistory:
		return FocusBuffer
	case FocusBuffer:
		return FocusInput
	case FocusInput:
		return FocusHistory
	default:
		return FocusBuffer
	}
}

// Prev returns the previous focus area in cyclic order.
func (f Focus) Prev() Focus {
	switch f {
	case FocusHistory:
		return FocusInput
	case FocusBuffer:
		return FocusHistory
	case FocusInput:
		return FocusBuffer
	default:
		return FocusBuffer
	}
}
