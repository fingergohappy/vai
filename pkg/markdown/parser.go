// Package markdown provides markdown parsing for vai.
package markdown

// Parser parses markdown text into structured blocks.
type Parser struct {
	// TODO: Add parser configuration
}

// NewParser creates a new markdown parser.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses markdown text and returns a list of blocks.
// TODO: Implement actual markdown parsing.
func (p *Parser) Parse(text string) []Block {
	// Placeholder implementation
	// TODO: Use a proper markdown library or implement parsing
	return []Block{
		&TextBlock{Content: text},
	}
}

// ParseCodeBlocks extracts code blocks from markdown text.
// TODO: Implement code block extraction.
func (p *Parser) ParseCodeBlocks(text string) []*CodeBlock {
	return []*CodeBlock{}
}
