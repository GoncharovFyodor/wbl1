package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup, doneChan <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case data, ok := <-ch:
			if !ok {
				//Канал ch закрыт, идет завершение работы
				return
			}
			time.Sleep(1500 * time.Millisecond) //Задержка в 1.5 секунды
			fmt.Printf("Воркер %d: Данные %d\n", id, data)
		case <-doneChan:
			//Получен сигнал завершения, идет завершение работы
			return
		}
	}
}

func main() {
	//Количество воркеров
	var N int

	fmt.Println("Введите количество воркеров")
	_, err := fmt.Scanf("%d\n", &N)
	if err != nil {
		log.Fatal(err)
	}
	//Канал для передачи данных от главного потока к воркерам
	dataChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(N)

	//Канал для сигнализации завершения работы воркеров
	doneChan := make(chan struct{})

	//Запуск N воркеров
	for i := 1; i <= N; i++ {
		go worker(i, dataChan, &wg, doneChan)
	}

	//Запись данных в канал из главного потока
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()

	//Ожидание сигнала завершения (Ctrl+C)
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, syscall.SIGTERM)
	<-stopSignalChan

	//Закрытие канала с целью сообщения воркерам о завершении работы
	close(doneChan)

	//Ожидание завершения работы всех воркеров
	wg.Wait()

	//Закрытие канала с целью завершения чтения из dataChan
	close(dataChan)
	fmt.Println("Программа завершена")

}
