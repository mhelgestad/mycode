package main

import (
    "html/template"
    "io"
    "log"
    "os"
    "strings"
)

func parse(path string) {
    
    t, err := template.ParseFiles(path)

    if err != nil {
        log.Print(err)
        return
    }

    path = strings.TrimRight(path, ".template")

    f, err := os.Create(path)
    if err != nil {
        log.Println("create file: ", err)
        return
    }

    config := map[string]string {
        "textColor": "#abcdef",
        "linkColorHover": "#ffaacc",
    }

    err = t.Execute(f, config)
    if err != nil {
        log.Print("execute: ", err)
        return
    }

    f.Close()
}

func main() {
    parse("example.css.template")

    f, _ := os.Open("example.css")

    io.Copy(os.Stdout, f)

    f.Close()
}
