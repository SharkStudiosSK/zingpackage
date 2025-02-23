package zing

import "fmt"

// Placeholder for configuration management.  Could use a struct to hold
// configuration values (e.g., repository URL).

// For now, we'll hardcode the repository URL.
func GetRepositoryURL() string {
	return "https://github.com/SharkStudiosSK/zinglets-repo.git" // Replace with the actual URL
}

// GetZingletURL returns the URL for a specific zinglet.
func GetZingletURL(zingletName string) string {
	baseURL := GetRepositoryURL()
	return fmt.Sprintf("%s/%s", baseURL, zingletName)
}
