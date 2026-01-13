package store_test

import (
	"dotman/internal/fs"
	"dotman/internal/model"
	"dotman/internal/store"
	"os"
	"path/filepath"
	"testing"
)

var testPaths = fs.Paths{
	PathsFile: filepath.Join("testdata", "paths.test.json"),
}

func TestSavePaths(t *testing.T) {
	t.Run("validar que se cree el archivo", func(t *testing.T) {

		defer os.Remove(testPaths.PathsFile)

		err := store.SavePaths(testPaths, []model.DotmanPath{})
		if err != nil {
			t.Fatal(err)
		}

		_, err = os.Stat(testPaths.PathsFile)
		if err != nil {
			t.Error("el archivo no se creo")
			t.Error(testPaths.PathsFile)
		}
	})

	t.Run("validar que se guardó el contenido", func(t *testing.T) {

		defer os.Remove(testPaths.PathsFile)
		paths := []model.DotmanPath{
			{
				Source:      "/tmp/source",
				Destination: "/tmp/dest",
				Linked:      false,
			},
		}

		err := store.SavePaths(testPaths, paths)
		if err != nil {
			t.Fatal(err)
		}

		data, err := os.ReadFile(testPaths.PathsFile)
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
