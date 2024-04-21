package main

import (
	"encoding/csv"
	"os"
)

func readCSVFile(csvFilePath string) ([][]string, error) {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
