package internal

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"strings"
)

func DotfilesPath() string {
	home := HomePath()
	return filepath.Join(home, ".dotfiles")
}
func HomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error al obtener el usuario actual:", err)
	}
	return usr.HomeDir
}
func ValidatePaths(paths []string) (map[string]struct{}, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("no hay nombres de directorio en la lista")
	}

	validNames := make(map[string]struct{})

	for _, name := range paths {
		n := strings.TrimSpace(name)
		if n == "" {
			return nil, fmt.Errorf("un nombre está vacío o solo tiene espacios")
		}

		validNames[n] = struct{}{}
	}

	return validNames, nil
}
