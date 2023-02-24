# challenge C

The requestURL function takes a single URL as an argument, sends an HTTP GET request to the URL using an HTTP client with a timeout of one second, and returns the response body as a string through a channel.

The requestURLs function takes a slice of URLs, creates a channel to receive the response strings, and launches a goroutine for each URL that calls the requestURL function with the URL and the channel as arguments. It then waits for all responses to be received within one second using a select statement with a timeout. If all requests are successful, it prints the response bodies to the console. Otherwise, it prints an error message.

The allSuccess function takes a slice of response strings and checks if all responses were successful. A response is considered successful if it is not empty and does not start with the prefixes "Request failed:" or "Request timed out".

The main function defines a slice of three URLs and calls the requestURLs function with the slice as an argument.

To run this code, save it to a file with a ".go" extension and use the Go compiler to build an executable file. Then, execute the file from the command line. Note that the URLs used in this example may change or become unavailable over time, so you may need to modify the code to use different URLs.