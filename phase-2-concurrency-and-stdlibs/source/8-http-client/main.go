package main

import "fmt"

func main() {
	fmt.Println("Basic API get request to Localhost test server:")
	BasicApiGetRequest()

	fmt.Println("\nBasic API get request to Localhost test server but with timeout of 1 sec:")
	BasicApiGetRequestWithTimeout()

	fmt.Println("\nApi With JSON payload:")
	ApiReqWithJsonResponse()

	fmt.Println("\nApi With Nested JSON payload:")
	ApiReqWithNestedJsonResponse()

	fmt.Println("\nSending Custom Requests [PUT With Predefined Headers]:")
	CustomPutRequest()

	fmt.Println("\nSending Custom Requests [GET With Predefined Headers And Req Query]:")
	CustomGetRequestWithQuery()
}
