package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(prevprime(num))
}

func findPrime(n int) (prime bool) {
	if n < 2 {
		return false
	}
	for x := 2; x < n; x++ {
		if n%x == 0 {
			return false
		}
	}
	return true
}

func prevprime(prime int) (prev int) {
	for i := prime; i >= prime-10; i-- {
		if !findPrime(i) {
			prev += 0
		} else {
			prev += i
			break
		}
	}
	return
}
