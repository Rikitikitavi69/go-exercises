/*
	What's the value of this expression: (true && false) || (false && true) || !(false && false)?
*/
package main

import "fmt"

func main() {
	expression := (true && false) || (false && true) || !(false && false)
	//                  false     or      false      or        true
	//                                    true
	fmt.Println(expression)
}
