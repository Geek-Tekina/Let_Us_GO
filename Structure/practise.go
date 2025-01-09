package main

import "fmt"

type Student struct { // declaration of struct
	name     string
	age      int
	subjects []string
}

// This is a constructor method, which is return a newly created instance of Student.
func NewStudent(name string, age int, sub []string) Student {
	return Student{
		name:     name,
		age:      age,
		subjects: sub,
	}
}

// a method for struct, this is using value reciever i.e. "stu Student"
func (stu Student) DisplayDetails() {
	fmt.Println("Name :", stu.name)
	fmt.Println("Age :", stu.age)
	fmt.Print("Subjects :")
	for _, el := range stu.subjects {
		fmt.Print(el, " ")
	}
	fmt.Println("\n----------------------------------------------")
}
func main() {
	var stu1 Student
	stu1.name = "Tekina"
	stu1.age = 21
	stu1.subjects = []string{"English", "Maths", "Science"}
	// this is one way of declaration
	stu1.DisplayDetails()

	// short hand for structure  is
	stu2 := Student{name: "Goku", age: 23, subjects: []string{"Geography", "Ecology"}}
	stu2.DisplayDetails()

	// In the above two methods, we are directly giving the values to structure instance. It is advised to have a constructor method to do this. It starts with "New" prefix.

	// Declaration using constructor

	stu3 := NewStudent("Pyhto", 25, []string{"Fighting", "Running"})
	stu3.DisplayDetails()
}
