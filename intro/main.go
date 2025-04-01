package main

import (
	"fmt"
	random "math/rand"

	"github.com/thenativeweb/main/calculator"
	"github.com/thenativeweb/main/composite"
)

func main() {

	fmt.Println("Hallo Welt", random.Intn(10))
	fmt.Println("23 + 43 = ", calculator.Add(23,42))
	fmt.Println(calculator.Divide(17,3))	

	fmt.Println(calculator.Sum(1, 10))
	fmt.Println(calculator.SumUntil(10))

	fmt.Println(calculator.IsSquareNumber(25))

	fmt.Println(composite.Add(23, 42))
}
