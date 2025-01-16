package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
)

type Course struct {
	Name      string
	TimeSpent string
	TotalTime string
}

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
	// when mapping the data, it is important to delte the first row
	rows = slices.Delete(rows, 0, 1) // delete from rows starting from 0 index and delete 1 row

	var courses []Course

	for _, row := range rows {
		newCourse := Course{row[0], row[1], row[2]}
		courses = append(courses, newCourse)
	}

	fmt.Println(courses)
}
