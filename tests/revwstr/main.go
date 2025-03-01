package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Revwstr(os.Args[1]))
}

func Revwstr(arg string) (rev string) {
	var reversed []string
	arr := strings.Fields(arg)

	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	rev = strings.Join(reversed, " ")
	return
}
