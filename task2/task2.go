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

func main() {
	//Массив чисел (2, 4, 6, 8, 10)
	arr := []int{2, 4, 6, 8, 10}

	//Создание канала для обмена результатами между горутинами
	resChan := make(chan int)

	//Инициализация WaitGroup для синхронизации горутин
	var wg sync.WaitGroup

	//Запуск горутин для расчета квадратов чисел
	for _, num := range arr {
		wg.Add(1)
		go calculateSquare(num, &wg, resChan)
	}

	//Дополнительная горутина, которая ожидает завершения всех остальных горутин,
	//после чего закроет канал resChan
	go func() {
		wg.Wait()
		close(resChan)
	}()

	//Основная горутина - вывод квадратов чисел из resChan в stdout
	for square := range resChan {
		fmt.Println(square)
	}

	//Главная горутина (main) создает отдельную горутину для каждого числа,
	//и каждая горутина вычисляет квадрат этого числа асинхронно.
	//Затем главная горутина считывает результаты из канал resChan и выводит их в stdout.
	//После завершение работы всех горутин главная горутина закрывает канал resChan,
	//что означает, что все данные получены, и таким образом главная горутина
	//выходит из цикла чтения из канала.
}
