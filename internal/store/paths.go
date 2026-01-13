package store

import (
	"dotman/internal/model"
	"encoding/json"
	"os"
)

func LoadPaths(file string) ([]model.DotmanPath, error) {
	// modelPaths := fs.NewPaths()
	// file := modelPaths.PathsFile
	//
	data, err := os.ReadFile(file)
	if os.IsNotExist(err) {
		return []model.DotmanPath{}, nil
	}
	if err != nil {
		return nil, err
	}

	var dotmanPaths []model.DotmanPath
	if len(data) > 0 {
		err = json.Unmarshal(data, &dotmanPaths)
		if err != nil {
			return nil, err
		}
	}

	return dotmanPaths, nil
}

func SavePaths(file string, paths []model.DotmanPath) error {
	// modelPaths := fs.NewPaths()
	// file := modelPaths.PathsFile

	data, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0644)
}
