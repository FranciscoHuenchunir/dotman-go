package internal

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func CreateFile(file string) {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		file, err := os.Create(file)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fmt.Println("archivo creando en: ", file.Name())
	}
}

func HomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error al obtener el usuario actual:", err)
	}
	return usr.HomeDir
}
func DotfilesPath() string {
	home := HomePath()
	return filepath.Join(home, ".dotfiles")
}
func ConfigPath() string {
	dirconfig, err := os.UserConfigDir()

	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dirconfig, "dotman")

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
