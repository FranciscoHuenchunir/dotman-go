package fs

import (
	"path/filepath"
	"strings"
)

func TildeToHome(path string) string {
	paths := NewPaths()
	if path == "~" {
		return paths.Home
	}
	if path == paths.Home {
		return path
	}
	return filepath.Join(paths.Home, path[2:])
}
func HomeToTilde(path string) string {
	paths := NewPaths()
	if rest, ok := strings.CutPrefix(path, paths.Home); ok {
		return "~" + rest
	}

	return path
}
