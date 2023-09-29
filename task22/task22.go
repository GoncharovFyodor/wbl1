package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создание нового объекта для чтения ввода пользователя
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите два числа (разделите пробелом):")
		// Чтение строки из ввода.
		in, _ := reader.ReadString('\n')
		// Удаление символа переноса строки из конца строки
		in = strings.Trim(in, "\r\n")
		// Разделение введенной строки на два числа
		nums := strings.Split(in, " ")
		if len(nums) != 2 {
			fmt.Println("Пожалуйста, введите два числа")
			continue // Продолжение цикла
		}

		// Создание объектов big.Int для хранения больших чисел
		var a, b big.Int
		// Преобразование введенных строк в большие числа
		_, errA := a.SetString(nums[0], 10)
		_, errB := b.SetString(nums[1], 10)
		if errA != true || errB != true {
			fmt.Println("Ошибка ввода чисел. Пожалуйста, введите целые числа")
			continue // Продолжение цикла
		}

		fmt.Println("Выберите операцию:")
		fmt.Println("1. Умножение")
		fmt.Println("2. Деление")
		fmt.Println("3. Сложение")
		fmt.Println("4. Вычитание")

		// Чтение ввода пользователя для выбора операции
		opIn, _ := reader.ReadString('\n')
		opIn = strings.Trim(opIn, "\r\n")
		// Преобразование строки в число
		op, err := strconv.Atoi(opIn)
		// Проверка корректности выбора операции.
		if err != nil || op < 1 || op > 4 {
			fmt.Println("Неверный выбор операции")
			continue
		}

		// Объект для хранения результата операции
		var res big.Int

		switch op {
		case 1:
			// Умножение
			res.Mul(&a, &b)
		case 2:
			// Деление
			res.Div(&a, &b)
		case 3:
			// Сложение
			res.Add(&a, &b)
		case 4:
			// Вычитание
			res.Sub(&a, &b)
		}

		// Вывод результата операции
		fmt.Println("Результат:", res.String())
		fmt.Println("Хотите выполнить еще операцию? (да/нет)")
		// Читаем ввод пользователя для продолжения или завершения программы.
		repeat, _ := reader.ReadString('\n')
		repeat = strings.Trim(repeat, "\r\n")
		if strings.ToLower(repeat) != "да" {
			fmt.Println("Программа завершена.")
			break // Завершаем цикл.
		}
	}
}
