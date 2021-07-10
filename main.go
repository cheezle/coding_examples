package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	type person struct{ name string }
	var wg sync.WaitGroup
	var p atomic.Value

	wg.Add(2)
	go func() {
		defer wg.Done()
		p.Store(person{name: "Steve"})
	}()

	go func(){
		defer wg.Done()
		p.Store(person{name: "Alex"})
	}()

	wg.Wait()
	fmt.Println("person", p.Load())
}