// Package chat provides the chat buffer component for displaying messages.
package chat

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ChatMessage handles rendering of individual chat messages with role-based styling.
type ChatMessage struct {
	// Styles for user messages
	userBorder    lipgloss.Style
	userLabel     lipgloss.Style
	userContent   lipgloss.Style
	userContainer lipgloss.Style

	// Styles for AI messages
	aiBorder    lipgloss.Style
	aiLabel     lipgloss.Style
	aiContent   lipgloss.Style
	aiContainer lipgloss.Style
}

// NewChatMessage creates a new ChatMessage renderer with predefined styles.
func NewChatMessage() *ChatMessage {
	userBorderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("142"))

	userLabelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("142")).
		Bold(true)

	userContentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("244"))

	userContainerStyle := lipgloss.NewStyle().
		Width(0). // Will be set dynamically
		Align(lipgloss.Right).
		MarginTop(1).   // Add top margin
		MarginBottom(1) // Add bottom margin

	aiBorderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("33"))

	aiLabelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("33")).
		Bold(true)

	aiContentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))

	aiContainerStyle := lipgloss.NewStyle().
		Width(0). // Will be set dynamically
		Align(lipgloss.Left).
		MarginTop(1).   // Add top margin
		MarginBottom(1) // Add bottom margin

	return &ChatMessage{
		userBorder:    userBorderStyle,
		userLabel:     userLabelStyle,
		userContent:   userContentStyle,
		userContainer: userContainerStyle,
		aiBorder:      aiBorderStyle,
		aiLabel:       aiLabelStyle,
		aiContent:     aiContentStyle,
		aiContainer:   aiContainerStyle,
	}
}

// Render renders a message with appropriate styling based on its role.
func (cm *ChatMessage) Render(msg Message, maxWidth int) string {
	switch msg.Role {
	case RoleUser:
		return cm.renderUserMessage(msg, maxWidth)
	case RoleAssistant:
		return cm.renderAssistantMessage(msg, maxWidth)
	default:
		return cm.renderAssistantMessage(msg, maxWidth)
	}
}

// renderUserMessage renders a user message with green border, right-aligned.
func (cm *ChatMessage) renderUserMessage(msg Message, maxWidth int) string {
	boxed := boxedMessage("You", msg, maxWidth, lipgloss.Color("142"))
	return cm.userContainer.Width(maxWidth).Render(boxed)
}

func bubbleMaxWidth(paneWidth int) int {
	w := paneWidth * 2 / 3
	if w < 30 {
		w = 30
	}
	return w
}

func boxedMessage(title string, msg Message, maxPaneWidth int, borderColor lipgloss.Color) string {
	maxBubbleWidth := bubbleMaxWidth(maxPaneWidth)
	if maxBubbleWidth < 10 {
		maxBubbleWidth = 10
	}

	padX := 1
	frameX := 2

	innerMaxWidth := maxBubbleWidth - frameX - padX*2
	if innerMaxWidth < 1 {
		innerMaxWidth = 1
	}

	var blocks []string
	for _, block := range msg.Blocks {
		blocks = append(blocks, block.Render(innerMaxWidth))
	}

	content := strings.Join(blocks, "\n")

	contentWidth, _ := lipgloss.Size(content)
	if contentWidth < 1 {
		contentWidth = 1
	}

	bubbleWidth := contentWidth + frameX + padX*2
	if bubbleWidth > maxBubbleWidth {
		bubbleWidth = maxBubbleWidth
	}
	if bubbleWidth < 10 {
		bubbleWidth = 10
	}

	innerWidth := bubbleWidth - frameX
	if innerWidth < 1 {
		innerWidth = 1
	}

	inner := lipgloss.NewStyle().Width(innerWidth).Padding(0, padX).Render(content)

	b := lipgloss.NormalBorder()
	availableTop := bubbleWidth - 2
	if availableTop < 0 {
		availableTop = 0
	}

	t := title
	if lipgloss.Width(t) > availableTop {
		if availableTop == 0 {
			t = ""
		} else {
			t = t[:1]
		}
	}

	tW := lipgloss.Width(t)
	spacedTitle := " " + t + " "
	stW := lipgloss.Width(spacedTitle)
	if stW > availableTop {
		spacedTitle = t
		stW = tW
	}

	leftFill := (availableTop - stW) / 2
	rightFill := availableTop - stW - leftFill
	b.Top = strings.Repeat(b.Top, leftFill) + spacedTitle + strings.Repeat(b.Top, rightFill)

	border := lipgloss.NewStyle().Border(b).BorderForeground(borderColor).Width(bubbleWidth)
	return border.Render(inner)
}

// renderAssistantMessage renders an AI message with blue border, left-aligned.
func (cm *ChatMessage) renderAssistantMessage(msg Message, maxWidth int) string {
	boxed := boxedMessage("AI", msg, maxWidth, lipgloss.Color("33"))
	return cm.aiContainer.Render(boxed)
}
