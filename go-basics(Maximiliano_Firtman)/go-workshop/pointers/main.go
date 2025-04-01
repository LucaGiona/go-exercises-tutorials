package main

import "fmt"

func birthday(age *int){
	*age++
}

func main() {

	defer fmt.Println("Good")
	defer fmt.Println("Bye")
	age := 22
	birthday(&age)
	
	fmt.Println(age)

}