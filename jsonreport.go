package main

import (
	"encoding/json"
	"sort"
	"fmt"
	"os"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	//Sort map keys
	keys := make([]string, 0, len(pages)) // Create empty slice to store sorted strings

    for k := range pages{ // Fill slice with unsorted keys
        keys = append(keys, k)
    }
	sort.Strings(keys) // Sort keys



	// Build a []PageData slice in sorted order
	sortedPageData := make([]PageData, 0, len(pages))

	for _, k := range keys {
		sortedPageData = append(sortedPageData, pages[k])
	}

	
	// Marshal it with json.MarshalIndent using indent=2

	data, err := json.MarshalIndent(sortedPageData, "", " ")
	if err != nil {
		fmt.Println("ERROR -- Unable to marshal page data: %v\n", err)
		return err
	}
	// Write the result to disk with os.WriteFile

	os.WriteFile(filename, data, 0644)

	return nil
}