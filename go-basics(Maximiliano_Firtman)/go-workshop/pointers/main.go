package main

import "fmt"

func birthday(age *int){
	*age++
}

func main() {

	age := 22
	birthday(&age)
	
	fmt.Println(age)

}