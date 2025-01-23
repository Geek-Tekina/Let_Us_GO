package main

import "fmt"

func main() {
	reportCard := make(map[string]int)
	reportCard["Eng"] = 97
	reportCard["Maths"] = 98
	reportCard["EVS"] = 89
	fmt.Println(reportCard)

	for key, value := range reportCard {
		fmt.Println(key, "->", value)
	}
	// to check if any key-value is present or not, as it always returns 0 as default value for any key that is not present
	grades, exists := reportCard["Science"]
	if !exists {
		fmt.Println("The key value pair does not exists", grades, "Not found")
	}
	// deleting a key value pair
	delete(reportCard, "EVS")

	for key, value := range reportCard {
		fmt.Println(key, "->", value)
	}

	//another declaration of map

	var mp = map[string]int{
		"hey":   1,
		"hello": 2,
	}
	fmt.Println(mp)
}
