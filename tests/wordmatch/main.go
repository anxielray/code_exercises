package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(0)
	}
	fmt.Println(WordMatch(os.Args[1], os.Args[2]))
}

func WordMatch(arg, arg2 string) (str string) {
	var wanted int
	for i, char := range arg {
		if RuneExist(char, arg2) {
			str += string(char)
			wanted = Index(char, arg2)
			arg2 = arg2[(wanted + 1):]
			i++
			continue
		}
	}
	if str != arg {
		os.Exit(0)
	}
	return arg
}

func RuneExist(r rune, arg2 string) bool {
	for _, char := range arg2 {
		if char == r {
			return true
		}
	}
	return false
}

func Index(r rune, arg2 string) (indx int) {
	for i, char := range arg2 {
		if char == r {
			indx = i
			break
		}
	}
	return
}
