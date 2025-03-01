package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]
	output := PigLatin(input)
	fmt.Println(output)
}

func PigLatin(input string) string {
	vowels := "aeiouAEIOU"
	words := strings.Split(input, " ")
	for i, word := range words {
		if !hasVowel(word) {
			words[i] = "No vowels"
		} else {
			if strings.HasPrefix(word, strings.ToLower(vowels)) {
				words[i] = word + "ay"
			} else {
				for j, char := range word {
					if strings.Contains(vowels, strings.ToLower(string(char))) {
						words[i] = word[j:] + word[:j] + "ay"
						break
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func hasVowel(word string) bool {
	vowels := "aeiouAEIOU"
	for _, char := range word {
		if strings.Contains(vowels, string(char)) {
			return true
		}
	}
	return false
}
