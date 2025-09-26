package queries

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type QueryFile struct {
	Category string
	Name     string
	Path     string
}

func Load(root string) ([]QueryFile, error) {
	var list []QueryFile
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.EqualFold(filepath.Ext(d.Name()), ".kql") {
			rel, _ := filepath.Rel(root, path)
			dir := filepath.Dir(rel)
			category := dir
			if category == "." {
				category = "uncategorized"
			}
			list = append(list, QueryFile{Category: category, Name: d.Name(), Path: path})
		}
		return nil
	})
	return list, err
}
