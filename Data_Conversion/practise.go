package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to string
	num := 123
	str := strconv.Itoa(num)
	fmt.Println(str)

	// string to int
	s := "1234"
	n, _ := strconv.Atoi(s)
	fmt.Println(n)

	// string to float 64
	s1 := "3.14"
	f, _ := strconv.ParseFloat(s1, 64)
	fmt.Println(f)
}
