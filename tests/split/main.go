package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Split(os.Args[1], " "))
}

func Split(arg, seperator string) (split string) {
	split = strings.ReplaceAll(arg, seperator, "\n")
	return
}
