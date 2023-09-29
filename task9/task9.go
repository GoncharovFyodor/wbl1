package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{2, 3, 5, 7, 11, 13, 17, 19}

	inCh := make(chan int)
	outCh := make(chan int)
	var wg sync.WaitGroup

	//Горутина для записи чисел в inCh
	go func() {
		for _, x := range input {
			inCh <- x
		}
		close(inCh)
	}()

	//Горутина для умножения чисел из inCh и оправки результатов в outCh
	go func() {
		for num := range inCh {
			outCh <- num * 2
		}
		close(outCh)
	}()

	//Горутина для вывода данных из outCh в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range outCh {
			time.Sleep(time.Second)
			fmt.Print(x, " ")
		}
		fmt.Println()
	}()

	//Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("Программа завершена")
}
