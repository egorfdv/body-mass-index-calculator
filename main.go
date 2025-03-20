package main

import (
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	userHeight := 1.8
	userWeight := 100.0
	BMI := userWeight / math.Pow(userHeight, BMIPower)
	fmt.Print(BMI)
}
