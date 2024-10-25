package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting the function...")
	start := time.Now() // Record the start time
	sleepOneSecond()
	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Time elapsed: %v\n", elapsed)
	fmt.Println("Function finished!")
	sleepOneSecond()
	elapsed = time.Since(start) // Calculate the elapsed time
	fmt.Printf("Time elapsed: %v\n", elapsed)
	fmt.Println("Function 2 finished!")

	var wg sync.WaitGroup
	wg.Add(2)
	start = time.Now() // Record the start time
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		sleepOneSecond()
	}()

	go func() {
		defer wg.Done()
		sleepOneSecond()
	}()
	wg.Wait()
	elapsed = time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
	fmt.Println("All goroutines completed!")
}

func sleepOneSecond() {
	time.Sleep(1 * time.Second)
}
