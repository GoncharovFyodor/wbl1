package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 14, 1, 7, 9, 8, 11, 6, 4, 2}
	fmt.Println("До сортировки:", arr)

	// Для быстрой сортировки в Go есть встроенный метод в пакете sort
	sort.Ints(arr)

	fmt.Println("После сортировки:", arr)
}
