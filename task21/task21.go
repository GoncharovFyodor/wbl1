package main

import "fmt"

// Европейская розетка
type EuroSocket struct{}

// Европейская вилка
type EuroPlug interface {
	getPower() string
}

// Американская вилка
type AmericanPlug interface {
	getPower() string
}

// Европейская вилка (реализация интерфейса AmericanPlug)
type SimpleAmericanPlug struct{}

// Европейская вилка (реализация интерфейса EuroPlug)
type SimpleEuroPlug struct{}

// Рабочее напряжение американской вилки
func (s SimpleAmericanPlug) getPower() string {
	return "110V"
}

// Рабочее напряжения европейской вилки
func (s SimpleEuroPlug) getPower() string {
	return "220V"
}

// Розетка в которую подключается европейская вилка
func (socket EuroSocket) connect(plug EuroPlug) {
	fmt.Println(plug.getPower())
}

// Адаптер американской вилки (переходник с американской вилки на европейскую)
type AmericanPlugAdapter struct {
	AmericanPlug
}

// Получение рабочего напряжения
func (adapter *AmericanPlugAdapter) getPower() string {
	return adapter.AmericanPlug.getPower()
}

func main() {
	euroSocket := EuroSocket{}

	euroPlug := SimpleEuroPlug{}

	amPlug := SimpleAmericanPlug{}

	// Переходник с американской вилки на европейскую
	amPlugAdapter := AmericanPlugAdapter{&amPlug}

	euroSocket.connect(&euroPlug)
	euroSocket.connect(&amPlugAdapter)
}
