package main

import (
	"fmt"
	"sync"
)

func main() {
	m := "First Value"
	var lock sync.Mutex
	lock.Lock()
	go func() {
		m = "Second Value"
		lock.Unlock()
	}()
	lock.Lock()
	v := m
	lock.Unlock()
	fmt.Printf("%s\n", v)
}
