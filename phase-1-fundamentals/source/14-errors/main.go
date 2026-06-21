package main

func main() {

	// 01. Returning errors as values in functions....
	BasicErrors()

	// 02. Sentinel Errors: Packages returning predefined errors instread of strings
	SentinelErrors()

	// 03. Wrapped Errors: Packages returning predefined errors instread of strings along with additional info at diffrent places
	WrappedErrors()
}
