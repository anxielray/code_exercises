package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	fmt.Println(Rostring(os.Args[1]))
}

func Rostring(arg string) (rotate string) {
	var reversed []string
	arr := strings.Fields(arg)

	for i := 1; i < len(arr); i++ {
		reversed = append(reversed, arr[i])
	}
	reversed = append(reversed, arr[0])
	rotate = strings.Join(reversed, " ")
	return
}
