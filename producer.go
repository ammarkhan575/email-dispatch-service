package main

import (
	"encoding/csv"
	"os"
)

func loadReciepients(filePath string, reciepients chan Reciepient) []Reciepient {
	// This function would contain logic to read a file and parse the reciepients
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	recipients := []Reciepient{}
	for _, record := range records[1:] { // Skip header
		reciepients <- Reciepient{Name: record[0], Email: record[1]}
		// fmt.Printf("Loaded reciepient: %s <%s>\n", record[0], record[1])
		// recipients = append(recipients, Reciepient{Name: record[0], Email: record[1]})
	}
	return recipients
}
