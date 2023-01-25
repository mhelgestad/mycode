package main

import (
    "fmt"
    "regexp"
)

var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+$`)
var shortPhone = regexp.MustCompile(`^[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]$`)
var longPhone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]$`)

func main() {
        contacts := []string{
                "(717) 566-4428",
                "rzfeeser@alta3.com",
                "566-4428",
                "123.555.4567",
                "4 bonnywick drive",
        }

        // loop across our slice of strings
        for _, contact := range contacts {
                switch {
                case email.MatchString(contact):
                        fmt.Println(contact, "is an email")
                case shortPhone.MatchString(contact):
                        fmt.Println(contact, "is a short phone number")
                case longPhone.MatchString(contact):
                        fmt.Println(contact, "is a long phone number")
                default:
                        fmt.Println(contact, "is not recognized")
                }
        }
}
