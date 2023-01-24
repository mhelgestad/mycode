package main

import "fmt"

type author struct {
    name   string
    branch string
    books  int
    salary int
}

func (a author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch's Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

func (a *author) bookup() {
    a.books++
}

func main() {
    res := author {
        name: "MaxHelgestad",
        branch: "Minnesota",
        books: 14,
        salary: 34000,
    }

    res.show()
    
    res.bookup()

    res.show()
}
