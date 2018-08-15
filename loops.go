package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	for diff := z; math.Abs(diff) > 1e-10; i++ {
		diff = ((z * z) - x) / (2 * z)
		z -= diff
	}
	return z, i
}

func main() {
	var n float64 = 2
	fmt.Println(Sqrt(n))
	fmt.Println(math.Sqrt(n))
}

