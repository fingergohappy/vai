// Package chat provides the chat buffer component for displaying messages.
package chat

import "time"

// Role represents the role of a message sender.
type Role string

const (
	// RoleUser is a message from the user.
	RoleUser Role = "user"

	// RoleAssistant is a message from the AI assistant.
	RoleAssistant Role = "assistant"
)

// Message represents a single message in the conversation.
type Message struct {
	ID        string    // Unique message identifier
	Role      Role      // "user" or "assistant"
	Blocks    []Block   // Ordered content blocks
	CreatedAt time.Time // Message timestamp
}

// NewMessage creates a new message with the given role and blocks.
func NewMessage(role Role, blocks []Block) Message {
	return Message{
		ID:        generateID(),
		Role:      role,
		Blocks:    blocks,
		CreatedAt: time.Now(),
	}
}

// generateID generates a unique ID for a message.
// TODO: Implement proper ID generation (UUID or similar).
func generateID() string {
	return "msg-" + time.Now().Format("20060102150405.000")
}
