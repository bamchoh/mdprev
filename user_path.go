// +build !windows

package main

import (
	"os"
	"path/filepath"
)

func findUserConfigPath() string {
	home := os.Getenv("HOME")
	dir := filepath.Join(home, ".config")

	appName := filepath.Base(os.Args[0])
	return filepath.Join(dir, appName)
}
