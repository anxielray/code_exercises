package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// fmt.Println("0123456789")
	s := "0123456789"
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(10 )
}
