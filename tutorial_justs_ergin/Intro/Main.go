package main

import (
	"fmt"
)

func main(){
	var age int
	var name string

	fmt.Println("Bitte gebe dein alter ein: ")
	fmt.Scan(&age)
	
	fmt.Println("Gib bitte deinen Namen an: ")
	fmt.Scan(&name)
	fmt.Printf(("Hallo Welt! Hallo, du bist %v Jahre alt und du heisst %v \n"),  age, name)

	if age < 18 {
		var missing = 18 - age
		var year = "Jahre"
		if missing <= 1{
			year = "Jahr"
		
		}
		fmt.Printf("Du bist nicht volljährig. Du brauchst %v %v, bis du volljährig bist \n", missing, year)
	}else {
		fmt.Println("Du bist volljährig")
	}

}

