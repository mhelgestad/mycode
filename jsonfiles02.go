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
        Id     int
        Name   string
        Colors []string
        P      Person `json:"Person"`
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    group := ColorGroup {
        Id: 24601,
        Name: "Greens",
        Colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
        P: per,
    }

    b, err := json.Marshal(group)

    if err != nil {
        panic(err)
    }

    os.Stdout.Write(b)
}
