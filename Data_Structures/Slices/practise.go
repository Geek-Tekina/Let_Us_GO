package main

import "fmt"

func main() {
	var arr []string
	arr = append(arr, "Hey")

	fmt.Println(arr)
	arr = append(arr, "Aniket", "here")

	arr2 := []string{"Let's Code guys"}
	arr = append(arr, arr2...)
	fmt.Println(arr)

	length := len(arr)
	capacity := cap(arr)

	fmt.Printf("The length of arr slice is %v and capacity is %v\n", length, capacity)

	slice := make([]int, 5, 10) // make is used to define slice with both initial capacity and the length, 5 is len and 10 is capacity
	fmt.Print(slice)

}
