package main

import "fmt"

type Player struct {
    Lives int
    Stage int
    Inventory []string
}

func (p *Player) Greenmushroom() {
    p.Lives++
}

func (p *Player) Pickup(powerup string) {
    p.Inventory = append(p.Inventory, powerup)
}

func (p Player) CanWhistle() bool {
    return p.Stage >= 5
}

func main() {
    mario := Player{3, 1, []string{"mushroom"}}
    // display mario's current lives
    fmt.Println(mario.Lives)
    // mario just touched a greenmushroom!
    mario.Greenmushroom()
    // display mario's current lives have changed
    fmt.Println(mario.Lives)


    // display mario's current inventory
    fmt.Println(mario.Inventory)
    // mario just picked up a flower!
    mario.Pickup("flower")
    // display mario's current inventory has changed
    fmt.Println(mario.Inventory)


    // below stage 5 mario cannot use the warp whistle
    fmt.Println(mario.CanWhistle())
    // move to a stage where mario can use the warp whistle
    mario.Stage = 7
    // mario can only use the warp whistle on or above stage 5
    fmt.Println(mario.CanWhistle())
}


