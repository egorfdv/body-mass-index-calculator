package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __ ")
	for {
		userHeight, userWeight := getUserInput()
		IMT, err := calculateIMT(userWeight, userHeight)
		if err != nil {
			fmt.Println("Не заданы параметры для расчета")
			continue
		}
		outputResult(IMT)
		menuForUser := menuContinueOrExit()
		if !menuForUser {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
	switch true {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У вас нормальный вес")
	case IMT < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userWeight float64, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("no params")
	}
	const IMTPower = 2
	IMT := userWeight / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func menuContinueOrExit() bool {
	var menuVariants string
	fmt.Print("Вы хотите повторить расчет?(y/n): ")
	fmt.Scan(&menuVariants)
	if menuVariants == "y" || menuVariants == "Y" {
		return true
	}
	return false
}

// Trash
// if IMT < 16 {
// 	fmt.Println("У вас сильный дефицит массы тела")
// } else if IMT < 18.5 {
// 	fmt.Println("У вас дефицит массы тела")
// } else if IMT < 25 {
// 	fmt.Println("У вас нормальный вес")
// } else if IMT < 30 {
// 	fmt.Println("У вас избыточный вес")
// } else {
// 	fmt.Println("У вас степень ожирения")
// }

// i := 0
// for i < 10 {
// 	fmt.Printf("%d\n", i)
// 	i++
// }

// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		continue
// 	}
// 	fmt.Printf("%d\n", i)
// }
