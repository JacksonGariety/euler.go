package main

import "fmt"

func multiplesOf3And5(n int) int {
	var sum int = 0
	for i := 0; i < n; i++ {
		if i % 3 == 0 || i % 5 == 0 { sum += i }
	}

	return sum
}

func main() {
	fmt.Println(multiplesOf3And5(1000))
}
