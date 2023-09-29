package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)           // Огромная строка
	justStringByteSlice := make([]byte, 100) // Выделение памяти для среза длиной в 100 байт
	copy(justStringByteSlice, v[:100])       // Копирование первых 100 байт огромной строки в срез байтов
	justString = string(justStringByteSlice) // Преобразование среза байт обратно в строку
}

func main() {
	someFunc()

	//Вывод строки
	fmt.Println(justString)
}

// Функция для создания огромной строки
func createHugeString(size int) string {
	// Создание строки, повторяющейся множество раз, чтобы получить огромный размер
	return strings.Repeat("A", size)
}

// Фрагмент кода на языке Go, приведенный в тексте задания, может привести к утечке памяти из-за использования
// среза v[:100] для присвоения значения переменной justString. При этом ссылка на огромную строку в памяти
// сохраняется, и даже если она больше не используется, сборщик мусора не освободит её,
// так как ссылка на огромную строку сохранена в justString.
