package main

import (
	"fmt"
	"time"
)

func SummationLoop(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func SummationFormula(n int) int {
	return n * (n + 1) / 2
}

func main() {
	n := 1000000

	start := time.Now()
	loopResult := SummationLoop(n)
	loopDuration := time.Since(start)

	start = time.Now()
	formulaResult := SummationFormula(n)
	formulaDuration := time.Since(start)

	fmt.Printf("Loop: %d (tempo: %v)\n", loopResult, loopDuration)
	fmt.Printf("Fórmula: %d (tempo: %v)\n", formulaResult, formulaDuration)
}
