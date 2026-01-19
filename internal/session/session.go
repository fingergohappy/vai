// Package session provides session management and persistence.
package session

import "time"

// Session represents a single chat session.
type Session struct {
	ID        string    // Unique session identifier
	Title     string    // Session title
	Messages  []Message // Ordered list of messages
	CreatedAt time.Time // Session creation timestamp
	UpdatedAt time.Time // Last update timestamp
	Model     string    // AI model used (e.g., "gpt-4", "claude-3")
}

// Message represents a message in a session.
type Message struct {
	ID        string      // Unique message identifier
	Role      string      // "user" or "assistant"
	Content   []Block     // Structured content blocks
	CreatedAt time.Time   // Message timestamp
}

// Block represents a content block in a message.
type Block interface {
	Kind() BlockType
	Render(width int) string
}

// BlockType represents the type of content block.
type BlockType int

const (
	// TextBlock is a plain text block.
	TextBlock BlockType = iota

	// CodeBlock is a code block with syntax.
	CodeBlock
)

// NewSession creates a new session with a generated title.
func NewSession() Session {
	now := time.Now()
	return Session{
		ID:        generateID(),
		Title:     "New Chat",
		Messages:  []Message{},
		CreatedAt: now,
		UpdatedAt: now,
		Model:     "gpt-4",
	}
}

// generateID generates a unique ID for a session.
// TODO: Implement proper ID generation (UUID or similar).
func generateID() string {
	return "session-" + time.Now().Format("20060102150405")
}

// AddMessage adds a message to the session.
func (s *Session) AddMessage(msg Message) {
	s.Messages = append(s.Messages, msg)
	s.UpdatedAt = time.Now()
}

// UpdateTitle updates the session title.
func (s *Session) UpdateTitle(title string) {
	s.Title = title
	s.UpdatedAt = time.Now()
}
