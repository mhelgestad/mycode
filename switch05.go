package main

import (
    "fmt"
)

func main() {
    nextstop := "Batcave"

    fmt.Println("Stops ahead of us:")

    switch nextstop {
        case "Superman":
            fmt.Println("Superman's Fortress of Solitude")
            fallthrough
        case "Batcave":
            fmt.Println("Batcave")
            fallthrough
        case "Strange":
            fmt.Println("Dr. Strange's Sanctum Sanctorum")
            fallthrough
        case "Justice":
            fmt.Println("Justice League's hall of Justice")
            fallthrough
        case "Xavier":
            fmt.Println("Xavier's School for gifted youngsters")
    }
}

