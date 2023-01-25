package main

import "fmt"

func main() {
    a := []string{"Alta3", "Research", "Loves", "GoLang"}

    for i, s := range a {
        fmt.Println("Position", i, "contains the string:", s)
    }

    for i := range a {
        fmt.Println("Position", i)
    }

    for pos := range a {
        fmt.Println("Position", pos)
    }
}
