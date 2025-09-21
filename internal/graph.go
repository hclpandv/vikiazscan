package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

func ExecuteKQLQuery(kqlFile string) ([]OrphanedResource, error) {
	// Read KQL file
	queryBytes, err := os.ReadFile(kqlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read KQL file: %v", err)
	}
	query := string(queryBytes)

	// Authenticate with Azure
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get Azure credential: %v", err)
	}

	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create ARG client: %v", err)
	}

	// Prepare ARG query request
	request := armresourcegraph.QueryRequest{
		Query:         &query,
		Subscriptions: nil, // optional: can specify subscription IDs
	}

	resp, err := client.Resources(context.Background(), request, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to run ARG query: %v", err)
	}

	var resources []OrphanedResource

	// ARG response is in JSON format; parse it
	if resp.Data != nil {
		// Convert to generic JSON
		var data []map[string]interface{}
		rawData, err := json.Marshal(resp.Data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal ARG response: %v", err)
		}
		if err := json.Unmarshal(rawData, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ARG response: %v", err)
		}

		for _, r := range data {
			resources = append(resources, OrphanedResource{
				Name:          fmt.Sprintf("%v", r["name"]),
				ResourceGroup: fmt.Sprintf("%v", r["resourceGroup"]),
				Type:          fmt.Sprintf("%v", r["type"]),
				Location:      fmt.Sprintf("%v", r["location"]),
				SKUName:       fmt.Sprintf("%v", r["skuName"]),
				DiskSize:      int(r["diskSizeGB"].(float64)),
				Tags:          fmt.Sprintf("%v", r["tags"]),
			})

		}
	}

	return resources, nil
}
