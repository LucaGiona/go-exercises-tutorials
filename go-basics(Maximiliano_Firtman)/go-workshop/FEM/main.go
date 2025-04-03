package main

import (
	"fmt"

	"frontendmasters.com/go/server/data"
)

func main() {

	max := data.Instructor{Id: 3, LastName: "Gionii"}
	max.FirstName = "Maximilliano"

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: max}

	//print(max.Print())

	fmt.Printf("%v", goCourse)

	swftWS := data.NewWorkshop("Swift with iOS", max)

	fmt.Printf("%v", swftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swftWS
	for _, course := range courses {
		fmt.Println(course)
	}
}
