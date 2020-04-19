package main

import (
	"../../ds/bst"

	"fmt"
)

/*
For array representation: [4,2,7,1,3,5,9]

                 4
               /   \
   (2*n + 1)  2     7 (2*n +2)
             / \   / \
            1   3 5   9

*/

func main() {
	// using actual BST
	var tree = new(bst.BST)
	tree.Add(4)
	tree.Add(2)
	tree.Add(7)
	tree.Add(1)
	tree.Add(5)
	tree.Add(3)
	tree.Add(9)
	fmt.Println(tree.BFSearch(1))

	// using array representation of tree
	var arr = []int{4, 2, 7, 1, 3, 5, 9}
	fmt.Println(BFS(arr, 1))
}

//BFS using array linear traversal as (2*n + 1) & (2*n + 2) can be achieved this way also(easier too :) ) .
func BFS(arr []int, key int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == key {
			return true
		}
	}
	return false
}
