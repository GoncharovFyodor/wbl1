package main

import (
	"fmt"
	"log"
)

// Установка i-го бита числа n в 1 или 0
func setBit(n int64, i uint, b int) int64 {
	//Создание маски с установкой бита в позиции i (используя побитовый сдвиг)
	mask := int64(1) << i

	if b == 0 {
		//Если b = 0, то используется побитовое И с инвертированной маской
		return n & ^mask
	}
	//Если b = 1, то используется побитовое ИЛИ с маской
	return n | mask
}

func main() {
	var n int64
	fmt.Print("Введите число (int64): ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal()
	}

	var i uint
	fmt.Print("Введите позицию бита (0-63): ")
	_, err = fmt.Scan(&i)
	if err != nil {
		log.Fatal()
	}

	//Проверка корректности введенной позиции бита
	if i < 0 || i > 63 {
		fmt.Println("Неверная позиция бита")
		return
	}

	var b int
	fmt.Print("Введите значение бита (0 или 1): ")
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal()
	}

	//Проверка корректности введенного значения бита
	if b != 0 && b != 1 {
		fmt.Println("Неверное значение бита")
		return
	}

	numWithOps := setBit(n, i, b)
	fmt.Println("Число с установленным битом:", numWithOps)
}
