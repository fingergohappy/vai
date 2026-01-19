// Package input provides the input area component.
package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textarea"
)

// Model is the input area Bubble Tea Model.
// It wraps bubbles.TextArea for multi-line text input.
type Model struct {
	// textarea is the underlying Bubbles textarea component.
	textarea textarea.Model

	// Placeholder is the placeholder text.
	Placeholder string

	// Focus indicates if the input area is focused.
	focused bool

	// Ready indicates if the model is initialized.
	ready bool
}

// NewModel creates a new input area model.
func NewModel() Model {
	ta := textarea.New()
	ta.Placeholder = "Type your message..."
	ta.ShowLineNumbers = false
	ta.Prompt = ""
	ta.CharLimit = 0 // No limit
	ta.SetWidth(80)
	ta.SetHeight(5)

	return Model{
		textarea:   ta,
		Placeholder: "Type your message...",
		focused:    false,
	}
}

// Init initializes the input area model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages for the input area.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.textarea.SetWidth(msg.Width)
		m.ready = true

	// TODO: Handle Vim-style movement in INSERT mode
	// TODO: Handle send message on Enter
	}

	// Update textarea
	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the input area.
func (m Model) View() string {
	if !m.ready {
		return "Loading input..."
	}
	return m.textarea.View()
}

// Value returns the current input text.
func (m Model) Value() string {
	return m.textarea.Value()
}

// SetValue sets the input text.
func (m *Model) SetValue(value string) {
	m.textarea.SetValue(value)
}

// Focus focuses the input area.
func (m *Model) Focus() {
	m.focused = true
	m.textarea.Focus()
}

// Blur removes focus from the input area.
func (m *Model) Blur() {
	m.focused = false
	m.textarea.Blur()
}

// Focused returns true if the input area is focused.
func (m Model) Focused() bool {
	return m.focused
}

// Reset clears the input text.
func (m *Model) Reset() {
	m.textarea.Reset()
}

// SetSize sets the size of the input area.
func (m *Model) SetSize(width, height int) {
	m.textarea.SetWidth(width)
	m.textarea.SetHeight(height)
	m.ready = true
}
