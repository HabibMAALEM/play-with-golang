package main

import (
  "fmt"
  "runtime"
)

func main() {
	fmt.Printf("GO: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
}
