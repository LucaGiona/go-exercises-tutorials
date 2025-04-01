package calculator

func SumFromAtoB (a,b int) int {
	if a > b {
		return 0
	}
	return a + SumFromAtoB(a +1, b)
}

func MultiplyFromAtoB (a,b int) int{
	if a > b {
		return 1
	}

	return a * MultiplyFromAtoB(a+1,b)
}


//muss oben dann noch eingefügt werden
func ProcessFromAtoB (a,b, initValue int, fn func(int, int)int ) int {
	if a > b{
		return initValue
	}
	return fn(a, ProcessFromAtoB(a+1, b, initValue, fn))

}