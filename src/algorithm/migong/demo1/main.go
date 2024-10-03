package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {
	n := 4
	test(n)
}
