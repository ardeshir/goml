package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"io"
)

func main() {
	// Open the iris dataset file
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file
	reader := csv.NewReader(f)

	// Assume we don't want the number of fileds per line. By setting
	// FieldsPerRecord negative, ech row may have a variable number of fields
	reader.FieldsPerRecord = -1

	// Read in all the csv records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("First method, reader.ReadAll() at once:\n")
	fmt.Println(rawCSVData)
	fmt.Println("\n\n")
	
	// second method
	f2, err := os.Open("iris.csv")
        if err != nil {
                log.Fatal(err)
        }
        defer f2.Close()

        // Create a new CSV reader reading from the opened file
        reader2 := csv.NewReader(f2)


	// rawCSVData2 will hold our succefully parsed rows
	var rawCSVData2 [][]string

	// read in the records one by one
	for {
		// Read in a row. Check if we are at the end of the file
		record, err := reader2.Read()
		if err == io.EOF {
			break
		}
		// Append the record to our data set
		rawCSVData2 = append(rawCSVData2, record)
	}
	fmt.Println("Second method, reader.Read() one row at a time appended to slice of slices:\n")
	fmt.Println(rawCSVData2)
}
