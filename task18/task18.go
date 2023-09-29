package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализация с мьютексами
type MutexCounter struct {
	num int
	sync.Mutex
}

// Инкрементирование счетчика
func (c *MutexCounter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.num++
}

// Получение значения счетчика
func (c *MutexCounter) Value() int {
	return c.num
}

// Реализация с использованием sync/atomic
type AtomicCounter struct {
	value int64
}

// Инкрементирование счетчика
func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

// Получение значения счетчика
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// Реализация с использованием каналов
type ChannelCounter struct {
	value int
}

// Инкрементирование счетчика
func Inc(counter *ChannelCounter, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	counter.value++
	ch <- counter.value
}

func main() {
	wg := sync.WaitGroup{}
	n := 10 // количество горутин, которые инкрементируют счетчик

	// Использование MutexCounter
	mCnt := &MutexCounter{num: 0}

	// запускаем n горутин, которые увеличивают счетчик на 1
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mCnt.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("Значение счетчика MutexCounter:", mCnt.Value())

	// Использование AtomicCounter
	aCnt := AtomicCounter{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			aCnt.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("Значение счетчика AtomicCounter:", aCnt.Value())

	// Использование ChannelCounter
	chanCnt := ChannelCounter{}
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Inc(&chanCnt, &wg, ch)
	}
	wg.Wait()
	close(ch)

	//Сбор значений из канала
	chanValue := 0
	for value := range ch {
		chanValue = value
	}
	fmt.Println("Значение счетчика ChannelCounter:", chanValue)
}
