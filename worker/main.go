package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// 3 Workers
	for x := 1; x <= 3; x++ {
		go worker(x, jobs, results)
	}

	// Give them jobs
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for the results
	for r := 1; r <= 6; r++ {
		fmt.Println("Result received from worker: ", <-results)
	}
}

func worker(ID int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker ", ID, " is working on job ", job)
		duration := time.Duration(rand.Intn(1e3)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Worker ", ID, " completed work on job ", job, " within ", duration)
		results <- ID
	}
}
