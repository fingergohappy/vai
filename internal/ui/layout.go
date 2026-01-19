// Package ui provides shared UI components and utilities for the vai application.
package ui

import tea "github.com/charmbracelet/bubbletea"

// paneInnerPadding is the vertical padding inside panes (for borders, spacing, etc.)
const paneInnerPadding = 5

// Layout represents the dimensions and positions of UI panes.
type Layout struct {
	// Width is the total terminal width.
	Width int

	// Height is the total terminal height.
	Height int

	// TitleBar is the layout for the title bar (top).
	TitleBar PaneLayout

	// SessionList is the layout for the session list pane (left).
	SessionList PaneLayout

	// ChatBuffer is the layout for the chat buffer pane (right).
	ChatBuffer PaneLayout

	// InputArea is the layout for the input area (bottom).
	InputArea PaneLayout
}

// PaneLayout represents the position and size of a single pane.
type PaneLayout struct {
	X      int // X position (0-indexed from left)
	Y      int // Y position (0-indexed from top)
	Width  int // Pane width
	Height int // Pane height
}

// CalculateLayout computes the layout based on terminal size.
func CalculateLayout(msg tea.WindowSizeMsg) Layout {
	width := msg.Width
	height := msg.Height

	titleBarHeight := 1
	inputHeight := 5

	// Session list: 20% width (minimum 20 chars)
	sessionWidth := width * 20 / 100
	if sessionWidth < 20 {
		sessionWidth = 20
	}
	if sessionWidth > width-40 {
		sessionWidth = width / 3
	}

	// Chat buffer: remaining width
	chatWidth := width - sessionWidth

	// Content height (total - title bar - input - pane padding)
	contentHeight := height - titleBarHeight - inputHeight - paneInnerPadding

	return Layout{
		Width:  width,
		Height: height,
		TitleBar: PaneLayout{
			X:      0,
			Y:      0,
			Width:  width,
			Height: titleBarHeight,
		},
		SessionList: PaneLayout{
			X:      0,
			Y:      titleBarHeight,
			Width:  sessionWidth,
			Height: contentHeight,
		},
		ChatBuffer: PaneLayout{
			X:      sessionWidth,
			Y:      titleBarHeight,
			Width:  chatWidth,
			Height: contentHeight,
		},
		InputArea: PaneLayout{
			X:      0,
			Y:      titleBarHeight + contentHeight,
			Width:  width,
			Height: inputHeight,
		},
	}
}
