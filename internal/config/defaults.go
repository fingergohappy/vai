// Package config provides application configuration.
package config

import (
	"os"
	"path/filepath"
)

// Defaults provides default values for configuration.
type Defaults struct{}

// EditorDefaults returns default editor configuration.
func EditorDefaults() EditorConfig {
	return EditorConfig{
		TabWidth:    4,
		WordWrap:    true,
		LineNumbers: true,
	}
}

// ThemeDefaults returns default theme configuration.
func ThemeDefaults() ThemeConfig {
	return ThemeConfig{
		Name: "default",
		Colors: map[string]string{
			"normal_mode": "252",
			"insert_mode": "142",
			"visual_mode": "33",
			"user_text":   "244",
			"assist_text": "252",
		},
	}
}

// GetDataDir returns the platform-specific data directory.
func GetDataDir() string {
	// Check for XDG_DATA_HOME
	if dataDir := os.Getenv("XDG_DATA_HOME"); dataDir != "" {
		return dataDir
	}

	// Default to ~/.local/share/vai
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".local", "share", "vai")
}

// GetSessionsDir returns the sessions data directory.
func GetSessionsDir() string {
	return filepath.Join(GetDataDir(), "sessions")
}
