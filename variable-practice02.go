package main

import (
    "fmt"
)

var zzz string = "package-level (global) variable\n"

func main() {
    fmt.Printf(zzz)

    var mug string = "filled with tea\n"
    fmt.Printf(mug)


    var a = "Hello"
    var b = 23
    var c = true
    var d = 2.3
    
    //    var a = "Goodbye"
    a = "Goodbye"

    fmt.Printf("The type of a, b, c, d is: %T, %T, %T, and %T respectively\n", a, b, c, d)
}
