package main

import (
    "fmt"
    "os"
)

func main() {
    
    argsWithProg := os.Args

    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    last_arg := os.Args[len(os.Args)-1]

    fmt.Println("All arguments, including the program name:")
    fmt.Println(argsWithProg)

    fmt.Println("All arguments, (without the program name:")
    fmt.Println(argsWithoutProg)

    fmt.Println(arg)

    fmt.Println(last_arg)
}
