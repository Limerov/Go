package main

import "fmt"

func main() {
	// Общее количество фруктов
	fruit := map[string]int{
		"Apple":             76,
		"orange":            81,
		"banana":            99,
		"yellow apple":      23,
		"Green appel":       33,
		"red appel":         45,
		"mango":             22,
		"melon":             18,
		"green mellon":      99,
		"beg mellon":        101,
		"tangerines":        11,
		"Chaina tangerines": 4,
		"small tangerines":  77,
	}
	// Минимальный остаток
	MinimumBalans := false

	// минимальное значение
	var min int

	// Наименование (ключ) минимального остатка фрукта
	var MinimumAmountFruit string

	// Приступик к взвешиванию
	for title, weight := range fruit {
		// устанавливаем минимальное значение
		if !MinimumBalans {
			MinimumBalans = true
			min = weight
			MinimumAmountFruit = title
		}

		// Определяем минимальный остаток
		if weight < min {
			min = weight
			MinimumAmountFruit = title
		}
	}

	fmt.Println("Minimum amaunt ", MinimumBalans, " title ", fruit[MinimumAmountFruit])
}
