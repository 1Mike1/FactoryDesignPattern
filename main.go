package main 

import (
  "fmt"
)

func main() {
    sam, _ := getMobile("samsung")
    app, _ := getMobile("apple")

    printDetails(sam)
    printDetails(app)
}

func printDetails(d Device) {
    fmt.Printf("Name: %s", d.getName())
    fmt.Printf("\nBrand: %s", d.getBrand())
    fmt.Printf("\nRAM: %s", d.getRAM())
    fmt.Printf("\nPrice: %s", d.getPrice())
    fmt.Println("\n**********************************************************************")
}
