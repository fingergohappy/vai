// Package session provides session management and persistence.
package session

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the session manager Bubble Tea Model.
// It manages session persistence and the session list.
type Model struct {
	// Sessions holds all loaded sessions.
	Sessions []Session

	// CurrentID is the ID of the active session.
	CurrentID string

	// SelectedIndex is the index of the selected session in the list.
	SelectedIndex int

	// Width is the available width for rendering.
	Width int

	// Height is the available height for rendering.
	Height int

	// Ready indicates if the model is initialized.
	ready bool
}

// NewModel creates a new session manager model.
func NewModel() Model {
	return Model{
		Sessions:      []Session{},
		CurrentID:     "",
		SelectedIndex: 0,
	}
}

// Init initializes the session manager model.
func (m Model) Init() tea.Cmd {
	// TODO: Load sessions from storage
	return nil
}

// Update handles messages for the session manager.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.ready = true

	// TODO: Handle session selection
	// TODO: Handle session operations (create, delete, rename)
	// TODO: Handle search
	}

	return m, nil
}

// View renders the session list.
// Phase 1: Rendering is handled by the top-level app model.
func (m Model) View() string {
	// Phase 1: This View is not used; rendering is done by app.Model
	return ""
}

// AddSession adds a new session to the list.
func (m *Model) AddSession(session Session) {
	m.Sessions = append(m.Sessions, session)
}

// SetCurrent sets the current session by ID.
func (m *Model) SetCurrent(id string) {
	m.CurrentID = id
}

// Current returns the current session.
func (m *Model) Current() *Session {
	for i := range m.Sessions {
		if m.Sessions[i].ID == m.CurrentID {
			return &m.Sessions[i]
		}
	}
	return nil
}

// GetCurrentTitle returns the title of the current session.
// Returns "New Chat" if there is no current session.
func (m *Model) GetCurrentTitle() string {
	if len(m.Sessions) == 0 {
		return "New Chat"
	}
	if m.SelectedIndex >= 0 && m.SelectedIndex < len(m.Sessions) {
		return m.Sessions[m.SelectedIndex].Title
	}
	return "New Chat"
}
