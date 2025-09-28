package main

import (
    "fmt"
)

func main() {
    var angka int
    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&angka)

    if angka%3 == 0 && angka%5 == 0 {
        fmt.Println("Kelipatan 3 Kelipatan 5")
    } else if angka%3 == 0 {
        fmt.Println("Kelipatan 3")
    } else if angka%5 == 0 {
        fmt.Println("Kelipatan 5")
    }
}
