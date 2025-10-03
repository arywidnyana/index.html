package main

import (
    "fmt"
)

func main() {
    // Menampilkan header
    fmt.Println("===================================")
    fmt.Println("       APLIKASI RATA-RATA SUHU      ")
    fmt.Println("===================================")

    var jumlahData int
    fmt.Print("Berapa data suhu yang ingin Anda inputkan: ")
    fmt.Scan(&jumlahData)

    var totalSuhu float64
    for i := 1; i <= jumlahData; i++ {
        fmt.Printf("\nMasukkan suhu dalam derajat Celcius (%d dari %d): ", i, jumlahData)
        var suhu float64
        fmt.Scan(&suhu)
        totalSuhu += suhu
    }

    // Menghitung rata-rata
    rataRataSuhu := totalSuhu / float64(jumlahData)

    // Menampilkan hasil dengan format rapi
    fmt.Println("\n===================================")
    fmt.Printf("Rata-rata suhu = %.2f derajat Celcius\n", rataRataSuhu)
    fmt.Println("===================================")
}
