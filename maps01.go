package main

import "fmt"

func main() {
    m:= make(map[string]int)

    m["key1"] = 400
    m["key2"] = 1000

    fmt.Println("map:", m)

    v1 := m["key1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "key2")

    fmt.Println("map:", m)

    fmt.Println(m["skeleton"])

    _, present := m["cookiemonster"]
    fmt.Println("present:", present)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map", n)

}
