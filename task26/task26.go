package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция, которая проверяет, что все символы в строке уникальные
func areAllCharsUnique(str string) bool {
	// Перевод строки в нижний регистр
	lowercaseStr := strings.ToLower(str)
	// Создание словаря для хранения символов и их значений
	chars := make(map[rune]bool)
	// Прохождение по каждому символу в строке
	for _, c := range lowercaseStr {
		// Проверка наличия символа в качестве ключа в map
		if chars[c] {
			// Если есть, значит строка не уникальна
			return false
		}
		// Добавление его в map с любым значением
		chars[c] = true
	}
	// Если повторения не были обнаружены, значит строка уникальна
	return true
}

func main() {
	fmt.Println("Введите строки для проверки уникальности символов (для завершения введите 'exit'):")

	// Создание сканнера для чтения ввода пользователя
	scanner := bufio.NewScanner(os.Stdin)
	// Бесконечный цикл для ввода строк
	for scanner.Scan() {
		// Считывание введенной строки
		in := scanner.Text()
		// Проверка на завершение программы
		if in == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		// Вызов функции для проверки уникальности символов
		res := areAllCharsUnique(in)
		// Вывод результата проверки
		fmt.Printf("%s — %v\n", in, res)
	}

	// Проверка на ошибки сканера
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка ввода:", err)
	}
}
