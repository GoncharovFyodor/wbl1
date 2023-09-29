package main

import (
	"fmt"
	"sort"
)

// Пересечение двух множеств с использованием словарей (Maps)
func intersectionWithMap(slice1, slice2 []int) []int {
	//Преобразование срезов в словари для быстрого доступа
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, el := range slice1 {
		set1[el] = true
	}
	for _, el := range slice2 {
		set2[el] = true
	}

	var res []int

	//Перебор элементов первого словаря
	for el := range set1 {
		//Если элемент также есть во втором словаре
		if set2[el] {
			res = append(res, el)
		}
	}
	return res
}

// Пересечение двух множеств с использованием срезов (Slices)
func intersectionWithSlice(slice1, slice2 []int) []int {
	//Сортируем срезы для удобства сравнения
	sort.Ints(slice1)
	sort.Ints(slice2)

	var res []int
	i, j := 0, 0

	// Проходим по срезам, сравнивая элементы
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] == slice2[j] {
			// Если элементы равны, добавляем в результат и двигаемся к следующим элементам
			res = append(res, slice1[i])
			i++
			j++
		} else if slice1[i] < slice2[j] {
			// Если элемент из первого множества меньше, увеличиваем индекс в первом множестве
			i++
		} else {
			// Если элемент из второго множества меньше, увеличиваем индекс во втором множестве
			j++
		}
	}
	return res
}

func main() {
	// Пример использования методов
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{5, 3, 10, 7, 6}
	resultMap := intersectionWithMap(set1, set2)
	fmt.Println("Пересечение с помощью словарей:", resultMap)

	resultSlice := intersectionWithSlice(set1, set2)
	fmt.Println("Пересечение с помощью срезов:", resultSlice)
}
