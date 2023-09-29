package main

import (
	"fmt"
	"time"
)

// Реализация sleep с помощью time.Sleep
func SleepUsingTime(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Реализация sleep с помощью таймера
func SleepUsingTimer(seconds int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
}

// Реализация sleep с помощью горутины и канала
func SleepUsingGoroutine(seconds int, done chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	done <- true
}

// Реализация sleep с помощью select
func SleepUsingSelect(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}

func main() {
	fmt.Println("В самом начале")

	// Используя time.Sleep
	SleepUsingTime(1)
	fmt.Println("После time.Sleep")

	// Используя таймер
	SleepUsingTimer(2)
	fmt.Println("После таймера")

	// Используя горутину и канал
	done := make(chan bool)
	go SleepUsingGoroutine(3, done)
	<-done
	fmt.Println("После горутины и канала")

	// Используя select
	SleepUsingSelect(4)
	fmt.Println("После select")
}
