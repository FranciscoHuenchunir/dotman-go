/*
Copyright © 2025 NAME HERE <franciscohuenchunir42@gmail.com francisco>
*/
package cmd

import (
	"dotman/internal/fs"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.MinimumNArgs(1),
	Short: "Adds a file or directory to the .dotfiles repository",
	Long: `The add command moves the specified file or directory into the ~/.dotfiles
		directory while preserving its relative path. If the 'symlink_on_add' option is
		enabled in the configuration, a symbolic link will be created at the original location.`,
	Run: addPathToDotfile,
}

func mover(origen, destino string) error {
	fmt.Println("se movio")
	return os.Rename(origen, destino)
}
func addToDotfiles(userPath, home, dotPath string) {

	// Get relative path from home
	relativePath := strings.Split(userPath, home)

	if len(relativePath) == 0 {
		log.Fatal(relativePath)
	}

	// Build complete destination path
	fullDestination := filepath.Join(dotPath, relativePath[1])

	// Create destination directory if it doesn't exist
	destinationDir := filepath.Dir(fullDestination)
	if err := os.MkdirAll(destinationDir, 0755); err != nil {
		log.Fatal(err)
	}
	// // Move file/directory
	if err := mover(userPath, fullDestination); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Source: ", userPath)
	fmt.Println("Destination: ", fullDestination)
	fmt.Println("✓ Moved successfully")

}
func addPathToDotfile(cmd *cobra.Command, args []string) {
	paths := *fs.NewPaths()

	home := paths.Home
	dotPath := paths.Dotfiles

	userPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Source: ", userPath)

	if len(args) == 1 && args[0] == "." {
		addToDotfiles(userPath, home, dotPath)
	}

	if len(args) >= 2 {

		dirPaths, err := fs.Filter(args, userPath)

		if err != nil {
			log.Fatal(err)
		}
		for _, dirPath := range dirPaths {
			addToDotfiles(dirPath, home, dotPath)
		}

	}
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
