package main

import (
	"fmt"
)

// ANSI escape sequences
const (
	greenColor  = "\033[1;32m"
	purpleColor = "\033[1;34m"
	resetColor  = "\033[0m"
)

func main() {

	c2jArt := `
   _________       __
  / ____/__ \     / /
 / /    __/ /__  / / 
/ /___ / __// /_/ /  
\____//____/\____/   
`

	fmt.Printf("%s%sEnter C2J file path:%s\n", greenColor, c2jArt, resetColor)
	var csvFilePath string
	_, err := fmt.Scanln(&csvFilePath)
	if err != nil {
		panic("Error reading CSV file path: " + err.Error())
	}

	// Prompt user to enter JSON file path
	fmt.Printf("%sEnter JSON file path:%s\n", greenColor, resetColor)
	var jsonFilePath string
	_, err = fmt.Scanln(&jsonFilePath)
	if err != nil {
		panic("Error reading JSON file path: " + err.Error())
	}

	// Read the CSV file.
	records, err := readCSVFile(csvFilePath)
	if err != nil {
		panic("Error reading CSV file: " + err.Error())
	}

	// Write the records to a JSON file.
	err = writeJSONToFile(records, jsonFilePath)
	if err != nil {
		panic("Error writing JSON file: " + err.Error())
	}

	fmt.Printf("%sCSV to JSON conversion complete.\nJSON file saved at: %s%s\n", purpleColor, jsonFilePath, resetColor)
}
