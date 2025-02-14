package zing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Zinglet represents a Zing package.
type Zinglet struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// Add other fields as needed (e.g., Description, Dependencies).
}

// LoadFromDir loads zinglet metadata from a zinglet.json file in the given directory.
func (z *Zinglet) LoadFromDir(dir string) error {
	manifestPath := filepath.Join(dir, "zinglet.json")
	file, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read zinglet.json: %v", err)
	}

	err = json.Unmarshal(file, z)
	if err != nil {
		return fmt.Errorf("failed to parse zinglet.json: %v", err)
	}

	if z.Name == "" {
		return fmt.Errorf("zinglet.json must contain a 'name' field")
	}
	if z.Version == "" {
		return fmt.Errorf("zinglet.json must contain a 'version' field")
	}

	return nil
}

// InstalledPath returns the installation path for a zinglet.
func (z *Zinglet) InstalledPath() string {
	zingHome := os.Getenv("HOME") + "/.zing"
	return filepath.Join(zingHome, "installed", z.Name)
}
