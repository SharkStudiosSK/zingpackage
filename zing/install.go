package zing

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Install handles the installation of a zinglet.
func Install(zingletName string) error {
	fmt.Printf("Installing zinglet: %s\n", zingletName)

	// 1. Clone or update the entire zinglets repository.
	err := CloneOrUpdate(zingletName)
	if err != nil {
		return err
	}

	// 2. Determine the path to the *zinglet's subdirectory* within the repository.
	zingHome := os.Getenv("HOME") + "/.zing"
	repoDir := filepath.Join(zingHome, "zinglets", "zinglets-repo")
	zingletDir := filepath.Join(repoDir, zingletName)

	// 3. Check if the zinglet directory exists within the repository.
	if _, err := os.Stat(zingletDir); os.IsNotExist(err) {
		return fmt.Errorf("zinglet '%s' not found in repository", zingletName)
	}

	// 4. Read the zinglet manifest (zinglet.json).
	var zinglet Zinglet
	if err := zinglet.LoadFromDir(zingletDir); err != nil {
		return err
	}

	// 5. Create the destination directory.
	installDir := zinglet.InstalledPath()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create installation directory: %v", err)
	}

	// 6. Copy files from the source to the destination directory.
	err = copyFiles(zingletDir, installDir)
	if err != nil {
		return fmt.Errorf("failed to copy files: %v", err)
	}

	fmt.Printf("Zinglet %s (%s) installed successfully to %s\n", zinglet.Name, zinglet.Version, installDir)
	return nil
}

// copyFiles recursively copies files and directories.
func copyFiles(srcDir, destDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(srcDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
			if err := copyFiles(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// Skip zinglet.json when copying
			if entry.Name() == "zinglet.json" {
				continue
			}
			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file.
func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}
