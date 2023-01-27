package main

import "fmt"

type Virtmach struct {
    ip int
    hostname string
    diskgb int
    ram int
}

func (v Virtmach) Show() {
    fmt.Println("IP:", v.ip)
    fmt.Println("Host Name:", v.hostname)
    fmt.Println("Disk GB:", v.diskgb)
    fmt.Println("RAM:", v.ram)
}

func (v *Virtmach) Updatehostname(hostname string){
    v.hostname = hostname
}

func (v *Virtmach) Increaseram(){
    v.ram += 8
}

func main() {
    v := Virtmach{12345, "Liam's Host", 256, 8}
    v.Show()

    v.Updatehostname("Different Host Now")
    v.Show()

    v.Increaseram()
    v.Show()
}
