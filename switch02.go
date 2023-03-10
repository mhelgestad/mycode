package main

import (
    "fmt"
    "time"
)

func main() {
    watch := time.Now()

    switch {
        case watch.Hour() < 6:
            fmt.Println("Go back to sleep.")
        case watch.Hour() < 12:
            fmt.Println("Good morning!")
        case watch.Hour() < 18:
            fmt.Println("Good Afternoon.")
        default:
            fmt.Println("Counting sheep. Z-z-z-z-z-z")
    }
}
