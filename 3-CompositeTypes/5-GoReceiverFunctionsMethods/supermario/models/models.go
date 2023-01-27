package models

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
