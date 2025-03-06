package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.8
	userWeight := 100.0
	var IMT = userWeight / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}

// var userHeight float64 = 1.8
// var userWeight float64 = 100
// var IMT = float64(userWeight) / (userHeight)
// var userKg float64
// userKg = 100
// var userHeight, userKg float64 = 1.8, 100
// userHeight, userKg := 1.8, 100
