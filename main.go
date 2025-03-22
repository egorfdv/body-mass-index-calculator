package main

import (
	"errors"
	"fmt"
	"math"
)

const BMIPower = 2

func main() {
	fmt.Println("__ Body mass index calculator __")
	for {
		var userWeight, userHeight float64 = getUserInput()
		BMI, err := calculateBMI(userWeight, userHeight)
		if err != nil {
			fmt.Println("No parameters specified for calculation")
		}
		outputResult(BMI)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
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

func calculateBMI(userWeight float64, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("noParameters")
	}
	BMI := userWeight / math.Pow(userHeight/100, BMIPower)
	return BMI, nil
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your body mass index: %.0f", BMI)
	fmt.Println(result)
	switch true {
	case BMI < 16:
		fmt.Println("You are severely underweight")
	case BMI < 18.5:
		fmt.Println("You are underweight")
	case BMI < 25:
		fmt.Println("You are of normal weight")
	case BMI < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("You have a degree of obesity")
	}
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to repeat the calculation?(y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	}
	return false
}
