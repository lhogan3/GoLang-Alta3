package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    people := []string{"John", "Jacob", "Jingle", "Heimer", "Schmidt"}

    index := rand.Intn(len(people)-1)
    fmt.Println("random number:", index)

    person := people[index]

    if person == "John" {
        fmt.Println("John is the leader of the pack.")
    } else if person == "Jingle" {
        fmt.Println("What a musical guy!")
    } else {
        fmt.Println("Jacob, Heimer, and Schmidt are cool too.")
    }
}
