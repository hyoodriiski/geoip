package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SupportedFormats lists all output formats supported by this tool.
var SupportedFormats = []string{
	"text",
	"json",
	"mmdb",
	"dat",
	"csv",
}

// IPType represents the IP address family.
type IPType uint8

const (
	IPv4 IPType = 4
	IPv6 IPType = 6
)

// Entry represents a single GeoIP entry with a name and associated CIDR list.
type Entry struct {
	Name  string
	CIDRs []string
}

// Config holds the top-level configuration parsed from config.json.
type Config struct {
	Input  []InputConfig  `json:"input"`
	Output []OutputConfig `json:"output"`
}

// InputConfig describes a single input source.
type InputConfig struct {
	Type   string                 `json:"type"`
	Action string                 `json:"action"`
	Args   map[string]interface{} `json:"args"`
}

// OutputConfig describes a single output destination.
type OutputConfig struct {
	Type   string                 `json:"type"`
	Action string                 `json:"action"`
	Args   map[string]interface{} `json:"args"`
}

// LoadConfig reads and parses the JSON configuration file at the given path.
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %q: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %q: %w", path, err)
	}

	return &cfg, nil
}

// NormalizeName converts a list name to uppercase and trims whitespace,
// ensuring consistent naming across all input/output operations.
func NormalizeName(name string) string {
	return strings.ToUpper(strings.TrimSpace(name))
}

// IsValidFormat returns true if the given format string is supported.
func IsValidFormat(format string) bool {
	for _, f := range SupportedFormats {
		if strings.EqualFold(f, format) {
			return true
		}
	}
	return false
}

// EnsureDir creates the directory at path if it does not already exist.
func EnsureDir(path string) error {
	if err := os.MkdirAll(path, 0750); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", path, err)
	}
	return nil
}
