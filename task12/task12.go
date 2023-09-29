package main

import "fmt"

// Метод 1: Использование map с пустыми значениями
func method1(sequence []string) {
	fmt.Println("Метод 1: Использование map с пустыми значениями")
	uniqueSet := make(map[string]struct{})

	for _, item := range sequence {
		//Запись пустого значения по ключу item
		uniqueSet[item] = struct{}{}
	}

	fmt.Println("Собственное множество (уникальные значения):")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 2: Использование map с булевыми значениями
func method2(sequence []string) {
	fmt.Println("Метод 2: Использование map с булевыми значениями")
	uniqueSet := make(map[string]bool)

	for _, item := range sequence {
		//Запись true по ключу item
		uniqueSet[item] = true
	}

	fmt.Println("Собственное множество (уникальные значения):")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 3: Использование среза для хранения уникальных значений
func method3(sequence []string) {
	fmt.Println("Метод 3: Использование среза для хранения уникальных значений")
	uniqueSet := []string{}

	for _, item := range sequence {
		found := false
		for _, uniqueItem := range uniqueSet {
			// Если текущее значение равно значению из множества уникальных значений,
			// то флагу found присваивается значение true
			if item == uniqueItem {
				found = true
				break
			}
		}
		// Если текущее значение не найдено во множестве уникальных значений, то оно добавляется в это множество
		if !found {
			uniqueSet = append(uniqueSet, item)
		}
	}

	fmt.Println("Собственное множество (уникальные значения):")
	for _, item := range uniqueSet {
		fmt.Println(item)
	}
}

// Метод 4: Использование структуры Set (пользовательский тип, по сути реализации как 2 метод)
type Set map[string]bool

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item string) {
	s[item] = true
}
func method4(sequence []string) {
	fmt.Println("Метод 4: Использование структуры Set")
	uniqueSet := NewSet()

	for _, item := range sequence {
		uniqueSet.Add(item)
	}

	fmt.Println("Собственное множество (уникальные значения):")
	for item := range uniqueSet {
		fmt.Println(item)
	}
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Исходное множество:")
	for _, item := range sequence {
		fmt.Println(item)
	}
	fmt.Println()
	method1(sequence)
	fmt.Println()
	method2(sequence)
	fmt.Println()
	method3(sequence)
	fmt.Println()
	method4(sequence)
	fmt.Println()
}
