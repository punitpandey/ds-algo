package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 8, 3, 9, 7}
	fmt.Println(BFS(arr, 6))
}

/*
                 4
               /   \
   (2*n + 1)  2     5 (2*n +2)
             / \   / \
            8   3 9   7
*/
func BFS(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return true
		}
	}
	return false
}
