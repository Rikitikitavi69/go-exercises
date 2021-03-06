/*
	Write a function which takes an integer.
	The function will have two returns.
	The first return should be the argument divided by 2.
	The second return should be a bool that let’s us know whether or not the argument was even.
	For example:
			half(1) returns (0, false)
			half(2) returns (1, true)

*/
package main

import (
	"fmt"
)

func half(i int) (float32, bool) {
	divided := float32(i) / 2
	var even bool
	if i%2 == 0 {
		even = true
	} else {
		even = false
	}
	return divided, even
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
