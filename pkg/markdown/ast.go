// Package markdown provides markdown parsing for vai.
package markdown

// BlockType represents the type of markdown block.
type BlockType int

const (
	// BlockTypeText is a plain text block.
	BlockTypeText BlockType = iota

	// BlockTypeCode is a code block.
	BlockTypeCode

	// BlockTypeHeading is a heading.
	BlockTypeHeading
)

// Block is the interface for all markdown block types.
type Block interface {
	Type() BlockType
}

// TextBlock represents plain text content.
type TextBlock struct {
	Content string
}

// Type returns the block type.
func (b *TextBlock) Type() BlockType {
	return BlockTypeText
}

// CodeBlock represents a code block.
type CodeBlock struct {
	Lang    string
	Content string
}

// Type returns the block type.
func (b *CodeBlock) Type() BlockType {
	return BlockTypeCode
}

// Heading represents a heading.
type Heading struct {
	Level   int
	Content string
}

// Type returns the block type.
func (h *Heading) Type() BlockType {
	return BlockTypeHeading
}

// AST represents the abstract syntax tree of a markdown document.
type AST struct {
	Blocks []Block
}

// NewAST creates a new AST.
func NewAST() *AST {
	return &AST{
		Blocks: []Block{},
	}
}

// AddBlock adds a block to the AST.
func (a *AST) AddBlock(block Block) {
	a.Blocks = append(a.Blocks, block)
}
