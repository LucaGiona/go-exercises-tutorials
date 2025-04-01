package calculator

const Pi float64 = 3.141592654


func Add (left int, right int) int {
	sum := left + right
	return sum
}

func Divide (left, right int) (quotient int, remainder int){
	quotient = left / right 
	remainder = left % right

	return
}