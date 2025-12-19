/*
Copyright © 2025 NAME HERE <franciscohuenchunir42@gmail.com francisco>
*/
package cmd

import (
	"dotman/internal/fs"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the dotfiles repository",
	Long: `Initialize the required structure for dotman to operate correctly.

	This command creates the ~/.dotfiles directory (if it does not exist) and 
	sets up the base structure needed to manage your configuration files. 
	It may also auto-configure Git so you can easily version and synchronize 
	your dotfiles across systems.

	Example:

  	dotman init

	After running this command, you can begin adding files using:

  dotman add <file>`,
	Run: initDotman,
}

func initDotfilesDir(dotPath string) {
	_, err := os.Stat(dotPath)

	if os.IsNotExist(err) {

		err = os.Mkdir(dotPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Carpeta creada en:", dotPath)
	}
}

func createDotignore(pathFile string) {
	internal.CreateFile(pathFile)

	content := `
.git/
.paths.yml
`
	err := os.WriteFile(pathFile, []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func createDataPaths(pathFile string) {
	internal.CreateFile(pathFile)

}

func initConfigDir(confPath string) {
	_, err := os.Stat(confPath)

	if os.IsNotExist(err) {
		err = os.Mkdir(confPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Carpeta creada en:", confPath)

	}
}
func createConfigFile(pathFile string) {
	_, err := os.Stat(pathFile)

	if os.IsNotExist(err) {
		file, err := os.Create(pathFile)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		content := `
symlink_on_add = false
dotfiles_dir = "~/.dotfiles"

ignore_file = ".dotmanignore"
`

		err = os.WriteFile(pathFile, []byte(content), 0666)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("archivo creando en: ", file.Name())

	}

}

func initDotman(cmd *cobra.Command, args []string) {
	paths := internal.NewPaths()

	initDotfilesDir(paths.Dotfiles)
	initConfigDir(paths.Config)

	createDotignore(paths.DotIngoreFile)
	createDataPaths(paths.PathsFile)

	createConfigFile(paths.ConfigFile)

}

func init() {
	rootCmd.AddCommand(initCmd)
}
