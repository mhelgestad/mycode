package main

import "os"

func main() {
    
    _, err := os.Create("/carolDenvers/msmarvel")

    if err != nil {
        panic(err)
    }
}
