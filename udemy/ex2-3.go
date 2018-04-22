/*
	Write a function with one variadic parameter that
	finds the greatest number in a list of numbers.
*/
package main

import "fmt"

func findMax(numbers ...int) int {
	var maxNumber int
	for _, v := range numbers {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber
}

func main() {
	n := []int{1, 2, 3, 4, 5, 8, 6, 4, 9, 1, 122, 75}
	maxN := findMax(n...)
	fmt.Println(maxN)
}
