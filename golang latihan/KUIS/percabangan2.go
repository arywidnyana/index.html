package main

import (
    "fmt"
)

func main() {
    var beratGr int
    fmt.Print("Masukkan berat parcel dalam gr: ")
    fmt.Scan(&beratGr)

    beratKg := beratGr / 1000
    sisaGr := beratGr % 1000

    biaya := beratKg * 10000
    tambahan := 0

    if beratKg < 10 {
        if sisaGr >= 500 {
            tambahan = sisaGr * 5
        } else {
            tambahan = sisaGr * 15
        }
    }

    total := biaya + tambahan

    fmt.Printf("%d kg + %d gr\n", beratKg, sisaGr)
    fmt.Printf("Rp. %d + Rp. %d = Rp. %d\n", biaya, tambahan, total)
}
