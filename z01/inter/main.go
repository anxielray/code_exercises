package main

import (
	"fmt"
	"os"
	"strings"
)

func inter(s1, s2 string) {
	seen := make(map[rune]bool)
	var result []rune

	for _, char := range s1 {
		if strings.ContainsRune(s2, char) && !seen[char] {
			result = append(result, char)
			seen[char] = true
		}
	}

	fmt.Println(string(result))
}

func main() {
	if len(os.Args) == 3 {
		inter(os.Args[1], os.Args[2])
	}
}

