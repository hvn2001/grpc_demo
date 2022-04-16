package main

import (
	"fmt"
	"math"
)

type MyFloat34 float64

func (f MyFloat34) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := MyFloat34(-math.Sqrt2)
	fmt.Println(f.Abs())
}
