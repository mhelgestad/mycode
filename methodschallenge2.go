package main

import "fmt"

type Virtmach struct {
    ip string
    hostname string
    diskgb int
    ram int
}

func (v *Virtmach) setram(ram int) {
    v.ram = ram
}

func (v *Virtmach) addDisk(diskgb int) {
    v.diskgb = v.diskgb + diskgb
}

func (v Virtmach)display(){
    fmt.Println("Primary IP Address:", v.ip)    // primary IP address 
    fmt.Println("Hostname:", v.hostname)        // hostname 
    fmt.Println("Disk GB:", v.diskgb)           // diskgb 
    fmt.Println("RAM:", v.ram)                  // ram 
}

func main() {
    vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8}
    vm1.display()

    vm1.setram(16)
    vm1.addDisk(100)
    vm1.display()
}
