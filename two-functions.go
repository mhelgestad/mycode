package main

import "fmt"

func shippingCalc(x, y, z int) int {
    return x + y + z
}

func main() {
    fmt.Println("Enter Height: 42")
    fmt.Println("Enter Width: 13")
    fmt.Println("Enter Length: 20")
    fmt.Println("Total: ", shippingCalc(42, 13, 20))
}
