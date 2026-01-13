package fs

import (
	"log"
	"os/user"
	"path/filepath"
)

type Paths struct {
	Home      string
	Dotfiles  string
	PathsFile string
}

func NewPaths() *Paths {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("error al obtener el usuario actual:", err)
	}

	return &Paths{
		Home:      usr.HomeDir,
		Dotfiles:  filepath.Join(usr.HomeDir, ".dotfiles"),
		PathsFile: filepath.Join(usr.HomeDir, ".dotfiles", ".paths.json"),
	}
}
