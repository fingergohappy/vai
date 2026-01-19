// Package clipboard provides cross-platform clipboard operations.
package clipboard

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Linux is the clipboard implementation for Linux.
type Linux struct {
	wayland bool
}

// NewLinux creates a new Linux clipboard instance.
func NewLinux() *Linux {
	// Detect Wayland
	_, err := exec.LookPath("wl-copy")
	return &Linux{
		wayland: err == nil,
	}
}

// Copy copies text to the Linux clipboard.
func (l *Linux) Copy(text string) error {
	if !l.Available() {
		return fmt.Errorf("no clipboard command available")
	}

	var cmd *exec.Cmd
	if l.wayland {
		cmd = exec.Command("wl-copy")
	} else {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	}

	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("clipboard command failed: %w", err)
	}

	return nil
}

// Available returns true if a clipboard command is available.
func (l *Linux) Available() bool {
	if l.wayland {
		return true
	}

	// Check for xclip
	_, err := exec.LookPath("xclip")
	return err == nil
}

// Detect Wayland session.
func isWayland() bool {
	return os.Getenv("WAYLAND_DISPLAY") != ""
}
