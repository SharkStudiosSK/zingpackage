package zing

import (
	"fmt"
	"os"
	"path/filepath"
)

// List lists all installed zinglets.
func List() error {
	zingHome := os.Getenv("HOME") + "/.zing"
	installedDir := filepath.Join(zingHome, "installed") // List from the installed directory

	// Check if the installed directory exists
	if _, err := os.Stat(installedDir); os.IsNotExist(err) {
		fmt.Println("No zinglets installed yet.")
		return nil
	}

	// Read the directory entries
	entries, err := os.ReadDir(installedDir)
	if err != nil {
		return fmt.Errorf("failed to read installed zinglets directory: %v", err)
	}

	// Print the names and versions of installed zinglets
	if len(entries) == 0 {
		fmt.Println("No zinglets installed yet.")
		return nil
	}

	fmt.Println("Installed zinglets:")
	for _, entry := range entries {
		if entry.IsDir() { // Make sure it's a directory
			zingletDir := filepath.Join(installedDir, entry.Name())
			var zinglet Zinglet
			if err := zinglet.LoadFromDir(zingletDir); err == nil { // Load metadata
				fmt.Printf("%s (%s)\n", zinglet.Name, zinglet.Version)
			} else {
				fmt.Printf("%s (Error loading meta %v)\n", entry.Name(), err) // Show name even if error
			}
		}
	}

	return nil
}
