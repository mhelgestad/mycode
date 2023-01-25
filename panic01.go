package main

import "fmt"

func main() {
    panic("Jim, we have a problem.")

    fmt.Println("You will not even see this line. the panic creates a fast fail.")
}
