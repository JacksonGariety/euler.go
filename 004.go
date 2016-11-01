package main

import . "fmt"
import . "math"

func rev(n int) int {
	num := n
	rev := 0
	for num > 0 {
		rev = rev * 10 + (num % 10)
		num = num / 10
	}
	return rev
}

func largestPalindromeProduct(ds int) int {
	n := int(Pow(10, float64(ds)))
	r := -1
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			p := j * i
			if (p == rev(p)) { r = p}
		} 
	}
	return r
}

func main() {
	Println(largestPalindromeProduct(2))
}
