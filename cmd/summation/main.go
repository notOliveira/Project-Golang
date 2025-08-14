package main

import (
	"fmt"
	"strconv"
)

func Summation(n int) string {
	txtOutput := "("
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i

		txtOutput += strconv.Itoa(i)
		if i < n {
			txtOutput += " + "
		}
	}

	txtOutput += ")"

	return fmt.Sprintf("%v -> %v %v", n, sum, txtOutput)
}

func main() {
	fmt.Println(Summation(5))
	fmt.Println(Summation(8))
	fmt.Println(Summation(12))
}