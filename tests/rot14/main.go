package main

import "fmt"

func main() {
	fmt.Print(
		`Watch what happens when you inplace a string: `)
	fmt.Println()
	var myR string
	fmt.Scanf("%s", &myR)
	fmt.Printf(`for you to obtain %s, you'd have to insert "%s"`, myR, RevRot14(myR))
	fmt.Println()
	fmt.Println()
	fmt.Println(
		`		The logic behind this concept,
		is a mathematical confrontation
		of the skills of critical  ability
		to think and reasoning. Write the alphabet
		from 'a' all the way to 'z' vertically.
		for 'a' you'll realize that you only 
		need to find its sum with 12.
		For letters from 'o', we find their
		difference with 14.`)
	fmt.Println()
}

// func Rot14(s string) string {
// 	var str string
// 	for _, r := range s {
// 		if (r >= 'a' && r <= 'l') || (r >= 'A' && r <= 'L') {
// 			r += 14
// 			str += string(r)
// 		} else if r >= 'm' && r <= 'z' {
// 			r = 'a' + ((r + 14) - ('z' + 1))
// 			str += string(r)
// 		} else if r >= 'M' && r <= 'Z' {
// 			r = 'A' + ((r + 14) - ('Z' + 1))
// 			str += string(r)
// 		}
// 	}
// 	return str
// }

func RevRot14(s string) string {
	var str string
	for _, r := range s {
		if (r >= 'a' && r <= 'n') || (r >= 'A' && r <= 'N') {
			r += 12
			str += string(r)
		} else if (r >= 'o' && r <= 'z') || (r >= 'O' && r <= 'Z') {
			r -= 14
			str += string(r)
		}
	}
	return str
}
