// Package clipboard provides cross-platform clipboard operations.
package clipboard

import "fmt"

// Dummy is a fallback clipboard implementation for systems without clipboard support.
type Dummy struct{}

// NewDummy creates a new dummy clipboard instance.
func NewDummy() *Dummy {
	return &Dummy{}
}

// Copy returns an error indicating clipboard is not available.
func (d *Dummy) Copy(text string) error {
	return fmt.Errorf("clipboard not available. Please install pbcopy (macOS) or xclip/wl-copy (Linux)")
}

// Available returns false for the dummy clipboard.
func (d *Dummy) Available() bool {
	return false
}
