package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Swapbits(os.Args[1]))
}

func Swapbits(arg string) (swapped string) {
	var (
		bit   = Bits(Atoi(arg))
		half  = len(bit) / 2
		first = bit[:half]
		last  = bit[half:]
	)
	swapped = first + " | " + last + "\n" + "   \\" + "   /" +
		"\n" + "   /" + "   \\" + "\n" + last + " | " + first
	return
}

func Bits(n int) (str string) {
	if n > 255 || n < 1 {
		os.Exit(0)
	}
	if n >= 128 && n < 256 {
		str += "1"
		n -= 128
	} else {
		str += "0"
	}
	if n >= 64 && n < 128 {
		str += "1"
		n -= 64
	} else {
		str += "0"
	}
	if n >= 32 && n < 64 {
		str += "1"
		n -= 32
	} else {
		str += "0"
	}
	if n >= 16 && n < 32 {
		str += "1"
		n -= 16
	} else {
		str += "0"
	}
	if n >= 8 && n < 16 {
		str += "1"
		n -= 8
	} else {
		str += "0"
	}
	if n >= 4 && n < 8 {
		str += "1"
		n -= 4
	} else {
		str += "0"
	}
	if n >= 2 && n < 4 {
		str += "1"
		n -= 2
	} else {
		str += "0"
	}
	if n%2 != 0 {
		str += "1"
	} else {
		str += "0"
	}
	return str
}

func Atoi(str string) (num int) {
	var digit, number = 1, 1
	for x := len(str) - 1; x >= 0; x-- {
		number = int(rune(str[x]) - 48)
		num += (digit * number)
		digit *= 10
	}
	return
}
