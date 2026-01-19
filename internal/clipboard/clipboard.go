// Package clipboard provides cross-platform clipboard operations.
package clipboard

// Clipboard is the interface for clipboard operations.
type Clipboard interface {
	// Copy copies text to the clipboard.
	Copy(text string) error

	// Available returns true if the clipboard is available.
	Available() bool
}

// New returns the appropriate clipboard implementation for the current platform.
// Platform detection will be done in build files.
func New() Clipboard {
	// Try to detect platform and return appropriate implementation
	// For now, return dummy
	return NewDummy()
}
