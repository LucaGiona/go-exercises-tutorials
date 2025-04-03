package main

import "fmt"

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
	}
}

func main() {
	printMessage("Go is great")

}
