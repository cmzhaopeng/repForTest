package main

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(1 * time.Second)
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	return int(result.Int64())
}

func main() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for i := range dataChan {
		println(i)
	}
}
