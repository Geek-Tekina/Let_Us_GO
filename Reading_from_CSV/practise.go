package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("We will implement reading from a CSV File.")
	f, err := os.Open("./course.csv") // opening the file

	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(f) // creating a reader object to read from CSV file

	rows, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for i, row := range rows {
		fmt.Println("Row", i+1, "->", row)
	}
}
