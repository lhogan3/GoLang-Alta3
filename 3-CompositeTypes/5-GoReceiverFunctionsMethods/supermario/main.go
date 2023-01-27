package main

import (
    "fmt"
    "github.com/lhogan3/supermario/models"
)

func main() {
    mario := models.Player{3, 1, []string{"mushroom"}}

    fmt.Println(mario.Lives)
    mario.Greenmushroom()
    fmt.Println(mario.Lives)

    fmt.Println(mario.Inventory)
    mario.Pickup("coin")
    fmt.Println(mario.Inventory)

    fmt.Println(mario.CanWhistle())
    mario.Stage = 11
    fmt.Println(mario.CanWhistle())
}
