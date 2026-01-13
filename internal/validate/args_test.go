package validate_test

import (
	"dotman/internal/validate"
	"testing"
)

func TestCleanUniqueArgs(t *testing.T) {
	// Test caso exitoso con duplicados
	t.Run("limpia espacios", func(t *testing.T) {
		args := []string{"María  ", "Juan", "Pedro"}
		args, err := validate.CleanUniqueArgs(args)

		if err != nil {
			t.Errorf("no esperaba error, obtuve: %v", err)
		}

		expected := []string{"Juan", "María", "Pedro"}

		if len(args) != len(expected) {
			t.Errorf("esperaba %d elementos, obtuve %d", len(expected), len(args))
			t.Errorf("resultados esperados: %s", expected)
			t.Errorf("obtuve: %s", args)
		}
	})

	// Test con argumento vacío
	t.Run("retorna error con argumento vacío", func(t *testing.T) {
		args := []string{"Juan", "", "María"}
		args, err := validate.CleanUniqueArgs(args)

		if err == nil {
			t.Error("esperaba un error pero no obtuve ninguno")
		}
	})

	// Test con solo espacios
	t.Run("retorna error con solo espacios", func(t *testing.T) {
		args := []string{"Juan", "   ", "María"}
		args, err := validate.CleanUniqueArgs(args)

		if err == nil {
			t.Error("esperaba un error pero no obtuve ninguno")
		}
	})

	// Test con slice vacío
	t.Run("maneja slice vacío", func(t *testing.T) {
		args := []string{}
		args, err := validate.CleanUniqueArgs(args)

		if err != nil {
			t.Errorf("no esperaba error, obtuve: %v", err)
		}

		if len(args) != 0 {
			t.Errorf("esperaba slice vacío, obtuve %d elementos", len(args))
		}
	})
}
