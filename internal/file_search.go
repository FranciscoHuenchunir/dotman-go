package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func Filter(arrDirNames []string, userPath string) ([]string, error) {
	entries, err := os.ReadDir(userPath)
	var pathDir []string

	if err != nil {
		return nil, fmt.Errorf("error al leer directorio %s: %w", userPath, err)
	}
	if len(arrDirNames) == 0 {
		return nil, fmt.Errorf("no hay nombres de directorios o archivos válidos en la lista")
	}

	validName, err := ValidatePaths(arrDirNames)

	for _, entry := range entries {
		if _, exists := validName[entry.Name()]; !exists {
			continue
		}

		pathDir = append(pathDir, filepath.Join(userPath, entry.Name()))

	}
	return pathDir, nil
}
