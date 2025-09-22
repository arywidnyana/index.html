package main

import (
    "fmt"
)

func main() {
    var umur int
    var buatkk bool

    fmt.Print("Masukkan umur: ")
    fmt.Scan(&umur)
    fmt.Print("Apakah bisa membuat kk (true/false): ")
    fmt.Scan(&buatkk)

    if umur >= 17 && buatkk {
        fmt.Println("bisa membuat KTP")
    } else {
        fmt.Println("belum bisa membuat KTP")
    }
}
