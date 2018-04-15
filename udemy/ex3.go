package main

import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Please, enter your name:")
    name, _ := reader.ReadString('\n')
    fmt.Println("Hello, " + name)
}
