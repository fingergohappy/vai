// Package chat provides the chat buffer component for displaying messages.
package chat

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Model is the chat buffer Bubble Tea Model.
// It displays and enables interaction with AI conversation history.
type Model struct {
	// Messages holds the conversation history.
	Messages []Message

	// ViewportOffset is the current scroll position.
	ViewportOffset int

	// CursorLine is the current cursor line position.
	CursorLine int

	// Selection holds VISUAL mode selection state.
	Selection Selection

	// Width is the available width for rendering.
	Width int

	// Height is the available height for rendering.
	Height int

	// messageRenderer handles rendering of individual messages.
	messageRenderer *ChatMessage

	// Ready indicates if the model is initialized.
	ready bool
}

// Selection represents text selection in VISUAL mode.
type Selection struct {
	Active bool
	Start  int // Start line
	End    int // End line
}

// NewModel creates a new chat buffer model with mock data for visual testing.
func NewModel() Model {
	return Model{
		Messages: []Message{
			NewMessage(RoleAssistant, []Block{
				NewTextBlock("Hello! How can I help you today?"),
			}),
			NewMessage(RoleUser, []Block{
				NewTextBlock("What is the weather like?"),
			}),
			NewMessage(RoleAssistant, []Block{
				NewTextBlock("I don't have access to real-time weather data, but I can help you with many other tasks!"),
			}),
		},
		ViewportOffset:  0,
		CursorLine:      0,
		Selection:       Selection{Active: false},
		messageRenderer: NewChatMessage(),
	}
}

// Init initializes the chat buffer model.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages for the chat buffer.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.ready = true

		// TODO: Handle scroll messages
		// TODO: Handle cursor movement
		// TODO: Handle code block navigation
		// TODO: Handle VISUAL mode selection
	}

	return m, nil
}

// View renders the chat buffer with styled messages.
func (m Model) View() string {
	// Handle empty state
	if len(m.Messages) == 0 {
		return "  [Chat Buffer]\n\n" +
			"  Welcome to vai!\n" +
			"  Start a conversation..."
	}

	// Render each message
	var renderedMessages []string
	for _, msg := range m.Messages {
		renderedMessages = append(renderedMessages, m.messageRenderer.Render(msg, m.Width))
	}

	// Join messages with single newline (messages have MarginTop/Bottom)
	return strings.Join(renderedMessages, "\n")
}

// SetWidth sets the available width for rendering.
func (m *Model) SetWidth(width int) {
	m.Width = width
}

func (m *Model) SetSize(width, height int) {

	m.Width = width
	m.Height = height
	m.ready = true
}

// AddMessage adds a new message to the chat buffer.
func (m *Model) AddMessage(msg Message) {
	m.Messages = append(m.Messages, msg)
}

// ScrollDown scrolls the buffer down by one line.
func (m *Model) ScrollDown() {
	if m.ViewportOffset < len(m.Messages)-1 {
		m.ViewportOffset++
	}
}

// ScrollUp scrolls the buffer up by one line.
func (m *Model) ScrollUp() {
	if m.ViewportOffset > 0 {
		m.ViewportOffset--
	}
}
