package main

import "fmt"

type Array struct {
	length int
	data   []interface{}
}

func (array *Array) GetElement(index int) interface{} {
	return array.data[index]
}

func (array *Array) Push(item interface{}) {
	array.data = append(array.data, item)
	array.length++
}

func (a *Array) GetLength() int {
	return a.length
}

func (a *Array) Pop() error {
	if a.length <= 0 {
		return fmt.Errorf("Error: empty array")
	}

	newLeng := a.GetLength()
	print(newLeng)
	a.data = a.data[:a.GetLength()-1]
	a.length--
	return nil
}

func (a *Array) Insert(index int, item interface{}) {
	// add an "empty" cell at the end of the array
	a.Push(0)
	for i := a.length - 1; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = item
}

func (a *Array) Delete(index int) error {
	if index < 0 || index > a.length-1 {
		return fmt.Errorf("index out of bound")
	}

	for i := index; i < a.length-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.Pop()
	return nil
}

// main function to test it
func main() {
	var array1 Array
	array1.Push("0")
	array1.Push("1")
	array1.Push("8")
	array1.Push("3")
	array1.Push("5")
	array1.Delete(0)
	array1.Pop()
	fmt.Println(array1.data)
}
