package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// fmt.Println(len(os.Args)-1)
	count := len(os.Args) - 1
	for _, char := range myItoa(count) {
		z01.PrintRune(char)
	}
}

func myItoa(num int) string {
	var result []byte
	var temp string
	if num == 0 {
		result = append(result, '0')
	}

	if num < 0 {
		result = append(result, '-')
		num = -num
	}

	for num > 0 {
		result = append(result, byte(num%10)+'0')
		num /= 10
	}

	for len(result) > 0 {
		temp = string(result[len(result)-1]) + temp
		result = result[:len(result)-1]
	}

	return temp
}
