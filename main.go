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
	if BMI < 16 {
		fmt.Println("You are severely underweight")
	} else if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI < 25 {
		fmt.Println("You are of normal weight")
	} else if BMI < 30 {
		fmt.Println("You are overweight")
	} else {
		fmt.Println("You have a degree of obesity")
	}
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
	fmt.Println(result)
}
