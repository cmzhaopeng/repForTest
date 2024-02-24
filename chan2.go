package main

import (
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {

	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func main() {
	dataChan := make(chan int)

	wg := sync.WaitGroup{}

	go func() {

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		wg.Wait()
		close(dataChan)

	}()
	wg.Add(1)
	wg.Done()
	wg.Done()
	wg.Done()
	wg.Wait()
	for i := range dataChan {
		println(i)
	}

}
