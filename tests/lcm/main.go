package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value, err := strconv.Atoi(os.Args[1])
	value2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(LCM(value, value2))
}

func LCM(v, v2 int) (lcm int) {
	var gcd = 1
	for x := 1; x <= v && x <= v2; x++ {
		if v%x == 0 && v2%x == 0 {
			gcd = x
		}
		lcm = (v * v2) / gcd
	}
	return
}
