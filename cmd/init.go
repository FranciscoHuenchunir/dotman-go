/*
Copyright © 2025 NAME HERE <franciscohuenchunir42@gmail.com francisco>
*/
package cmd

import (
	"dotman/internal/fs"
	"fmt"
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
func createDataPaths(pathFile string) {
	fs.CreateFile(pathFile)

}

func initDotman(cmd *cobra.Command, args []string) {
	paths := fs.NewPaths()

	initDotfilesDir(paths.Dotfiles)
	createDataPaths(paths.PathsFile)
}

func init() {
	rootCmd.AddCommand(initCmd)
}
