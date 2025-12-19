package internal

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Paths struct {
	Home          string
	Dotfiles      string
	Config        string
	DotIngoreFile string
	PathsFile     string
	ConfigFile    string
}

func NewPaths() *Paths {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("error al obtener el usuario actual:", err)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	return &Paths{
		Home:          usr.HomeDir,
		Dotfiles:      filepath.Join(usr.HomeDir, ".dotfiles"),
		Config:        filepath.Join(configDir, "dotman"),
		DotIngoreFile: filepath.Join(usr.HomeDir, ".dotfiles", ".dotignore"),
		PathsFile:     filepath.Join(usr.HomeDir, ".dotfiles", ".paths.yml"),
		ConfigFile:    filepath.Join(configDir, "dotman", "config.toml"),
	}
}
