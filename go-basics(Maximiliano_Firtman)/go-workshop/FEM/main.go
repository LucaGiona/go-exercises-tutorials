package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {
	max := data.Instructor{Id: 3, LastName: "Gionii"}
	max.FirstName = "Maximilliano"

	kyle := data.NewInstructor("Kyle", "Simpson")

	goCourse := data.Course{Id: 2, Name: "Go Fundametals", Instructor: max}

	//print(max.Print())
	

	fmt.Printf("%v", goCourse)
	print(kyle.Print())
}
