// Package config provides application configuration.
package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Loader loads configuration from file.
type Loader struct {
	configPath string
}

// NewLoader creates a new configuration loader.
func NewLoader() *Loader {
	return &Loader{
		configPath: getConfigPath(),
	}
}

// Load loads the configuration from file.
// If the file doesn't exist, returns default config.
func (l *Loader) Load() (Config, error) {
	cfg := DefaultConfig()

	// Check if config file exists
	if _, err := os.Stat(l.configPath); os.IsNotExist(err) {
		// Return default config
		return cfg, nil
	}

	// Read config file
	data, err := os.ReadFile(l.configPath)
	if err != nil {
		return cfg, err
	}

	// Parse YAML
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

// Save saves the configuration to file.
func (l *Loader) Save(cfg Config) error {
	// Ensure directory exists
	dir := filepath.Dir(l.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Marshal to YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(l.configPath, data, 0644)
}

// getConfigPath returns the platform-specific config path.
func getConfigPath() string {
	// Check for XDG_CONFIG_HOME
	if configDir := os.Getenv("XDG_CONFIG_HOME"); configDir != "" {
		return filepath.Join(configDir, "vai", "config.yaml")
	}

	// Default to ~/.config/vai/config.yaml
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "vai", "config.yaml")
}
