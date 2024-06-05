package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	out := make(chan int)
	go func() {
		for i := 0; i < 25; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()
	return out
}

func randNumsForCubic(n int) <-chan int {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- rand.Intn(6) + 1
		}
		close(out)
	}()
	return out
}

func main() {
	for num := range randNumsGenerator(18) {
		fmt.Print(num, " ")
	}
	fmt.Println()
	for num := range randNumsForCubic(33) {
		fmt.Print(num, " ")
	}
	fmt.Println()

}
