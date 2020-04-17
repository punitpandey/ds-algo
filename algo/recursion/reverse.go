package main

import "fmt"

func main() {
	reverse("this is recursion")
	fmt.Printf("\nreverseWithOutput: %v \n", reverseWithOutput("this is recursion"))

}

func reverseWithOutput(str string) string {
	if len(str) == 1 {
		return str
	}
	return reverseWithOutput(str[1:]) + str[0:1]
}

func reverse(str string) {
	if len(str) == 1 {
		fmt.Print(str)
		return
	}
	reverse(str[1:])
	fmt.Print(str[0:1])
}
