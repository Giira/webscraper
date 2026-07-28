package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	var keys []string
	for k := range pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sorted []PageData
	for _, key := range keys {
		sorted = append(sorted, pages[key])
	}

	data, err := json.MarshalIndent(sorted, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	return nil
}
