package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 7, 8, 9}
	fmt.Println(BinarySearch(arr, 6))
}

//BinarySearch expects sorted arr
func BinarySearch(arr []int, key int) bool {
	if len(arr) == 1 {
		return arr[0] == key
	}
	middle := len(arr) / 2
	if arr[middle] == key {
		return true
	} else if arr[middle] < key {
		return BinarySearch(arr[middle+1:], key)
	} else {
		return BinarySearch(arr[:middle], key)
	}
}
