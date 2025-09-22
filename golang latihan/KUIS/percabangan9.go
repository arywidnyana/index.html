package main

import (
    "fmt"
)

func main() {
    var jam, menit int
    var jarak float64

    fmt.Print("Masukkan waktu pemesanan (jam menit): ")
    fmt.Scan(&jam, &menit)
    fmt.Print("Masukkan jarak (km): ")
    fmt.Scan(&jarak)

    if jam < 5 || jam >= 22 || jarak > 20 {
        fmt.Println("Maaf, kami tidak bisa melayani pesanan Anda.")
        return
    }

    var tarifPerKm int
    if (jam >= 5 && jam <= 9) || (jam >= 16 && jam <= 19) {
        if jarak <= 10 {
            tarifPerKm = 5000
        } else {
            tarifPerKm = 4500
        }
    } else {
        tarifPerKm = 4000
    }

    harga := jarak * float64(tarifPerKm)
    fmt.Printf("Harga yang harus dibayar penumpang: Rp. %.0f\n", harga)
}
