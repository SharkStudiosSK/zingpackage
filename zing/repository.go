package zing

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CloneOrUpdate clones or updates the *entire* zinglets repository.
func CloneOrUpdate(zingletName string) error {
	repoURL := GetRepositoryURL() // Get the repository URL from config.go
	zingHome := os.Getenv("HOME") + "/.zing"
	targetDir := filepath.Join(zingHome, "zinglets", "zinglets-repo") // Consistent directory for the entire repo.

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		// Clone the repository
		fmt.Printf("Cloning zinglets repository from %s to %s\n", repoURL, targetDir)
		cmd := exec.Command("git", "clone", repoURL, targetDir)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	} else {
		// Update the repository (assuming it's already cloned)
		fmt.Printf("Updating zinglets repository in %s\n", targetDir)
		cmd := exec.Command("git", "-C", targetDir, "pull")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}
