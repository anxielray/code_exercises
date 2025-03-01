package main

import "fmt"

func main() {
	fmt.Println(SortWord([]string{"a", "2", "3", "$", ",", " ", "/", "1", "A", "!", "#", ")"}))
}

func SortWord(ascii []string) (sorted []string) {
	if len(ascii) == 1 {
		sorted = append(sorted, ascii[0])
	}
	var index int
	small := ascii[0]
	for i := range ascii {
		if ascii[i] < small {
			index = i
			small = ascii[i]
		}
	}
	smallest := small
	for i := 1; i < len(ascii); i++ {
		sorted = append(sorted, smallest)
		asci := append(ascii[:index], ascii[(index+1):]...)
		sorted = append(sorted, SortWord(asci)...)
		break
	}
	return
}
