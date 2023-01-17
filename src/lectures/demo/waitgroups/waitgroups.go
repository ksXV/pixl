package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	counter := 0

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		counter += 1
		go func() {
			defer func() {
				fmt.Println(counter, "goroutines remaining")
				counter--
				wg.Done()
			}()
			duration := time.Duration(rand.Intn(5000) * int(time.Millisecond))
			fmt.Println("Waiting :", duration)
			time.Sleep(duration)
		}()
	}
	fmt.Println("Waiting for go routines to finish")
	wg.Wait()
	fmt.Println("Done")
}
