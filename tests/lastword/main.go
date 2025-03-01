package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(LastWord(os.Args[1]))
}

func LastWord(arg string) (last string) {
	args := strings.Fields(arg)
	return args[len(args)-1]
}
