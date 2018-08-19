package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	s := fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot Sqrt negative number: %v", s)
}

func Sqrt(x float64) (interface{}, int) {
	if x < 0 {
		return ErrNegativeSqrt(x), 0
	}
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
	fmt.Println(Sqrt(-n))
}

