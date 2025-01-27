package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Hey, Hello, Hi, Yeah"
	// split function
	parts := strings.Split(data, ",")
	fmt.Println("Parts :")
	for _, part := range parts {
		fmt.Println(part)
	}

	// Count fnc
	elements := "Goku Naruto Goku Nobita goku"
	cnt := strings.Count(elements, "Goku")
	fmt.Println("Count of Goku :", cnt)

	// TrimSpace() fnc - only removes spaces from starting and end
	payload := "              Hey this is      payload    from response of     an API                .          "
	trimmedPayload := strings.TrimSpace(payload)
	fmt.Println("Paylod >>>>>>>>", trimmedPayload)

}
