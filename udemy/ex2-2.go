/*
Modify the previous program to use a func expression.
*/
package main

import "fmt"

func main() {
	result := func(i int) (float64, bool) {
		divided := float64(i) / 2
		var even bool
		if i%2 == 0 {
			even = true
		} else {
			even = false
		}
		return divided, even
	}
	fmt.Println(result(1))
	fmt.Println(result(2))
}
