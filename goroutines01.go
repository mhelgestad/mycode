package main

import (
    "fmt"
    "time"
)

func countDown(s int) {
    for i := s; i > 0; i-- {
        fmt.Println(i)
        time.Sleep(time.Second)
    }
}

func main() {
    fmt.Println("Nasa launch sequence starting:")

    go countDown(10)

    fmt.Println("launch sequence starting:")

    time.Sleep(time.Second)
    fmt.Println("Hey wait a second...")

    time.Sleep(time.Second)
    fmt.Println("I forgot my wallet!")

    fmt.Println("TAKE OFF!")
}

