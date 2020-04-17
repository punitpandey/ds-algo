package main

import "fmt"

func main() {
	fmt.Printf("factorialWithRecursion: %v \n", factorialWithRecursion(5))
	fmt.Printf("factorialWithLoop: %v \n", factorialWithLoop(5))
}

//factorialWithRecursion returns 1 for numbers less than 2 or else the factorial using recursion
func factorialWithRecursion(num int) int {
	if num < 2 {
		return 1
	}
	return num * factorialWithRecursion(num-1)
}

//factorialWithLoop returns 1 for numbers less than 2 or else the factorial using loop
func factorialWithLoop(num int) int {
	var factorial = 1
	for ; num > 1; num-- {
		factorial *= num
	}
	return factorial
}
