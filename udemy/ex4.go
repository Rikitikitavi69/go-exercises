package main

import (
    "fmt"
)

func main() {
    var num1, num2 int
    fmt.Print("Please enter small number:")
    fmt.Scan(&num1)
    fmt.Print("Please enter larger number:")
    fmt.Scan(&num2)
    remainder := num2 % num1
    fmt.Println("Remainder ", num2, " % ", num1, " = ", remainder)
}
