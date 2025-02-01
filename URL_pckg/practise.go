package main

import (
	"fmt"
	"net/url"
)

func main() {
	link := "https://www.funkyflavors.com/search?type=fabulous&category=festive&intensity=ferocious"
	parsedUrl, _ := url.Parse(link)

	fmt.Printf("Now type of url is : %T\n", parsedUrl)

	fmt.Println("Host :", parsedUrl.Host)
	fmt.Println("Path :", parsedUrl.Path)
	fmt.Println("Raw Query :", parsedUrl.RawQuery)

	// extracting params IMP
	params := parsedUrl.Query()

	typeParam := params.Get("type")
	fmt.Println("Type param :", typeParam)

	// adding
	par := url.Values{}
	par.Add("id", "123")
	fullURL := fmt.Sprintf("%s?%s", link, par.Encode())

	fmt.Println("New URL ->>>>> ", fullURL)

}
