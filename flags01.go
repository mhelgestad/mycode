package main

import (
    "flag"
    "fmt"
)

func main() {
    
    num := flag.Int("n", 5, "How angry is the Hulk? (# of iterations)")
    flag.Parse()

    n := *num

    i := 0
    for i < n {
        fmt.Println("Hulk Smash!")
        i++
    }
}
