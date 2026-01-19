// Package ui provides shared UI components and utilities for the vai application.
package ui

// TitleBar is a component that displays the current session name.
type TitleBar struct {
	styles *Styles
	width  int
}

// NewTitleBar creates a new TitleBar component.
func NewTitleBar(styles *Styles) *TitleBar {
	return &TitleBar{
		styles: styles,
		width:  80, // default width
	}
}

// SetWidth sets the width of the title bar.
func (t *TitleBar) SetWidth(width int) {
	t.width = width
}

// Render renders the title bar with the given session title.
func (t *TitleBar) Render(sessionTitle string) string {
	title := "Sessions - " + sessionTitle
	return t.styles.TitleBar.
		Width(t.width).
		Height(1).
		Render(title)
}
