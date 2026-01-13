package fs_test

import (
	"dotman/internal/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestTildeToHome(t *testing.T) {
	t.Run("obtener path con home", func(t *testing.T) {
		tilde := "~/dotman/cmd"
		expected, err := os.UserHomeDir()

		if err != nil {
			t.Error(err)
		}
		expected = expected + "/dotman/cmd"

		home := fs.TildeToHome(tilde)

		if expected != home {
			t.Error("no se esperaba un error, tenían que ser iguales")
			t.Error(home)
		}
	})

}
func TestHomeToTilde(t *testing.T) {
	t.Run("obtener path con tilde", func(t *testing.T) {
		home, err := os.UserHomeDir()
		expected := "~/dotman/cmd"

		if err != nil {
			t.Error(err)
		}
		home = filepath.Join(home, "dotman/cmd")
		tilde := fs.HomeToTilde(home)

		if expected != tilde {
			t.Error("no se esperaba un error, tenían que ser iguales")
			t.Error(tilde)
		}
	})

}
