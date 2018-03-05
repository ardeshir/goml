package main

import (
	"fmt"
	"log"
	"os"
	
    "github.com/kniren/gota/dataframe"
)

func main() {
	// Open the iris dataset file
	f, err := os.Open("iris_label.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
    irisDF := dataframe.ReadCSV(f)
    
    // Create a filter for the dataframe
    filter := dataframe.F{
    	Colname: "species",
    	Comparator: "==",
    	Comparando: "Iris-versicolor",
    }
    
    // Filter the dataframe to see only the rows where 
    // the iris species is "Iris-versicolor"
    versicolorDF := irisDF.Filter(filter)
    if versicolorDF.Err != nil {
    	log.Fatal(versicolorDF.Err)
    }
    
	fmt.Println("Gota will format the dataframe\n")
	fmt.Println(irisDF)
}
