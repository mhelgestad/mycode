package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

type nasaMission struct {
    people  []Astro
    number  int
    message string
}
    
func main() {
    astro1 := Astro{"max", 24, "moon", true}
    astro2 := Astro{"jacob", 42, "mars", true}
    astro3 := Astro{"helgestad", 12, "jupiter", false}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

    p := []Astro{astro1, astro2, astro3}

    fmt.Println(p)

    fmt.Println(p[2].mission)

    s := nasaMission{p, 3, "success"}

    fmt.Println(s)
    fmt.Printf("%+v", s)
}
