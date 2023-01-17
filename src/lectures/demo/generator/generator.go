package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()
	return out
}

func generateRandIntn(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()
	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randInt := generateRandIntn(2, 1, 100)

	fmt.Println("generate rand int infite")
	for i := range randInt {
		fmt.Println(i)
	}
}
