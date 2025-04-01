package data

import "fmt"

var Countries [10]string //value is nill
var Slice [] int
var Codes map[int] bool

func init(){
	Countries[0] = "Switzerland"
	Countries[1] = "Germany"
	Countries[2] = "Italy"
	Countries[8] = "Argentinia"

	qty := len(Countries)

	fmt.Println("Countries saved", qty)
}