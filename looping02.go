package main

import "fmt"
import "math/rand"

func main() {
    
    drive := 0
    fmt.Print("Get a mulligan on any drive under 60 yards.\n")

    for ; drive <= 60; {
        drive = rand.Intn(251)
        fmt.Print("Swing!\n")
    }

    fmt.Println("Your longest drive was", drive, "yards")
}
