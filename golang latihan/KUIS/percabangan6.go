package main

import (
    "fmt"
)

func main() {
    var gigi int
    var kopling, gas bool

    fmt.Print("Masukkan posisi gigi (0-4): ")
    fmt.Scan(&gigi)
    fmt.Print("Apakah kopling sedang ditarik? (true/false): ")
    fmt.Scan(&kopling)
    fmt.Print("Apakah gas menyala? (true/false): ")
    fmt.Scan(&gas)

    if gigi == 0 {
        if kopling && gas {
            fmt.Println("Mesin menyala dan motor tidak berjalan")
        } else {
            fmt.Println("Mesin mati")
        }
    } else if kopling {
        fmt.Println("Mesin menyala dan motor tidak berjalan")
    } else if gas {
        fmt.Println("Motor berjalan")
    } else {
        fmt.Println("Mesin mati")
    }

}
