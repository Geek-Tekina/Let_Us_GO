package main

import "fmt"

// fmt.Scanf("%[^\n]", &name)

func main() {
	fmt.Println("Hello Peeps, I am learning Golang fundamentals.")
	var name string
	var age int
	fmt.Println("Enter your name, future go developer :) ")
	fmt.Scanln(&name) // only for single word name

	fmt.Println("Enter your age !!")
	fmt.Scan(&age)
	fmt.Printf("Hey %s, let's deep dive into go learning. And Congratulations for wasting %v years on earth, lol", name, age)
}
