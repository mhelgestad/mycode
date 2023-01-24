package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    
    rand.Seed(time.Now().UnixNano())

    if growth := rand.Intn(21 - -100) + -100; growth < 0 {
        fmt.Println("looks like growth was negative:", growth)
    } else if growth < 10 {
        fmt.Println("Growth was a big flat:", growth)
    } else {
        fmt.Println("Double digit growth:", growth)
    }
}
