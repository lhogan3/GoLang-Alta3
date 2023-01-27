package main

import (
    "fmt"
    "regexp"
)

var isEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+$`)
var isNumeric = regexp.MustCompile(`[0-9]`)
var isTwoCharacterString = regexp.MustCompile(`^[a-zA-Z]{2}$`)

func main() {
    values := []string{"thisisanemail@gmail.com", "12345", "PF", "thisisnotanemail@asdfasdfasdf", "a", "-2", "cd", "42"}

    for _, value := range values {
        switch {
            case isEmail.MatchString(value):
                fmt.Println(value, "is an email.")
            case isNumeric.MatchString(value):
                fmt.Println(value, "is numeric.")
            case isTwoCharacterString.MatchString(value):
                fmt.Println(value, "is a two character string.")
            default:
                fmt.Println(value, "is not an email, numeric, or a two character string.")
        }
    }
}
