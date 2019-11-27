package main

import "fmt"

func main() {
	// Проинициализированная карта возрастов пациентов
	patients := map[string]int{
		"oleg":     76,
		"pavel":    81,
		"grisha":   99,
		"sasha":    23,
		"andrey":   33,
		"tanya":    45,
		"zaya":     22,
		"tolik":    18,
		"kostya":   99,
		"shantsun": 101,
		"lesha":    11,
		"sam":      4,
		"valya":    77,
	}
	// флаг о первоначальной инициализации минимального возраста
	flagMinInit := false

	// минимальное значение
	var min int

	// имя (ключ) самого молодого пациента
	var youngestPatientName string

	// начинаем обход пациентов
	for name, age := range patients {
		// устанавливаем минимальное значение
		if !flagMinInit {
			flagMinInit = true
			min = age
			youngestPatientName = name
		}

		// выявляем самого молодого из пациентов
		if age < min {
			min = age
			youngestPatientName = name
		}
	}

	fmt.Println("Youngest Patient is ", youngestPatientName, " aged ", patients[youngestPatientName])
}
