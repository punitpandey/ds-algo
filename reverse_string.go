package main

import (
	"fmt"
	"strings"
)

func ReverseString(str string) string {
	if len(str) <= 0 {
		return "wrong input: empty string"
	}
	strArray := strings.Split(str, "")
	var resersedArray []string
	for i := len(strArray) - 1; i >= 0; i-- {
		resersedArray = append(resersedArray, strArray[i])
	}
	reversedString := strings.Join(resersedArray, "")
	return reversedString
}

func main() {
	string1 := "-"
	fmt.Print(ReverseString(string1))
}
