package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord contains a sucessfully parsed row of the CSV file
type CSVRecord struct {
	SepalLength  float64
	SepalWidth   float64
	PetalLength	 float64
	PetalWidth	 float64
	Species      string
	ParseError   error
}

func main() {
	// Open the iris dataset file
	f, err := os.Open("iris_mix.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file
	reader := csv.NewReader(f)
	
	// Create a slice value that will hold all of the successful parsed 
	// records from CSV
	var csvData []CSVRecord 
	
	//  line will help us keep track of line number of logging
	line := 1
	
	for {
		  // Read in a row. Check if we are at the end of file. 
		  record, err := reader.Read()
		  if err == io.EOF {
		  	  break
		  }
		  // Create a CSVRecord value for the row.
		  var csvRecord CSVRecord 
		  
		  // Parse each of the values in the record based on an exprected type
		  for idx, value := range record {
		  	     //  Rase the value in the record 
		  	     if idx == 4 {
		  	     	  // validate that the value is not an empty string
		  	     	    if value == "" {
		  	     	    	   log.Printf("Parsing line %d failed, unexpected type in column %d\n", line,  idx)
		  	     	    	   csvRecord.ParseError = fmt.Errorf("Empty string value")
		  	     	    	   break
		  	     	    }
		  	     	    
		  	     	    // Add string value to the CSVRecord
		  	     	    csvRecord.Species = value
		  	     	    continue 
		  	     }
		  	     
		  	     // Otherwise, parse the value in the record as a float64
		  	     // floatValue will hold the parsed float value of the record
		  	     // for the numeric columns
		  	     var floatValue float64
		  	     
		  	     // if the value cannot be parsed as a float, log and break the parsing loop
		  	     if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
		  	     	  log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
		  	     	  
		  	     	  csvRecord.ParseError = fmt.Errorf("Could not parse float")
		  	     	  break 
		  	     }
		  	     
		  	      // add the float value to the respective filed in the CSVRecord
		  	      switch idx {
		  	      	
		  	      	case 0:
		  	      			csvRecord.SepalLength = floatValue
		  	      	case 1: 
		  	      			csvRecord.SepalWidth  = floatValue
		  	      	case 2: 
		  	      			csvRecord.PetalLength = floatValue
		  	      	case 3: 
		  	      	        csvRecord.PetalWidth  = floatValue
		  	      } 
		  	      
		  	      // Append successfully parsed records to the slice defined above.
		  	      if csvRecord.ParseError == nil {
		  	      	    csvData = append(csvData, csvRecord)
		  	      }
		  	      
		  	      // increment line counter
		  	      line++
		  	       
		  }
	}
	
	fmt.Printf("Successuflly parse %d lines\n", len(csvData))


	// Assume we don't want the number of fileds per line. By setting
	// FieldsPerRecord negative, ech row may have a variable number of fields
	/*
	reader.FieldsPerRecord = -1

	// Read in all the csv records
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("First method, reader.ReadAll() at once:\n")
	fmt.Println(rawCSVData)
	fmt.Println("\n\n")
	*/ 
	
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
