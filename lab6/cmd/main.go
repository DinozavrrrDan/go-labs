package main

import (
	"fmt"
	"sync"
)

type Turnstile struct {
	count int
	mux   sync.Mutex
}

func (t *Turnstile) IncreaseCount() {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.count++
}

func (t *Turnstile) GetCount() int {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.count
}

func main() {
	turnstile := Turnstile{}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			turnstile.IncreaseCount()
		}()
	}
	wg.Wait()

	fmt.Println("С использованием примитива синхронизации:", turnstile.GetCount())

	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			turnstile.count++ //  без синхронизации
		}()
	}
	wg2.Wait()

	fmt.Println("Без использования примитива синхронизации:", turnstile.GetCount())
}
