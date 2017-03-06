package main

import (
	"os"
	"path/filepath"
)

func findUserConfigPath() string {
	home := os.Getenv("USERPROFILE")
	dir := os.Getenv("APPDATA")
	if dir == "" {
		dir = filepath.Join(home, "Application Data")
	}

	basename := filepath.Base(os.Args[0])
	ext := filepath.Ext(basename)
	appName := basename[:(len(basename) - len(ext))]
	return filepath.Join(dir, appName)
}
