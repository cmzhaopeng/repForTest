package main

import "sync"

var once sync.Once
var wg sync.WaitGroup

func initial() {
	println("initial")
}

func dostuff() {
	defer wg.Done()
	once.Do(initial)
	println("dostuff")
}

func main() {
	wg.Add(2)
	go dostuff()
	go dostuff()

	wg.Wait()

}
