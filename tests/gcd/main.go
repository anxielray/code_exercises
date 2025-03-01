package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value, err := strconv.Atoi(os.Args[1])
	value2, _ := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(GCD(value, value2))
}

func GCD(val, val2 int) (gcd int) {
	gcd = 1
	for x := 1; x <= val && x <= val2; x++ {
		if val%x == 0 && val2%x == 0 {
			gcd = x
		}
	}
	return
}
