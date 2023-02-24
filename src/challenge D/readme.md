# challenge D

This code is written in Go and demonstrates the use of goroutines and channels to create a pipeline for processing data. The program calculates the sum of squared numbers in a slice of integers using a worker pool.

The powerSum function takes a slice of integers as an argument, squares each integer, and returns the sum of squared numbers.

The worker function takes an integer from an input channel, squares it, and sends the result to an output channel. This function is designed to be run concurrently by multiple goroutines.

The main function defines a slice of integers and the number of worker goroutines to use. It creates two channels, numsChan and squaredChan, and sends the integers to numsChan using a goroutine. It then starts multiple worker goroutines that receive integers from numsChan, square them, and send the results to squaredChan. Finally, it collects the squared numbers from squaredChan, calculates the sum of squared numbers using powerSum, and prints the result to the console.

To run this code, save it to a file with a ".go" extension and use the Go compiler to build an executable file. Then, execute the file from the command line. The number of worker goroutines can be adjusted by changing the numWorkers variable in the main function. Note that the numbers used in this example are arbitrary and can be replaced with any slice of integers.




