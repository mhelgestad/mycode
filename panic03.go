package main

import "fmt"

func main() {
    magic := make(map[string]string)

    magic["forests"] = "green"
    magic["mountains"] = "red"
    magic["plains"] = "white"
    magic["seas"] = "blue"
    magic["swamp"] = "black"

    fmt.Println(magic[100])
}
