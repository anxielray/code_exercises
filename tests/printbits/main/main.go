package main

import (
	bits "bits/BitConverter"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		fmt.Println(`this is the printbits function that converts your decimal number to its corresponding binary number.`)
		fmt.Println("pass a number as an argument...")
		os.Exit(0)
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("the argument you passed cannot be converted to a number...")
		os.Exit(0)
	}
	fmt.Println(bits.Bits(num))
}
