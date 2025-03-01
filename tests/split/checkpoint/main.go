package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Split(os.Args[1], os.Args[2])
}

func Split(arg, seperator string) {
	for i, char := range arg {
		var x int
		if string(char) == seperator {
			x = i
			z01.PrintRune(10)
		} else {
			z01.PrintRune(char)
		}
		arg = arg[:x] + arg[(x+1):]
		Split(arg, seperator)
		break
	}
}
