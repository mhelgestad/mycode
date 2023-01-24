package main

import (
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().Unix())
}

func main() {
    namesSlice := []string{"max", "jacob", "mackenzie", "jo"}

    if x := namesSlice[rand.Intn(len(namesSlice))]; x == "max" {
        fmt.Println("max works out of the NorthEast")
    } else if x == "jo" {
        fmt.Println("jo lives in the Pacific NorthWest")
    } else {
        fmt.Println("Hmm, I don't know much about this name", x)
    }
}
