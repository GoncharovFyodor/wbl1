package main

import (
	"fmt"
	"strings"
)

// Переворачивание строки с использованием среза рун
func reverseWithSlice(s string) string {
	// Преобразование строки в срез рун
	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем руны местами
	}

	return string(runes) //Преобразование среза рун обратно в строку
}

// Переворачивание строки с использованием string.Builder
func reverseWithStringsBuilder(s string) string {
	var reversed strings.Builder

	// Преобразование строки в срез рун
	runes := []rune(s)
	length := len(runes)

	// Построение строки с обратным порядком символов
	for i := length - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i]) // Запись руны в обратном порядке
	}

	return reversed.String() //Получение строки из strings.Builder
}

// Переворачивание строки с использованием байтового массива
func reverseWithBytes(s string) string {
	// Преобразование строки в срез рун
	runes := []rune(s)
	length := len(runes)

	//Создание байтового среза
	reversed := make([]byte, 0, len(s))

	// Запись байт в обратном порядке
	for i := length - 1; i >= 0; i-- {
		reversed = append(reversed, []byte(string(runes[i]))...) // Преобразование руны в байты и добавление к срезу
	}

	return string(reversed) //Преобразование байтового среза в строку
}

func main() {
	s := "абырвалг"
	fmt.Println("Исходная строка:", s)
	fmt.Println("Перевернутая строка (с использованием среза рун):", reverseWithSlice(s))
	fmt.Println("Исходная строка (с использованием strings.Builder):", reverseWithStringsBuilder(s))
	fmt.Println("Исходная строка (с использованием байтового массива):", reverseWithBytes(s))
}
