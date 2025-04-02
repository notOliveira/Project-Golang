package main

import "fmt"

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", 1.1, 1.1)
	fmt.Printf("Type: %T - Value: %v\n", "string", "string")
	fmt.Printf("Type: %T - Value: %v\n", []int{1, 2, 3}, []int{1, 2, 3})
	fmt.Printf("Type: %T - Value: %v\n", map[string]int{"key": 1}, map[string]int{"key": 1})
	fmt.Printf("Type: %T - Value: %v\n", struct{ Name string }{"struct"}, struct{ Name string }{"struct"})
	fmt.Printf("Type: %T - Value: %v\n", nil, nil)
	fmt.Printf("Type: %T - Value: %v\n", make(chan int), make(chan int))
	fmt.Printf("Type: %T - Value: %v\n", []interface{}{1, 2, 3}, []interface{}{1, 2, 3})
	fmt.Printf("Type: %T - Value: %v\n", map[interface{}]interface{}{"key": 1}, map[interface{}]interface{}{"key": 1})
	fmt.Printf("Type: %T - Value: %v\n", struct{ Name string }{"struct"}, struct{ Name string }{"struct"})
}
