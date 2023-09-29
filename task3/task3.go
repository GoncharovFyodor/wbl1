package main

import (
	"fmt"
	"sync"
)

// Функция вычисления квадрата числа num с последующей его отправкой в канал resChan
func calculateSquare(num int, wg *sync.WaitGroup, resChan chan int) {
	defer wg.Done()
	square := num * num
	resChan <- square
}

// Функция вычисления квадрата числа num с последующим добавлением его к общей сумме.
// Для защиты переменной общей суммы от race condition при параллельном доступе используется mutex
func calculateSquareSumWithMutex(num int, wg *sync.WaitGroup, mu *sync.Mutex, totalSum *int) {
	defer wg.Done()
	square := num * num
	mu.Lock()
	*totalSum += square
	mu.Unlock()
}

func main() {
	//Массив чисел (2, 4, 6, 8, 10)
	arr := []int{2, 4, 6, 8, 10}

	//==Вариант 1: с использованием каналов==
	fmt.Println("==Вариант 1: с использованием каналов==")
	//Создание буферизированного канала для передачи результатов вычислений
	resChan := make(chan int, len(arr))
	//Инициализация WaitGroup для синхронизации горутин
	var chanWg sync.WaitGroup

	//Запуск горутин для расчета квадратов чисел
	for _, num := range arr {
		chanWg.Add(1)
		go calculateSquare(num, &chanWg, resChan)
	}

	//Дополнительная горутина, которая ожидает завершения всех остальных горутин,
	//после чего закроет канал resChan
	go func() {
		chanWg.Wait()
		close(resChan)
	}()

	chanTotalSum := 0
	//Основная горутина - вывод квадратов чисел из resChan в stdout
	for square := range resChan {
		chanTotalSum += square
	}
	fmt.Println("Сумма квадратов чисел (с использованием каналов):", chanTotalSum)

	//==Вариант 2: с использованием мьютексов==
	fmt.Println("==Вариант 2: с использованием мьютексов==")
	mutexTotalSum := 0

	//WaitGroup для ожидания завершения горутин вычисления квадратов
	var mutexWg sync.WaitGroup

	//Mutex для защиты переменной общей суммы от race condition при параллельном доступе
	var mu sync.Mutex

	for _, num := range arr {
		mutexWg.Add(1)
		go calculateSquareSumWithMutex(num, &mutexWg, &mu, &mutexTotalSum)
	}

	//Ожидание завершения всех горутин
	mutexWg.Wait()
	fmt.Println("Сумма квадратов чисел (с использованием мьютексов):", mutexTotalSum)
}
