package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time :", time.Now()) //this gives default format
	currentTime := time.Now()

	formatted := currentTime.Format("02-01-2006, 15:04:05") // this formats the time, 02-01-2006 -> Go was live on this date
	fmt.Println("Formatted time :", formatted)
	// Format is fixed
	timeFormat := currentTime.Format("02/01/2005, 3:04 PM")
	fmt.Println(timeFormat)

}
