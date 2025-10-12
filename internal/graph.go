package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

// ExecuteKQLQuery runs a KQL file and returns headers and rows
func ExecuteKQLQuery(kqlFile string) ([]string, [][]string, error) {
	queryBytes, err := os.ReadFile(kqlFile)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read KQL file: %v", err)
	}
	query := string(queryBytes)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get Azure credential: %v", err)
	}

	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create ARG client: %v", err)
	}

	request := armresourcegraph.QueryRequest{
		Query:         &query,
		Subscriptions: nil,
	}

	resp, err := client.Resources(context.Background(), request, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run ARG query: %v", err)
	}

	// Convert resp.Data to []map[string]interface{}
	b, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal resp.Data: %v", err)
	}

	var arr []map[string]interface{}
	if err := json.Unmarshal(b, &arr); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal resp.Data: %v", err)
	}

	// Build headers
	var headers []string
	if len(arr) > 0 {
		for h := range arr[0] {
			headers = append(headers, h)
		}
	}

	// Build rows
	var rows [][]string
	for _, m := range arr {
		var r []string
		for _, h := range headers {
			v, ok := m[h]
			if !ok || v == nil {
				r = append(r, "")
				continue
			}
			switch val := v.(type) {
			case map[string]interface{}, []interface{}:
				j, _ := json.Marshal(val)
				r = append(r, string(j))
			default:
				r = append(r, fmt.Sprintf("%v", val))
			}
		}
		rows = append(rows, r)
	}

	return headers, rows, nil
}

// GetKQLFiles scans a folder and returns all .kql files
func GetKQLFiles(folder string) []string {
	var files []string
	err := filepath.WalkDir(folder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".kql") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("Error reading queries folder: %v", err))
	}
	return files
}
