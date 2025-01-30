/*
Web req are HTTP request made to a web server. These are used to get or post data over the internet. In GO, net/http package is used to hit web requests.
*/

// This is an illustration where a simple get request is hit and struct is populated using json.Unmarshal
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

func main() {
	// this is simple for making open api calls
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close() // this is resource management and avoiding resource leaks, here network connection is closed.

	fmt.Printf("type : response which we get from API : %T\n", response) // pointer to http.Response

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Printf("type : response body which we get from API : %T\n", body)

	// *************************************** IMPORTANT
	//  data := []uint8{72, 101, 108, 108, 111}

	// // Converting the uint8 slice to a string
	// str := string(data)

	// fmt.Println(str) // Output: Hello
	// ***************************************

	fmt.Println(string(body)) // this converts the uint8 slice into string
	jsonBodyString := string(body)
	// fmt.Printf("%T", string(body))

	fmt.Println("API call excecuted buddy !")
	fmt.Println("Now we will try to use Unmarshal to store this into our struct !!")

	var user1 User
	errr := json.Unmarshal([]byte(jsonBodyString), &user1) // changes JSON data into go data strcutures (struct)
	if errr != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
	fmt.Printf("%+v\n", user1) //  %v outputs the struct in the default format

}
