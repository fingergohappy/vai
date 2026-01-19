// Package config provides application configuration.
package config

// Config holds application configuration.
type Config struct {
	// Editor settings
	Editor EditorConfig `yaml:"editor"`

	// Keybindings
	Keybindings KeybindingsConfig `yaml:"keybindings"`

	// Theme
	Theme ThemeConfig `yaml:"theme"`
}

// EditorConfig contains editor-related settings.
type EditorConfig struct {
	// TabWidth is the number of spaces per tab.
	TabWidth int `yaml:"tab_width"`

	// WordWrap enables word wrap.
	WordWrap bool `yaml:"word_wrap"`

	// LineNumbers enables line numbers in code blocks.
	LineNumbers bool `yaml:"line_numbers"`
}

// KeybindingsConfig contains custom keybindings.
type KeybindingsConfig struct {
	// Custom keybindings override defaults.
	// TODO: Define keybinding structure
	Overrides map[string]string `yaml:"overrides"`
}

// ThemeConfig contains theme settings.
type ThemeConfig struct {
	// Name of the theme (default, dark, light).
	Name string `yaml:"name"`

	// Custom colors override theme colors.
	Colors map[string]string `yaml:"colors"`
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		Editor: EditorConfig{
			TabWidth:    4,
			WordWrap:    true,
			LineNumbers: true,
		},
		Keybindings: KeybindingsConfig{
			Overrides: make(map[string]string),
		},
		Theme: ThemeConfig{
			Name:   "default",
			Colors: make(map[string]string),
		},
	}
}
