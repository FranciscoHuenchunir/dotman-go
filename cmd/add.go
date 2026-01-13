/*
Copyright © 2025 NAME HERE <franciscohuenchunir42@gmail.com francisco>
*/
package cmd

import (
	"dotman/internal/fs"
	"dotman/internal/model"
	"dotman/internal/validate"
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
	Run: handleAddPathCommand,
}

func move(source, destination string) error {
	return os.Rename(source, destination)
}
func resolveDotfileDestination(sourcePath string, paths fs.Paths) model.DotmanPath {
	dotfilesRoot := paths.Dotfiles
	homeDir := paths.Home

	pathSegments := strings.Split(sourcePath, homeDir)

	if len(pathSegments) == 0 {
		log.Fatal(pathSegments)
	}

	destination := filepath.Join(dotfilesRoot, pathSegments[1])
	return model.DotmanPath{
		Source:      sourcePath,
		Destination: destination,
		Linked:      false,
	}
}

func moveToDotfiles(path model.DotmanPath) {
	destinationDir := filepath.Dir(path.Destination)
	if err := os.MkdirAll(destinationDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := move(path.Source, path.Destination); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✔ Path moved successfully")
	fmt.Println("  ├─ Source:     ", path.Source)
	fmt.Println("  └─ Destination:", path.Destination)
}

func handleAddPathCommand(cmd *cobra.Command, args []string) {
	sourcePath, err := os.Getwd()
	paths := fs.NewPaths()

	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 1 && args[0] == "." {
		dotfilePath := resolveDotfileDestination(sourcePath, paths)
		moveToDotfiles(dotfilePath)
		return
	}
	args, err = validate.CleanUniqueArgs(args)

	if err != nil {
		log.Fatal(err)
	}

	targetPaths, err := fs.Filter(args, sourcePath)

	if err != nil {
		log.Fatal(err)
	}

	for _, target := range targetPaths {
		dotfilePath := resolveDotfileDestination(target, paths)
		moveToDotfiles(dotfilePath)
	}

}

func init() {
	rootCmd.AddCommand(addCmd)
}
