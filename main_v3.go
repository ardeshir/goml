package main

import (
	"fmt"
	"log"
	"os"
	
    "github.com/kniren/gota/dataframe"
)

func main() {
	// Open the iris dataset file
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
    irisDF := dataframe.ReadCSV(f)
    
	fmt.Println("Gota will format the dataframe\n")
	fmt.Println(irisDF)
}
