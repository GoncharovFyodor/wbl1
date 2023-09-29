package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)
	var mu sync.Mutex     // Mutex для синхронизации записи
	var rwmu sync.RWMutex // RWMutex для синхронизации чтения и записи
	var wg sync.WaitGroup

	numRoutines := 10 // Количество горутин
	numOps := 100     // Количество операций в горутине

	//Способ 1: Mutex
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				mu.Lock()
				data[n] = j
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	//Вывод результатов
	fmt.Println("Результат работы 1-го способа:")
	for key, value := range data {
		fmt.Printf("%d : %d\n", key, value)
	}

	//Очистка данных
	data = make(map[int]int)

	//Способ 2: RWMutex (чтение)
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				rwmu.RLock()
				_ = data[n]
				rwmu.RUnlock()
			}
		}(i)
	}
	wg.Wait()

	//Вывод результатов
	fmt.Println("Результат работы 2-го способа:")
	for key, value := range data {
		fmt.Printf("%d : %d\n", key, value)
	}

	//Очистка данных
	data = make(map[int]int)

	//Способ 3: RWMutex (запись)
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(n int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				rwmu.Lock()
				data[n] = j
				rwmu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	//Вывод результатов
	fmt.Println("Результат работы 3-го способа:")
	for key, value := range data {
		fmt.Printf("%d : %d\n", key, value)
	}
}
