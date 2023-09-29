package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var stopFlag bool // Флаг для остановки горутин

// Использование каналов для синхронизации горутин
func goroutineUsingChannels() {
	fmt.Println("С использованием каналов:")
	quitCh := make(chan struct{}) //Канал для сигнала завершения горутины
	go func() {
		defer fmt.Println("Горутина на каналах: Завершено!")
		defer close(quitCh)
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на каналах:", i)
			time.Sleep(time.Second)
		}
	}()
	//Ожидание сигнала завершения
	<-quitCh
}

// Использование контекста для остановки горутины
func goroutineUsingContext() {
	fmt.Println("С использованием контекста:")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer fmt.Println("Горутина на контексте: Завершено!")
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Горутина на контексте:", i)
				time.Sleep(time.Second)
			}
		}
	}()

	//Ожидание сигнала завершения (нажатия клавиши)
	fmt.Println("Нажмите Enter, чтобы остановить горутину")
	fmt.Scanln()
	cancel()
	// Ожидание для завершения вывода перед завершением программы
	time.Sleep(3 * time.Second)
}

// Использование WaitGroup для синхронизации горутин
func goroutineUsingWg() {
	fmt.Println("С использованием WaitGroup:")
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer fmt.Println("Горутина на WaitGroup: Завершено!")
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("Горутина на WaitGroup:", i)
			time.Sleep(time.Second)
		}
	}()

	//Ожидание завершения горутины
	wg.Wait()
}

// Использование сигналов для остановки горутины
func goroutineUsingSignals() {
	fmt.Println("С использованием сигналов:")
	stopSignalChan := make(chan os.Signal, 1)
	signal.Notify(stopSignalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan struct{})

	go func() {
		defer fmt.Println("Горутина на сигналах: Завершено!")
		<-stopSignalChan
		close(doneChan)
	}()

	//Ожидание сигнала завершения (Ctrl+C)
	fmt.Println("Нажмите Ctrl+C, чтобы остановить горутину на сигналах")
	<-doneChan
}

// Использование флага для остановки горутины
func goroutineUsingFlag() {
	fmt.Println("С использованием флага:")

	stopFlag = false
	go func() {
		defer fmt.Println("Горутина c флагом: Завершено!")
		for i := 1; ; i++ {
			if stopFlag {
				return
			}
			fmt.Println("Горутина с флагом:", i)
			time.Sleep(time.Second)
			if i == 3 {
				runtime.Goexit() //Завершение текущей горутины
			}
		}
	}()

	// Ожидание нажатия клавиши для остановки
	fmt.Println("Нажмите Enter, чтобы остановить горутину")
	fmt.Scanln()
	stopFlag = true
}

// Использование паники для остановки горутины (плохая практика!)
func goroutineUsingPanic() {
	fmt.Println("С использованием паники:")

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Паника перехвачена:", r)
			}
		}()
		defer fmt.Println("Горутина c паникой: Завершено!")
		for i := 1; ; i++ {
			fmt.Println("Горутина с паникой:", i)
			if i == 3 {
				panic("Panic") //Имитация паники
			}
			time.Sleep(time.Second)
		}
	}()

	// Ожидание для завершения вывода
	time.Sleep(3 * time.Second)
}

// Использование runtime.Goexit для остановки горутины
func goroutineUsingGoexit() {
	fmt.Println("С использованием runtime.Goexit:")

	go func() {
		defer fmt.Println("Горутина на runtime.Goexit: Завершено!")
		for i := 1; i < 5; i++ {
			fmt.Println("Горутина на runtime.Goexit:", i)
			time.Sleep(time.Second)
			if i == 3 {
				runtime.Goexit() //Завершение текущей горутины
			}
		}
	}()

	// Ожидание для завершения вывода перед завершением программы
	time.Sleep(5 * time.Second)
}

func main() {
	goroutineUsingChannels()
	fmt.Println()

	goroutineUsingContext()
	fmt.Println()

	goroutineUsingWg()
	fmt.Println()

	goroutineUsingSignals()
	fmt.Println()

	goroutineUsingFlag()
	fmt.Println()

	goroutineUsingPanic()
	fmt.Println()

	goroutineUsingGoexit()
	fmt.Println()

}
