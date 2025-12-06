package pathutil

import (
	"log"
	"os/user"
	"path/filepath"
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
