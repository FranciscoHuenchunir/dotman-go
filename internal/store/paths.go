package store

import (
	"dotman/internal/fs"
	"dotman/internal/model"
	"encoding/json"
	"os"
)

func LoadPaths(paths fs.Paths) ([]model.DotmanPath, error) {

	data, err := os.ReadFile(paths.PathsFile)
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

func SavePaths(paths fs.Paths, items []model.DotmanPath) error {

	data, err := json.MarshalIndent(paths, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(paths.PathsFile, data, 0644)
}
