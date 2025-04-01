package main

import "fmt"

func main(){

	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("=================")
	fmt.Println("Which operation you want to perform? (add, sub, multi, divide) ")

	fmt.Scanf("%s", &operation)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)
	fmt.Println("The result is:")
	switch operation{
	case "add":
		fmt.Println(number1 + number2)
	case "sub":
		fmt.Println(number1 - number2)
	case "multi":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)

	}

}