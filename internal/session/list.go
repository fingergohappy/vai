// Package session provides session management and persistence.
package session

// List represents the session list component.
type List struct {
	// Sessions is the list of sessions to display.
	Sessions []Session

	// SelectedIndex is the index of the selected session.
	SelectedIndex int

	// Width is the available width for rendering.
	Width int

	// Height is the available height for rendering.
	Height int

	// Offset is the scroll offset for the list.
	Offset int
}

// NewList creates a new session list component.
func NewList() *List {
	return &List{
		Sessions:      []Session{},
		SelectedIndex: 0,
		Offset:        0,
	}
}

// SetSessions sets the sessions to display.
func (l *List) SetSessions(sessions []Session) {
	l.Sessions = sessions
	l.SelectedIndex = 0
	l.Offset = 0
}

// Select moves the selection to the given index.
func (l *List) Select(index int) {
	if index < 0 || index >= len(l.Sessions) {
		return
	}
	l.SelectedIndex = index

	// Adjust offset if needed
	if l.SelectedIndex < l.Offset {
		l.Offset = l.SelectedIndex
	}
	// TODO: Adjust offset for bottom of visible area
}

// SelectNext moves the selection to the next session.
func (l *List) SelectNext() {
	if l.SelectedIndex < len(l.Sessions)-1 {
		l.Select(l.SelectedIndex + 1)
	}
}

// SelectPrev moves the selection to the previous session.
func (l *List) SelectPrev() {
	if l.SelectedIndex > 0 {
		l.Select(l.SelectedIndex - 1)
	}
}

// Selected returns the currently selected session.
func (l *List) Selected() *Session {
	if l.SelectedIndex < 0 || l.SelectedIndex >= len(l.Sessions) {
		return nil
	}
	return &l.Sessions[l.SelectedIndex]
}

// VisibleRange returns the range of visible sessions based on offset.
func (l *List) VisibleRange() (start, end int) {
	return l.Offset, min(l.Offset+l.Height, len(l.Sessions))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
