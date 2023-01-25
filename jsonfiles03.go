package main

import (
    "encoding/json"
    "os"
)

func main() {
    type Person struct {
        Fn string
        Ln string
    }

    type ColorGroup struct {
        ID int
        name string
        colors []string
        p Person `json:"Person"`
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    group := ColorGroup{
        ID: 24601,
        name: "greens",
        colors: []string{"Moss", "Shamrock", "lime", "Hnter"},
        p: per,
    }

    b, err := json.Marshal(group)

    if err != nil {
        panic(err)
    }

    os.Stdout.Write(b)
}
