package main

import (
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	fmt.Println("__ Body mass index calculator __")
	var userWeight, userHeight float64 = getUserInput()
	BMI := calculateBMI(userWeight, userHeight)
	outputResult(BMI)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter your height (centimeters): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight (kilograms): ")
	fmt.Scan(&userWeight)
	return userWeight, userHeight
}

func calculateBMI(userWeight float64, userHeight float64) float64 {
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your body mass index: %.0f", BMI)
	fmt.Print(result)
}
