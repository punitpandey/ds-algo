package main

import "fmt"

var calculation int

func main() {
	fibO := fibonacci()
	fmt.Println(fibO(10))
	fmt.Println(calculation)
}

// fibonacci function uses dynamic programming principle
func fibonacci() func(n int) int {
	var fib func(int) int
	var cache = map[int]int{}
	fib = func(n int) int {
		calculation++
		if result, ok := cache[n]; ok {
			return result
		}
		if n < 2 {
			return n
		}
		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}
	return fib
}

//fib function is simple implementation
func fib(n int) int {
	calculation++
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
