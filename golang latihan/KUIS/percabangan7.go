package main

import (
    "fmt"
)

func main() {
    var n int
    var total, uang int

    fmt.Print("Masukkan jumlah minggu: ")
    fmt.Scan(&n)

    fmt.Println("Masukkan jumlah uang yang dikumpulkan setiap minggu:")
    for i := 1; i <= n; i++ {
        fmt.Printf("Minggu %d: ", i)
        fmt.Scan(&uang)

        if uang < 0 {
            fmt.Println("Jumlah uang tidak boleh negatif")
        } else if uang == 0 {
            fmt.Println("Tidak ada uang yang dikumpulkan minggu ini")
        } else {
            total += uang
        }
    }

    fmt.Printf("Total uang yang dikumpulkan dalam sebulan: %d\n", total)
}
