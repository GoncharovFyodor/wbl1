package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Введите значение: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	var value interface{}

	// Попытка преобразования в int
	if intValue, err := strconv.Atoi(input); err == nil {
		value = intValue
	} else {
		// Попытка преобразования в bool
		if boolValue, err := strconv.ParseBool(input); err == nil {
			value = boolValue
		} else {
			// Попытка создания chan
			if makeChan, ok := makeChanFromString(input); ok {
				value = makeChan
			} else {
				value = input
			}
		}
	}
	fmt.Printf("Введенное значение: %v\n", value)
	fmt.Printf("Тип введенного значения: %T\n", value)
}

// Вспомогательная функция, создающая chan из строки input
func makeChanFromString(input string) (chan int, bool) {
	if input == "chan" {
		return make(chan int), true
	}
	return nil, false
}
