package main

import (
	"fmt"
	"math"
)

func main() {
	//Температуры
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//Словарь для группировки температур по диапазонам
	groupedTemperatures := make(map[int][]float64)

	for _, t := range temperatures {
		//Определение номера группы для температуры
		var group int
		if t < 0 {
			group = int(math.Ceil(t/10.0)) * 10
		} else {
			group = int(math.Floor(t/10.0)) * 10
		}

		//Добавление температуры в соответствующую группу
		groupedTemperatures[group] = append(groupedTemperatures[group], t)
	}

	//Вывод групп температур
	for group, temps := range groupedTemperatures {
		fmt.Printf("%d:%v\n", group, temps)
	}
}
