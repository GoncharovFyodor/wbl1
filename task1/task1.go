package main

import "fmt"

// Структура Human
type Human struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

// Структура Action, в которую встраивается структура Human
type Action struct {
	Human
	university string
}

// Метод для структуры Human
func (h Human) Introduce() {
	fmt.Printf("Привет! Меня зовут %s %s. Мне %d лет. Мой пол: %s\n", h.firstName, h.lastName, h.age, h.gender)
}

// Метод для структуры Action
func (a Action) DescribeUniversity() {
	fmt.Printf("Мой ВУЗ: %s", a.university)
}

func main() {
	//Создание объекта структуры Action
	person := Action{
		Human: Human{
			firstName: "Федор",
			lastName:  "Гончаров",
			age:       21,
			gender:    "Мужской",
		},
		university: "НИУ МИЭТ",
	}

	//person имеет доcтуп к методам Human
	person.Introduce()
	person.DescribeUniversity()
}
