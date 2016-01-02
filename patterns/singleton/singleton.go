package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//import "sync"

type singleton struct {
}

var instance *singleton

// Non thread safe singleton
func GetInstance1() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// Aggressive locking singeton
var mu sync.Mutex

func GetInstance2() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// Check lock check non atomic singleton
func GetInstance3() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

// Check lock check atomic singleton
var initialized uint32

func GetInstance4() *singleton {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// Idiomatic sync once singleton
var once sync.Once

func GetInstance5() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	fmt.Println("Hello singleton")
}
