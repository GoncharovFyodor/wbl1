package main

import "fmt"

// Удаление i-го элемента из слайса с использованием функции append
func remove1(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

// Удаление i-го элемента из слайса с использованием функции copy
func remove2(slice []int, i int) []int {
	// Копирование элементов справа от i-го элемента на одну позицию влево
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{2, 3, 5, 7, 11}
	fmt.Println("Слайс до удаления:", slice)
	i := 3 //Индекс удаляемого элемента

	if i < len(slice) {
		fmt.Println("Удаление i-го элемента с помощью append")
		slice1 := remove1(slice, i)
		fmt.Println("Слайс после удаления:", slice1)
		fmt.Println()
		fmt.Println("Удаление i-го элемента с помощью copy")
		slice2 := remove2(slice, i)
		fmt.Println("Слайс после удаления:", slice2)
	}
}
