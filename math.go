package main

import (
	"fmt"
	"math"
)

func main() {
	q := 34.0

	fmt.Println(isEven(q))
}

func isEven(x float64) bool {
	if math.Mod(x, 2) == 0 {
		return true
	} else {
		return false
	}

}
