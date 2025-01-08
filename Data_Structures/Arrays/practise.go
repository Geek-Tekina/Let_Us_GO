package main

import "fmt"

func main() {
	var arr [3]string

	arr[0] = "Hey"
	arr[1] = "Aniket"
	arr[2] = "here !"

	length := len(arr)
	fmt.Println("Length :", length)

	fmt.Println("The array is : ", arr)
	fmt.Println("Array using for - range")
	for i, element := range arr {
		fmt.Printf("The element at %v is %v \n", i, element)
	}

	arr2 := [3]string{"abc", "pyq", "xyz"} // another way of initialising the array
	fmt.Println(arr2)
}
