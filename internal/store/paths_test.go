package store_test

import (
	"dotman/internal/model"
	"dotman/internal/store"
	"os"
	"path/filepath"
	"testing"
)

func TestSavePaths(t *testing.T) {
	var file string = filepath.Join("testdata", "paths.test.json")
	t.Run("validar que se cree el archivo", func(t *testing.T) {

		defer os.Remove(file)

		err := store.SavePaths(file, []model.DotmanPath{})
		if err != nil {
			t.Fatal(err)
		}

		_, err = os.Stat(file)
		if err != nil {
			t.Error("el archivo no se creo")
		}
	})

	t.Run("validar que se guardó el contenido", func(t *testing.T) {

		defer os.Remove(file)
		paths := []model.DotmanPath{
			{
				Name:        "tmp",
				Source:      "/tmp/source",
				Destination: "/tmp/dest",
				Linked:      false,
			},
		}

		err := store.SavePaths(file, paths)
		if err != nil {
			t.Fatal(err)
		}

		data, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}

		if len(data) == 0 {
			t.Fatal("archivo vacío")
		}
	})

}

func TestLoadPaths(t *testing.T) {
}
