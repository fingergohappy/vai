// Package app provides the top-level Bubble Tea Model for the vai application.
package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/fingergohappy/vai/internal/chat"
	"github.com/fingergohappy/vai/internal/config"
	"github.com/fingergohappy/vai/internal/input"
	"github.com/fingergohappy/vai/internal/session"
	ui "github.com/fingergohappy/vai/internal/ui"
	"github.com/fingergohappy/vai/internal/vim"
)

// Model is the top-level Bubble Tea Model for the vai application.
// It combines sub-models for different UI components and manages global state.
type Model struct {
	// Mode is the current Vim mode (NORMAL, INSERT, VISUAL)
	Mode vim.Mode

	// Focus is the currently focused UI area (History, Buffer, Input)
	Focus ui.Focus

	// Config holds application configuration
	Config config.Config

	// Sub-models for each UI component
	Session session.Model // Session list (left pane)
	Chat    chat.Model    // Chat buffer (right pane)
	Input   input.Model   // Input area (bottom)

	// UI components
	TitleBar *ui.TitleBar // Title bar (top)
	Layout   ui.Layout    // Computed layout for panes
	Styles   *ui.Styles   // Lipgloss styles

	// Ready flag indicates if the layout has been calculated
	ready bool

	// Quitting flag
	quitting bool
}

// NewModel creates a new top-level Model with default state.
func NewModel(cfg config.Config) Model {
	styles := ui.DefaultStyles()
	titleBar := ui.NewTitleBar(styles)

	return Model{
		Mode:     vim.ModeNormal,
		Focus:    ui.FocusBuffer, // Default to chat buffer
		Config:   cfg,
		Styles:   styles,
		TitleBar: titleBar,
		ready:    false,
		// Sub-models initialized with defaults
		Session: session.NewModel(),
		Chat:    chat.NewModel(),
		Input:   input.NewModel(),
	}
}

// Init initializes the top-level Model.
func (m Model) Init() tea.Cmd {
	// Initialize all sub-models and return their commands
	return tea.Batch(
		m.Session.Init(),
		m.Chat.Init(),
		m.Input.Init(),
	)
}

// Update handles messages and routes them to appropriate sub-models.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle global messages first
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle quit keys
		if msg.Type == tea.KeyCtrlC {
			m.quitting = true
			return m, tea.Quit
		}

		// Handle focus switching with Ctrl+w (only in NORMAL mode)
		if msg.Type == tea.KeyCtrlW && m.Mode == vim.ModeNormal {
			m.Focus = m.Focus.Next()
			return m, nil
		}

		// Handle mode transitions
		switch msg.Type {
		case tea.KeyRunes:
			// Enter INSERT mode with 'i' or 'a' in NORMAL mode
			if m.Mode == vim.ModeNormal {
				for _, r := range msg.Runes {
					if r == 'i' || r == 'a' {
						m.Mode = vim.ModeInsert
						m.Focus = ui.FocusInput
						m.Input.Focus()
						return m, nil
					}
				}
			}
		case tea.KeyEsc:
			// Return to NORMAL mode from INSERT mode
			if m.Mode == vim.ModeInsert {
				m.Mode = vim.ModeNormal
				m.Focus = ui.FocusBuffer
				m.Input.Blur()
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		// Calculate layout based on terminal size
		m.Layout = ui.CalculateLayout(msg)
		m.ready = true
		// Update sub-model sizes
		m.TitleBar.SetWidth(msg.Width)

		chatStyle := m.getPaneStyle(m.Focus == ui.FocusBuffer)
		chatFrameX, chatFrameY := chatStyle.GetFrameSize()
		chatInnerWidth := m.Layout.ChatBuffer.Width - chatFrameX
		chatInnerHeight := m.Layout.ChatBuffer.Height - chatFrameY
		if chatInnerWidth < 0 {
			chatInnerWidth = 0
		}
		if chatInnerHeight < 0 {
			chatInnerHeight = 0
		}
		m.Chat.SetSize(chatInnerWidth, chatInnerHeight)

		inputStyle := m.getPaneStyle(m.Focus == ui.FocusInput)
		inputFrameX, inputFrameY := inputStyle.GetFrameSize()
		inputInnerWidth := m.Layout.InputArea.Width - inputFrameX
		inputInnerHeight := m.Layout.InputArea.Height - inputFrameY
		if inputInnerWidth < 0 {
			inputInnerWidth = 0
		}
		if inputInnerHeight < 0 {
			inputInnerHeight = 0
		}
		m.Input.SetSize(inputInnerWidth, inputInnerHeight)
	}

	// Route messages to sub-models based on Mode and Focus
	var cmd tea.Cmd
	switch m.Mode {
	case vim.ModeInsert:
		// In INSERT mode, route keyboard messages to Input
		if m.Focus == ui.FocusInput {
			var model tea.Model
			model, cmd = m.Input.Update(msg)
			m.Input = model.(input.Model)
		}
	case vim.ModeNormal:
		// In NORMAL mode, sub-models handle non-mode-switching messages
		// TODO: Route to Session/Chat for navigation when implemented
	}

	return m, cmd
}

// View renders the entire UI.
func (m Model) View() string {
	if m.quitting {
		return "Thanks for using vai!\n"
	}

	if !m.ready {
		return "Initializing vai..."
	}

	// Render title bar
	titleBar := m.renderTitleBar()

	// Render session list pane with placeholder content
	sessionPane := m.renderSessionPane()

	// // Render chat buffer pane with placeholder content
	chatPane := m.renderChatPane()

	// Render input area pane with placeholder content
	inputPane := m.renderInputPane()

	// Join session list and chat buffer horizontally (top section)
	topSection := lipgloss.JoinHorizontal(
		lipgloss.Left,
		sessionPane,
		chatPane,
	)

	// Join title bar, top section, and input area vertically
	mainContent := lipgloss.JoinVertical(
		lipgloss.Top,
		titleBar,
		topSection,
		inputPane,
	)

	return mainContent
}

// renderSessionPane renders the session list pane with static placeholder content.
func (m Model) renderSessionPane() string {
	style := m.getPaneStyle(m.Focus == ui.FocusHistory)

	placeholder := "  [Sessions]\n\n" +
		"  • Session 1\n" +
		"  • Session 2\n" +
		"  • Session 3\n\n" +
		"  (TODO)"

	frameX, frameY := style.GetFrameSize()
	w := m.Layout.SessionList.Width - frameX
	h := m.Layout.SessionList.Height - frameY
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}

	return style.
		Width(w).
		Height(h).
		Render(placeholder)
}

// renderChatPane renders the chat buffer pane with static placeholder content.
func (m Model) renderChatPane() string {
	style := m.getPaneStyle(m.Focus == ui.FocusBuffer)

	frameX, frameY := style.GetFrameSize()
	w := m.Layout.ChatBuffer.Width - frameX
	h := m.Layout.ChatBuffer.Height - frameY
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}

	return style.
		Width(w).
		Height(h).
		Render(m.Chat.View())
}

// renderInputPane renders the input area pane with the Input sub-model.
func (m Model) renderInputPane() string {
	style := m.getPaneStyle(m.Focus == ui.FocusInput)

	frameX, frameY := style.GetFrameSize()
	w := m.Layout.InputArea.Width - frameX
	h := m.Layout.InputArea.Height - frameY
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}

	inputContent := m.Input.View()

	return style.
		Width(w).
		Height(h).
		Render(inputContent)
}

// renderTitleBar renders the title bar with the current session title.
func (m Model) renderTitleBar() string {
	currentTitle := m.Session.GetCurrentTitle()
	if currentTitle == "" {
		currentTitle = "New Chat"
	}
	return m.TitleBar.Render(currentTitle)
}

// getPaneStyle returns the appropriate border style based on focus and mode.
func (m Model) getPaneStyle(isFocused bool) lipgloss.Style {
	if isFocused {
		return m.Styles.FocusedBorder // Thick cyan border
	}

	// Non-focused: use mode color
	switch m.Mode {
	case vim.ModeNormal:
		return m.Styles.NormalModeBorder
	case vim.ModeInsert:
		return m.Styles.InsertModeBorder
	case vim.ModeVisual:
		return m.Styles.VisualModeBorder
	default:
		return m.Styles.NormalModeBorder
	}
}
