package zing

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CloneOrUpdateZinglet clones or updates the specific zinglet directory.
func CloneOrUpdateZinglet(zingletName string) error {
	repoURL := GetZingletURL(zingletName) // Get the specific zinglet URL from config.go
	zingHome := os.Getenv("HOME") + "/.zing"
	targetDir := filepath.Join(zingHome, "zinglets", zingletName) // Consistent directory for the specific zinglet.

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// Clone the specific zinglet
		fmt.Printf("Cloning zinglet from %s to %s\n", repoURL, targetDir)
		cmd := exec.Command("git", "clone", repoURL, targetDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		// Update the specific zinglet (assuming it's already cloned)
		fmt.Printf("Updating zinglet in %s\n", targetDir)
		cmd := exec.Command("git", "-C", targetDir, "pull")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}
