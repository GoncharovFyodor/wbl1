package main

import (
	"fmt"
	"log"
	"time"
)

// Функция отправки значений в канал ch
func sender(ch chan<- int) {
	for i := 1; ; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) //Значение отправляется в канал каждые 100 мс
	}
}

// Функция приема значений из канала ch. Канал doneChan используется для сигнализации завершения работы
func receiver(ch <-chan int, doneChan <-chan struct{}) {
	for {
		select {
		case value, ok := <-ch:
			if !ok {
				//Канал ch закрыт, идет завершение работы
				return
			}
			fmt.Println("Полученное значение:", value)
		case <-doneChan:
			//Получен сигнал завершения, идет завершение работы
			return
		}
	}
}

func main() {
	var N int //Количество секунд
	fmt.Println("Введите таймаут выполнения программы (в секундах):")
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		log.Fatal(err)
	}

	//Канал передачи данных
	ch := make(chan int)

	//Канал сигнализации завершения работы
	doneChan := make(chan struct{})

	//Таймер на N секунд
	timeoutChan := time.After(time.Second * time.Duration(N))

	//Запуск горутины для отправки сообщений в канал ch
	go sender(ch)

	//Запуск горутины для чтения сообщений из канала ch
	go receiver(ch, doneChan)

	//Ожидание сигнала завершения или истечения времени
	select {
	case <-timeoutChan:
		//Время выполнения истекло, завершение работы программы
		fmt.Println("Время выполнения истекло")
	case <-doneChan:
		//Получен сигнал завершения, завершение работы программы
	}

	//Закрытие канала сигнала завершения с целью завершения горутин sender и receiver
	close(doneChan)

	// Ожидание завершения горутин sender и receiver возможность завершиться
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Программа завершена")
}
