package main

import (
    "fmt"
)

func main() {
    var hasil [5]string
    fmt.Println("Masukkan lima hasil pertandingan (menang/kalah/draw):")
    for i := 0; i < 5; i++ {
        fmt.Scan(&hasil[i])
    }

    if hasil[0] == "kalah" && hasil[1] == "kalah" && hasil[2] == "kalah" && hasil[3] == "kalah" && hasil[4] == "kalah" {
        fmt.Println("dipecat")
    } else {
        fmt.Println("tidak dipecat")
    }
}
