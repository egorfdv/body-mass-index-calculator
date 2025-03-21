package main

import (
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	var userHeight float64
	var userWeight float64
	fmt.Print("__ Body mass index calculator __\n")
	fmt.Print("Enter your height (centimeters): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight (kilograms): ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	fmt.Print("Your body mass index: ")
	fmt.Print(BMI)
}
