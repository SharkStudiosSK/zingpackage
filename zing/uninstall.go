package zing

import (
	"fmt"
	"os"
)

// Uninstall removes an installed zinglet.
func Uninstall(zingletName string) error {
	// Construct a Zinglet object to use the InstalledPath method
	zinglet := Zinglet{Name: zingletName}
	installDir := zinglet.InstalledPath()

	// Check if the zinglet is installed
	if _, err := os.Stat(installDir); os.IsNotExist(err) {
		return fmt.Errorf("zinglet '%s' is not installed", zingletName)
	}

	// Remove the installed directory
	err := os.RemoveAll(installDir)
	if err != nil {
		return fmt.Errorf("failed to uninstall zinglet '%s': %v", zingletName, err)
	}

	fmt.Printf("Zinglet '%s' uninstalled successfully.\n", zingletName)
	return nil
}
