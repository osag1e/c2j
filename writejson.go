package main

import (
	"encoding/json"
	"os"
)

func writeJSONToFile(data interface{}, jsonFilePath string) error {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
