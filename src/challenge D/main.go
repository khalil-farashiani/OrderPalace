package main

import "fmt"

func powerSum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n * n
	}
	return sum
}

func worker(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
}

func main() {
	nums := []int{12, 54, 89, 21, 66, 47, 14, 285, 96}
	numWorkers := 2

	// create the pipeline
	numsChan := make(chan int)
	squaredChan := make(chan int)
	go func() {
		defer close(numsChan)
		for _, n := range nums {
			numsChan <- n
		}
	}()
	// start the workers
	for i := 0; i < numWorkers; i++ {
		go worker(numsChan, squaredChan)
	}
	// collect the results
	squaredNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		squaredNums[i] = <-squaredChan
	}
	close(squaredChan)
	// calculate the sum of squared numbers
	sum := powerSum(squaredNums)
	fmt.Printf("Sum of squared numbers: %d\n", sum)
}
