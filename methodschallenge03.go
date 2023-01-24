package main

import "fmt"

type Player struct{
    Lives int
    Stage int
    Inventory []string
}

func (p *Player)Greenmushroom(){
    p.Lives++
}

func (p *Player)Pickup(item string){
    p.Inventory = append(p.Inventory, item)
}

func (p Player)CanWhistle()bool{
    return p.Stage >= 5
}

func main() {
    mario := Player{3, 1, []string{"mushroom"}}

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
