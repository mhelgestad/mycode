package main

import "fmt"
import "os"

func main() {

    for i, v := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, v)
    }
}
