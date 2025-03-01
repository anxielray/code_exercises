package bitconverter

import (
	"fmt"
	"os"
)

func Bits(n int) (str string) {
	if n > 255 || n < 1 {
		fmt.Println(`
		try numbers from 1-255.
		the number you provided is out of the scoop that the 9-digit binary
		can present for you!...
		`)
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
