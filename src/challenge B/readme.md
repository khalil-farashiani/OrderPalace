# solution description

`findExceptionNumber` takes a slice of integers arr and returns the exception number by performing a bitwise XOR operation on all the numbers in the slice.
The ^ operator in Go is used for bitwise XOR.

Notice that this solution can respond correct answer only when slice or array item repeated twice except for 1 number
(exactly that is question say) more detail on readme file beside main function
We can start by setting a variable called "result" to 0, and then loop through each element of the slice. For each element, we perform a bitwise XOR operation between the current element and the current value of "result". This means that if the current element is repeated twice, its value will be XORed with itself and result in 0, leaving us with the original value of "result". However, if the current element is the exception number, its value will only be XORed once with "result", leaving us with the value of the exception number.