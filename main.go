package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("__ Калькулятор индекса массы тела __ ")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

// var userHeight float64 = 1.8
// var userWeight float64 = 100
// var IMT = float64(userWeight) / (userHeight)
// var userKg float64
// userKg = 100
// var userHeight, userKg float64 = 1.8, 100
// userHeight, userKg := 1.8, 100
// fmt.Print("__ Калькулятор индекса массы тела __ \n")
// fmt.Print("Ваш индекс", IMT)
// fmt.Printf("Ваш индекс массы тела: %v", IMT)
// fmt.Print(`__ Калькулятор индекса массы тела __
// Введите свой рост в сантиметрах: %.1f`, 0,1)
