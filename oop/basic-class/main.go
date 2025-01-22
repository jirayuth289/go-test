package oop

import (
	"fmt"
)

func TestCreateBasicClass() {
	// Create an array of Student
	var students [3]Student
	students[0] = Student{FirstName: "Mark", LastName: "Zuckerberg", Weight: 54, Height: 180, Grade: "A"}
	students[1] = Student{FirstName: "Dan", LastName: "Test", Weight: 54, Height: 180, Grade: "A"}
	students[2] = Student{FirstName: "John", LastName: "Test", Weight: 54, Height: 180, Grade: "A"}
	for i := 0; i < len(students); i++ {
		fmt.Printf("Student[%d] = %s\n", i, students[i].FullName())
	}

	// Create Map of Student
	myStudent := make(map[string]Student)
	myStudent["mark"] = Student{FirstName: "Mark", Weight: 54, Height: 180, Grade: "A"}
	myStudent["dan"] = Student{FirstName: "Dan", Weight: 54, Height: 180, Grade: "A"}
	myStudent["john"] = Student{FirstName: "John", Weight: 54, Height: 180, Grade: "A"}
	fmt.Println("Mark", myStudent["mark"])
	// Delete a key-value pair
	delete(myStudent, "mark")

	for key, value := range students {
		fmt.Println(key, "->", value)
	}

	val, ok := myStudent["wadee"]
	if ok {
		fmt.Println("Wadee's value: ", val)
	} else {
		fmt.Println("Wadee not found in map")
	}
}
