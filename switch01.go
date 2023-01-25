package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    var gover string = runtime.Version()

    switch {
        case strings.Contains(gover, "go1.19"):
            fmt.Println("Released in 2022")
        case strings.Contains(gover, "go1.18"):
            fmt.Println("You are using the latest version of GoLang")
        case strings.Contains(gover, "go1.17"):
            fmt.Println("This version of Go is fine")
        case strings.Contains(gover, "go1.16"):
            fmt.Println("You are using an older, but acceptable version of GoLang")
        default:
            fmt.Println("I cannot make a reccomendation")
    }
}
