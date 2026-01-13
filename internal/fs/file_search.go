package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func Filter(names []string, userPath string) ([]string, error) {
	entries, err := os.ReadDir(userPath)
	var paths []string

	if err != nil {
		return nil, fmt.Errorf("error al leer el directorio %s: %w", userPath, err)
	}
	for _, entry := range entries {
		if !slices.Contains(names, entry.Name()) {
			continue
		}
		paths = append(paths, filepath.Join(userPath, entry.Name()))

	}
	return paths, nil
}
