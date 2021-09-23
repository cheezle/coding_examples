package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var p atomic.Value

	type person struct { name string }

	wg.Add(2)
	go func() {
		defer wg.Done()
		p.Store(person{name: "Steve"})
	}()

	go func() {
		defer wg.Done()
		p.Store(person{name: "Alex"})
	}()

	wg.Wait()
	fmt.Println("person", p.Load())
}