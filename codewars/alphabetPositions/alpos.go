package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(alps("The sunset sets at twelve o' clock."))
}

func alps(s string) string {
	var result string
	arr := strings.Split(s, " ")
	for i, v := range arr {
		for x, c := range v {
			if c >= 'A' && c <= 'Z' {
				if i != len(arr)-1 {
					result += strconv.Itoa(int(c-64)) + " "
				} else if (i == len(arr)-1) && (x == len(v)-1) {
					result += strconv.Itoa(int(c - 96))
				}
			} else if c >= 'a' && c <= 'z' {
				if i != len(arr)-1 {
					result += strconv.Itoa(int(c-96)) + " "
				} else if (i == len(arr)-1) && (x == len(v)-1) {
					result += strconv.Itoa(int(c - 96))
				}
			} else {
				continue
			}
		}
	}
	return result
}
