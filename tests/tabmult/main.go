package main

import (
	"fmt"
	ray "numbers/tabulate"
)

func main() {
	fmt.Print("enter a number:... ")
	var num int
	fmt.Scanf("%v", &num)
	fmt.Println()
	fmt.Println(ray.Tabulate(num))
}
