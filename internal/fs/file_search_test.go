package fs_test

import (
	"dotman/internal/fs"
	"os"
	"testing"
)

func TestFilter(t *testing.T) {
	testPath, _ := os.Getwd()
	t.Run("obtener paths de los achivos o direcotios por nombre", func(t *testing.T) {
		names := []string{"file.go", "paths.go"}
		entries, _ := fs.Filter(names, testPath)

		if len(entries) == 0 {
			t.Error("no esperaba que estuviera vacio ")
		}
		t.Logf("resultado: \n%s", entries)
	})
	t.Run("validar la existencia de los path obtenido por el nombre", func(t *testing.T) {
		names := []string{"file.go", "paths.go"}
		entries, _ := fs.Filter(names, testPath)

		for _, entry := range entries {
			t.Log(entry)

			if _, err := os.Stat(entry); err != nil {
				t.Errorf("el path %s no es valido", entry)
			}
		}
	})
	t.Run("espera un error, se proporciona un directorio que no existe", func(t *testing.T) {
		names := []string{"file.go", "paths.go"}
		_, err := fs.Filter(names, "/home/test/")

		if err == nil {
			t.Errorf("Se esperaba un error")
		}
	})

}
