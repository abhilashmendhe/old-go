package main

import (
	"sync"
	"testing"
)

func TestDataRaceConditions(t *testing.T) {

	var state int32

	// state2
	// state3
	// state4
	var mu sync.RWMutex

	for i := 0; i < 10; i++ {
		go func(i int) {
			mu.Lock()
			state += int32(i)
			// business logic
			mu.Unlock()
		}(i)
	}
	// fmt.Println(state)
}
