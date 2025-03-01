package tabulate

import "fmt"

func Tabulate(n int) (s, s1, s2, s3, s4, s5, s6, s7, s8, s9 string) {
	fmt.Printf("This is TABLE %v", n)
	fmt.Println()
	s = fmt.Sprintf("%v * 1 = %v", n, (n * 1))
	s1 = fmt.Sprintf("%v * 2 = %v", n, (n * 2))
	s2 = fmt.Sprintf("%v * 3 = %v", n, (n * 3))
	s3 = fmt.Sprintf("%v * 4 = %v", n, (n * 4))
	s4 = fmt.Sprintf("%v * 5 = %v", n, (n * 5))
	s5 = fmt.Sprintf("%v * 6 = %v", n, (n * 6))
	s6 = fmt.Sprintf("%v * 7 = %v", n, (n * 7))
	s7 = fmt.Sprintf("%v * 8 = %v", n, (n * 8))
	s8 = fmt.Sprintf("%v * 9 = %v", n, (n * 9))
	s9 = fmt.Sprintf("%v * 10 = %v", n, (n * 10))
	return " " + s + "\n", s1 + "\n", s2 + "\n", s3 + "\n", s4 + "\n",
		s5 + "\n", s6 + "\n", s7 + "\n", s8 + "\n", s9
}
