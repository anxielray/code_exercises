package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Memory(os.Args[1]))
	fmt.Println(os.Args[1])
}
func Memory(arg string) (memory string) {
	arr := []rune(arg)
	for _, str := range arr {
		memory += fmt.Sprintf("%x ", str)
	}
	return
}
