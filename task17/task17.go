package main

import (
	"fmt"
	"sort"
)

func main() {
	// Исходный срез
	arr := []int{1, 2, 3, 4, 5, 6, 8, 15, 22}
	fmt.Println("Исходный срез:", arr)

	// Искомый элемент
	target := 8
	fmt.Println("Искомый элемент:", target)

	// Для бинарного поиска используется функция sort.SearchInts()
	index := sort.SearchInts(arr, target)

	// Проверка результата
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Число %d найдено, оно имеет индекс %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено")
	}
}
