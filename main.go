package main

import "C"

//export s
func s() string {
	return "HELLO"
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
