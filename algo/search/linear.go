package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 8, 3, 9, 7}
	fmt.Println(LinearSearch(arr, 6))
}

func LinearSearch(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return true
		}
	}
	return false
}
