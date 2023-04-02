package main

import (
	"encoding/json"
	"log"
)

func parseFile(fileContent []byte) ([]byte, error) {
	// Ensure the uploaded JSON is valid
	log.Print("Ensuring JSON is valid")
	if !json.Valid(fileContent) {
		log.Fatal("Invalid JSON provided")
	}

	// Create a new string mapping for json data with any type (as json values can be multiple types)
	log.Print("Unmarshalling JSON data into string map")
	var result map[string]any
	if err := json.Unmarshal(fileContent, &result); err != nil {
		log.Print("Error unmarshalling JSON: ", err)
		return []byte(""), err
	}

	log.Print("Finding number of components")
	components := result["components"].([]interface{})
	numComponents := len(components)
	log.Printf("Number of components found: %d", numComponents)

	log.Print("Remarshalling JSON data")
	returnData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return []byte(""), err
	}

	log.Print("Returning JSON data")
	return returnData, nil
}
