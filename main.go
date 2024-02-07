package main

import (
	"fmt"
	"math"
)

func main() {
	// create n, dynamic variable
	var n int = 64

	for i := 1; i <= n; i++ {
		// convert n to be float64, so it can be used in math.Sqrt() & math.Cbrt()
		var x float64 = float64(i)
		powerFloat2 := math.Sqrt(x)
		powerFloat3 := math.Cbrt(x)

		var powerInt2 int = int(powerFloat2)
		var powerInt3 int = int(powerFloat3)

		// compare integer to float to decide SquareCube, Square, Cube, or not
		if float64(powerInt2) == powerFloat2 && float64(powerInt3) == powerFloat3 {
			fmt.Println("SquareCube")
			// fmt.Println("SquareCube", float64(powerInt2), powerFloat2, float64(powerInt3), powerFloat3)
		} else if float64(powerInt2) == powerFloat2 {
			fmt.Println("Square")
			// fmt.Println("Square", float64(powerInt2), powerFloat2, float64(powerInt3), powerFloat3)
		} else if float64(powerInt3) == powerFloat3 {
			fmt.Println("Cube")
			// fmt.Println("Cube", float64(powerInt2), powerFloat2, float64(powerInt3), powerFloat3)
		} else {
			fmt.Println(i)
			// fmt.Println(i, float64(powerInt2), powerFloat2, float64(powerInt3), powerFloat3)
		}
	}
}
