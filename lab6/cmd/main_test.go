package main

import (
	"sync"
	"testing"
)

func TestTurnstile(t *testing.T) {
	var wg sync.WaitGroup
	turnstile := Turnstile{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			turnstile.IncreaseCount()
		}()
	}
	wg.Wait()

	if turnstile.GetCount() != 100 {
		t.Errorf("Ожидается: %d, Фактическое: %d", 100, turnstile.GetCount())
	}

}
