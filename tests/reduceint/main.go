package main

import "fmt"

func add(m, n int) int {
	return m + n
}
func multiply(m, n int) int {
	return m * n
}
func divide(m, n int) int {
	return m / n
}
func main() {
	a := []int{500, 2}
	fmt.Println(reduceInt(a, add))
	fmt.Println(reduceInt(a, multiply))
	fmt.Println(reduceInt(a, divide))
}

func reduceInt(a []int, f func(int, int) int) (result int) {
	result = a[0]
	for _, num := range a[1:] {
		result = f(result, num)
	}
	return
}
