package main

import "fmt"

func main(){
	fmt.Println(Max([]int{1, 2, 3, 4, 5}))
}

func Max(a []int)(max int) {
	max = a[0]
	for _, num := range a {
		if num > max {
			max = num
		}
	}
	return
}