package main

import (
	"fmt"
)

// Метод 1: Использование кортежей
func method1(a int, b int) (int, int) {
	fmt.Println("Метод 1: Использование кортежей")
	b, a = a, b // меняем местами значения с помощью присваивания кортежей
	return a, b
}

// Метод 2: Использование XOR
func method2(a int, b int) (int, int) {
	fmt.Println("Метод 2: Использование XOR")
	a = a ^ b // применяем операцию XOR к a и b и сохраняем результат в a
	b = a ^ b // применяем операцию XOR к a и b и сохраняем результат в b
	a = a ^ b // применяем операцию XOR к a и b и сохраняем результат в a
	return a, b
}

// Метод 3: Использование арифметических операций
func method3(a int, b int) (int, int) {
	fmt.Println("Метод 3: Использование арифметических операций")
	a = a + b // складываем a и b и сохраняем результат в a
	b = a - b // вычитаем b из a и сохраняем результат в b
	a = a - b // вычитаем b из a и сохраняем результат в a
	return a, b
}

// Метод 4: Использование указателей
func method4(a int, b int) (int, int) {
	fmt.Println("Метод 4: Использование указателей")
	swapByPointers(&a, &b)
	return a, b
}

// Функция для обмена значениями через указатели
func swapByPointers(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 50
	b := -20
	fmt.Println("До обмена:", a, b)
	fmt.Println(method1(a, b))
	fmt.Println()
	fmt.Println(method2(a, b))
	fmt.Println()
	fmt.Println(method3(a, b))
	fmt.Println()
	fmt.Println(method4(a, b))
	fmt.Println()
}
