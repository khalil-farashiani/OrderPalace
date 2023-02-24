package main

import "fmt"

// findExceptionNumber takes a slice of integers arr and
// returns the exception number by performing a bitwise XOR operation on all the numbers in the slice.
// The ^ operator in Go is used for bitwise XOR.
// Notice that this solution can respond correct answer only when slice or array item repeated twice except for 1 number
// (exactly that is question say) more detail on readme file beside main function
func findExceptionNumber(arr []int) int {
	result := 0
	for _, num := range arr {
		result ^= num
	}
	return result
}
func main() {
	// this is sample for testing you can change this slice
	arr := []int{2, 2, 5, 6, 5, 6, 9}
	exceptionNumber := findExceptionNumber(arr)
	fmt.Println(exceptionNumber) // prints 9
}
