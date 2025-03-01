package main

import (
	// "fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	size := len(os.Args)
	if size > 1 {
		for _, char := range os.Args[size-1] {
			z01.PrintRune(char)
		}
	}
}
