package main

import (
	"example/internal/version"
	"fmt"
)

func main() {
	fmt.Printf("Hello World (%v)\n", version.GetFullVersion())
}
