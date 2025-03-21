package main

import (
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	var userHeight float64
	var userWeight float64
	fmt.Println("__ Body mass index calculator __")
	fmt.Print("Enter your height (centimeters): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight (kilograms): ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	fmt.Printf("Your body mass index: %.0f", BMI)
}
