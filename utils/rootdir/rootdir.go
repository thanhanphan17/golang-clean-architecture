package utils

import (
	"os"
	"path/filepath"
)

func FindGoModDir() string {
	// Get the current working directory.
	cwd, _ := os.Getwd()

	// Check for the existence of the go.mod file in the current directory and its parent directories.
	for {
		goModFilePath := filepath.Join(cwd, "go.mod")
		_, err := os.Stat(goModFilePath)
		if err == nil {
			return cwd
		}

		// Move up one directory level.
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			return ""
		}
		cwd = parentDir
	}
}
