package main

import "fmt"

func main() {
    var age int = 42

    fmt.Println("given that age is", age)

    if age%2 == 0 {
        fmt.Println("age is even")
    } else {
        fmt.Println("age is odd")
    }
}
