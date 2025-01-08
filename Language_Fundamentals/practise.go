package main

import "fmt"

// fmt.Scanf("%[^\n]", &name)

func main() {
	// -------------------- Printing and Input ----------------------------
	fmt.Println("Hello Peeps, I am learning Golang fundamentals.")
	var name string
	var age int
	fmt.Println("Enter your name, future go developer :) ")
	fmt.Scanln(&name) // only for single word name
	// var discard string
	// fmt.Scanln(&discard) // is this the right way ??????
	fmt.Println("Enter your age !!")
	fmt.Scan(&age)
	fmt.Printf("Hey %s, let's deep dive into go learning. And Congratulations for wasting %v years on earth, lol", name, age)
	// -------------------- Printing and Input ----------------------------
	// -------------------- Control Structures ---------------------------
	mathsMarks := 90
	engMarks := 25

	if mathsMarks+engMarks > 150 {
		fmt.Println("Ooo Yeah, Passed !!")
	} else {
		fmt.Println("Fuck you beattchhhhhh")
	}

	role := "Reader" // can be used in role based authorization
	switch role {
	case "Admin":
		fmt.Println("You have all the access.")
		fallthrough
	case "Writer":
		fmt.Println("You have access of writing and reading both.")
		fallthrough
	case "Reader":
		fmt.Println("You have only access to reading.")
	}

	for i := 1; i < 5; i++ {
		fmt.Print("Hi-")
	}
	// -------------------- Control Structures ---------------------------

}
