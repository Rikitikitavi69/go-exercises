/*
Find a problem at projecteuler.net then create the solution.
Add a comment beneath your solution that includes a description of the problem.
Upload your solution to github. Tweet me a link to your solution.

https://projecteuler.net/problem=1 :

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import "fmt"

func findSum(maxN int) int {
	var slice []int
	for i := 1; i < maxN; i++ {
		if i%3 == 0 || i%5 == 0 {
			slice = append(slice, i)
		}
	}
	var mul int
	for _, v := range slice {
		mul += v
	}
	return mul

}

func main() {
	result := findSum(1000)
	fmt.Println(result)
}
