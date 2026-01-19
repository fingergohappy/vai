// Package clipboard provides cross-platform clipboard operations.
package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
)

// MacOS is the clipboard implementation for macOS.
type MacOS struct{}

// NewMacOS creates a new macOS clipboard instance.
func NewMacOS() *MacOS {
	return &MacOS{}
}

// Copy copies text to the macOS clipboard using pbcopy.
func (m *MacOS) Copy(text string) error {
	if !m.Available() {
		return fmt.Errorf("pbcopy not available")
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pbcopy failed: %w", err)
	}

	return nil
}

// Available returns true if pbcopy is available.
func (m *MacOS) Available() bool {
	_, err := exec.LookPath("pbcopy")
	return err == nil
}
