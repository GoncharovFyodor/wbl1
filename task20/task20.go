package main

import (
	"fmt"
	"strings"
)

// Переворачивание слов в строке с использованием среза строк
func reverseWordsSlice(s string) string {
	// Разбиение строки на слова
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем слова местами
	}

	return strings.Join(words, " ") // Объединяем слова в строку с пробелами
}

// Переворачивание строки с использованием string.Builder
func reverseWordsBuilder(s string) string {
	// Разбиение строки на слова
	words := strings.Fields(s)

	var reversed strings.Builder

	// Построение строки с обратным порядком символов
	for i := len(words) - 1; i >= 0; i-- {
		reversed.WriteString(words[i]) // Добавление слова в strings.Builder
		if i > 0 {
			reversed.WriteRune(' ') // Добавление пробела, кроме последнего слова
		}
	}

	return reversed.String() //Получение строки из strings.Builder
}

// Переворачивание строки с использованием байтового массива
func reverseWordsBytes(s string) string {
	// Разбиение строки на слова
	words := strings.Fields(s)

	//Создание байтового среза
	var reversed []byte

	// Запись байт в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, []byte(words[i])...)
		if i > 0 {
			reversed = append(reversed, ' ')
		}
	}

	return string(reversed) //Преобразование байтового среза в строку
}

func main() {
	input := "Каждый суслик - агроном"
	fmt.Println("Тестовая строка:", input)

	// Переворот слов с использованием среза рун
	reversedSlice := reverseWordsSlice(input)
	fmt.Println("Результат с использованием среза слов:", reversedSlice)

	// Переворот слов с использованием strings.Builder
	reversedBuilder := reverseWordsBuilder(input)
	fmt.Println("Результат с использованием strings.Builder:", reversedBuilder)

	// Переворот слов с использованием байтового среза
	reversedBytes := reverseWordsBytes(input)
	fmt.Println("Результат с использованием байтового среза:", reversedBytes)
}
