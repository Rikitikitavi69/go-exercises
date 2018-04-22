/*
   Create a program that prints to the terminal asking for a user to enter their name.
   Use fmt.Scan to read a user’s name entered at the terminal.
   Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please, enter your name:")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello, " + name)
}
