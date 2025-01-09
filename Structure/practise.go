package main

import "fmt"

type Student struct { // declaration of struct
	name     string
	age      int
	subjects []string
}

// a method for struct, this is using value reciever i.e. "stu Student"
func (stu Student) DisplayDetails() {
	fmt.Println("Name :", stu.name)
	fmt.Println("Age :", stu.age)
	fmt.Print("Subjects :")
	for _, el := range stu.subjects {
		fmt.Print(el, " ")
	}
}
func main() {
	var stu1 Student
	stu1.name = "Tekina"
	stu1.age = 21
	stu1.subjects = []string{"English", "Maths", "Science"}
	stu1.DisplayDetails()
}
