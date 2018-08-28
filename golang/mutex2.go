package main

import (
	"fmt"
	"sync"
)

func main() {
	m := "First Value"
	var lock sync.Mutex
	go func() {
		lock.Lock()
		m = "Second Value"
		lock.Unlock()
	}()
	lock.Lock()
	v := m
	lock.Unlock()
	fmt.Printf("%s\n", v)
}
