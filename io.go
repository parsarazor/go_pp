package main

import "os"

func main() {
	// Wrting to StdOut 
	os.Stdout.WriteString("Hello, 世界!\n")

	// Writing to Stderr
	os.Stderr.WriteString("Error message\n")
}
