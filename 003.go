package main

import . "fmt"
import . "math"

func rho(n int) int {
	a, c, r := n, 2, 0
	for Pow(float64(c), 2) <= float64(a) {
		if a % c == 0 {
			a = a / c
			r = c
		} else {
			if c == 2 { c = 3 } else { c = c + 2 }
		}
	}

	if (a > r) { r = a }
	return r
}
 
func main() {
	Println(rho(600851475143))
} 
