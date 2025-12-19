package internal

import (
	"fmt"
	"log"
	"os"
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
