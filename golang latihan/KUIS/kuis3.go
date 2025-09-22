package main

import (
    "fmt"
)

func main() {
    // Header
    fmt.Println("========================================")
    fmt.Println("       APLIKASI HITUNG DISKON           ")
    fmt.Println("========================================")

    var harga, diskon, total float64

    for {
        fmt.Print("\nMasukkan harga barang dalam rupiah (ketik 0 untuk selesai): ")
        fmt.Scan(&harga)
        if harga == 0 {
            break
        }

        fmt.Print("Masukkan diskon dalam persen (%): ")
        fmt.Scan(&diskon)

        hargaSetelahDiskon := harga * (1 - diskon/100)
        fmt.Println("----------------------------------------")
        fmt.Printf("Harga awal           : Rp %.2f\n", harga)
        fmt.Printf("Diskon (%v%%)       : Rp %.2f\n", diskon, harga*(diskon/100))
        fmt.Printf("Harga setelah diskon : Rp %.2f\n", hargaSetelahDiskon)
        fmt.Println("----------------------------------------")

        total += hargaSetelahDiskon
    }

    // Footer dengan total
    fmt.Println("\n========================================")
    fmt.Printf("Total belanja Anda = Rp %.2f\n", total)
    fmt.Println("========================================")
    fmt.Println("  Terima kasih telah menggunakan aplikasi!")
    fmt.Println("========================================")
}
