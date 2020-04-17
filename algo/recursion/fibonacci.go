package main

import "fmt"

func main() {
	fmt.Printf("fibonacciWithLoop: %v \n", fibonacciWithLoop(9))
	fmt.Printf("fibonacciWithRecursion: %v \n", fibonacciWithRecursion(9))
}

func fibonacciWithRecursion(index int) int {
	if index < 2 {
		return index
	}
	return fibonacciWithRecursion(index-1) + fibonacciWithRecursion(index-2)
}

func fibonacciWithLoop(index int) int {
	first, second := 0, 1
	if index < 2 {
		return index
	}
	for i := 2; i < index+1; i++ {
		first, second = second, first+second
	}
	return second
}
