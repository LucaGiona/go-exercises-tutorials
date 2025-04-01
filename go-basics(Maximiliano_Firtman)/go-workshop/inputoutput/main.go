package main

import (
	"fmt"
	
)

//global
var url = "https://google.com"

func init() {
	fmt.Println("A")
}
func init() {
	fmt.Println("B")
}

func calculateTax(price float32)(float32, float32) {

	return price*0.09, price*0.02
}
func calculateTaxWithName(price float32)(stateTax float32, cityTax float32) {

	return price*0.09, price*0.02
}


func main(){
	//function-scoped variables

	 message  := "Hello from go \n"
	 price := 34.5


	print(message, price, url)
	print("\n")
	printData()

	cityTax, _ := calculateTax(100)

	fmt.Println("City: ",cityTax,)

	age:=22
	fmt.Println(age)

	fmt.Println(calculateTaxWithName(90.9))
}	

