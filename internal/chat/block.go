// Package chat provides the chat buffer component for displaying messages.
package chat

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BlockType represents the type of content block.
type BlockType int

const (
	// TypeText is a plain text block.
	TypeText BlockType = iota

	// TypeCode is a code block with syntax.
	TypeCode
)

// Block is the interface for all content block types.
type Block interface {
	Kind() BlockType
	Render(width int) string
}

// TextBlock represents plain text content.
type TextBlock struct {
	Text string // Plain text or markdown content
}

// Kind returns the block type.
func (b *TextBlock) Kind() BlockType {
	return TypeText
}

// Render renders the text block with word wrapping.
func (b *TextBlock) Render(width int) string {
	if width <= 0 {
		return ""
	}

	var out strings.Builder
	lines := strings.Split(b.Text, "\n")
	for i, line := range lines {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(wrapPlainText(line, width))
	}
	return out.String()
}

func wrapPlainText(s string, width int) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}

	var out strings.Builder
	lineW := 0

	writeWord := func(w string) {
		if lineW > 0 {
			out.WriteByte(' ')
			lineW++
		}
		out.WriteString(w)
		lineW += lipgloss.Width(w)
	}

	newline := func() {
		out.WriteByte('\n')
		lineW = 0
	}

	for _, w := range words {
		wW := lipgloss.Width(w)
		if wW > width {
			if lineW != 0 {
				newline()
			}
			r := []rune(w)
			start := 0
			for start < len(r) {
				end := start
				segW := 0
				for end < len(r) {
					cw := lipgloss.Width(string(r[end]))
					if segW+cw > width {
						break
					}
					segW += cw
					end += 1
				}
				if end == start {
					end = start + 1
				}
				out.WriteString(string(r[start:end]))
				start = end
				if start < len(r) {
					newline()
				}
			}
			continue
		}

		if lineW == 0 {
			writeWord(w)
			continue
		}

		if lineW+1+wW <= width {
			writeWord(w)
			continue
		}

		newline()
		writeWord(w)
	}

	return out.String()
}

// CodeBlock represents a code block with language information.
type CodeBlock struct {
	Lang   string   // Language identifier (go, python, bash, etc.)
	Lines  []string // Code content split by lines
	Number int      // Sequential code block number in message
}

// Kind returns the block type.
func (b *CodeBlock) Kind() BlockType {
	return TypeCode
}

// Render renders the code block with syntax highlighting placeholder.
func (b *CodeBlock) Render(width int) string {
	// TODO: Implement proper code block rendering
	// TODO: Add syntax highlighting
	var sb strings.Builder
	sb.WriteString("[")
	if b.Number < 10 {
		sb.WriteString(string(rune('0' + b.Number)))
	} else {
		sb.WriteString(string(rune('0' + b.Number/10)))
		sb.WriteString(string(rune('0' + b.Number%10)))
	}
	sb.WriteString("] ")
	if b.Lang != "" {
		sb.WriteString(b.Lang)
		sb.WriteString("\n")
	}
	for _, line := range b.Lines {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	return sb.String()
}

// Content returns the full content of the code block.
func (b *CodeBlock) Content() string {
	return strings.Join(b.Lines, "\n")
}

// NewTextBlock creates a new text block.
func NewTextBlock(text string) *TextBlock {
	return &TextBlock{Text: text}
}

// NewCodeBlock creates a new code block.
func NewCodeBlock(lang string, lines []string, number int) *CodeBlock {
	return &CodeBlock{
		Lang:   lang,
		Lines:  lines,
		Number: number,
	}
}
