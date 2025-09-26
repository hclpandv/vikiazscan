package queries

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hclpandv/vikiazscan/internal"
)

// LoadQueriesFromFolders reads queries from the given root folder and groups them by subfolder
func LoadQueriesFromFolders(root string) (map[string][]internal.TableData, error) {
	tabs := make(map[string][]internal.TableData)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		// Only process .kql files
		if !strings.HasSuffix(d.Name(), ".kql") {
			return nil
		}

		// Determine folder name (category)
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		folder := filepath.Dir(rel)
		if folder == "." {
			folder = "uncategorized"
		}

		// Read query content (if needed) or just store name for report
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// For demo, just create empty headers/rows; you would run the query here
		table := internal.TableData{
			Name:    strings.TrimSuffix(d.Name(), ".kql"),
			Headers: []string{"Column1", "Column2"},    // replace with actual headers
			Rows:    [][]string{{string(content), ""}}, // replace with actual rows
		}

		tabs[folder] = append(tabs[folder], table)
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tabs, nil
}
