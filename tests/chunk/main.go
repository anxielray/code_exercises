package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3))
}

func Chunk(a []int, size int) (chunks [][]int) {
	var length, i int
	for length = range a {
		length++
	}
	if size == 0 {
		os.Exit(0)
	}
	if length%size == 0 {
		chunks = append(chunks, a[0:size])
		for i < len(a) {
			i = size
			chunks = append(chunks, a[i:(i+i)])
			a = a[i:]
		}
	} else {
		working := a[:(length - (length % size))]
		chunks = append(chunks, working[0:size])
		for i < len(working) {
			if len(working) != size {
				i = size
				chunks = append(chunks, working[i:(i+i)])
				working = working[i:]
			} else {
				break
			}

		}
		remainder := a[length-(length%size):]
		chunks = append(chunks, remainder)
	}
	return
}
