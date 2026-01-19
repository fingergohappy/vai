// Package ui provides shared UI components and utilities for the vai application.
package ui

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles contains all Lipgloss style definitions for the application.
type Styles struct {
	// Mode colors
	NormalMode lipgloss.Color
	InsertMode lipgloss.Color
	VisualMode lipgloss.Color

	// Mode border styles
	NormalModeBorder lipgloss.Style
	InsertModeBorder lipgloss.Style
	VisualModeBorder lipgloss.Style

	// Panes
	SessionList   lipgloss.Style
	ChatBuffer    lipgloss.Style
	InputArea     lipgloss.Style
	FocusedBorder lipgloss.Style

	// Text
	UserText      lipgloss.Style
	AssistantText lipgloss.Style
	CodeBlock     lipgloss.Style
	CodeBlockNum  lipgloss.Style

	// Messages
	InfoMessage  lipgloss.Style
	ErrorMessage lipgloss.Style

	// Title bar
	TitleBar lipgloss.Style
}

// DefaultStyles returns the default style definitions.
func DefaultStyles() *Styles {
	return &Styles{
		// Mode colors
		NormalMode: lipgloss.Color("252"), // White
		InsertMode: lipgloss.Color("142"), // Green
		VisualMode: lipgloss.Color("33"),  // Blue

		// Mode border styles
		NormalModeBorder: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")), // Gray

		InsertModeBorder: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("142")), // Green

		VisualModeBorder: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("33")), // Blue

		// Panes
		SessionList: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")),

		ChatBuffer: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")),

		InputArea: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")),

		FocusedBorder: lipgloss.NewStyle().
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("151")), // Cyan

		// Text
		UserText: lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")), // Gray

		AssistantText: lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")), // White

		CodeBlock: lipgloss.NewStyle().
			Background(lipgloss.Color("235")).
			Foreground(lipgloss.Color("230")).

			Padding(0, 1),

		CodeBlockNum: lipgloss.NewStyle().
			Foreground(lipgloss.Color("142")), // Green

		// Messages
		InfoMessage: lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")),

		ErrorMessage: lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")),

		// Title bar
		TitleBar: lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("252")). // White
			Background(lipgloss.Color("235")),  // Dark gray
	}
}
