package main

import "fmt"

func main() {
	fmt.Println(Strlen("Dinero"))
}

func Strlen(s string) (i int) {
	for i = range s {
		i++
	}
	return
}
