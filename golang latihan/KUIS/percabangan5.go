package main

import (
    "fmt"
)

func main() {
    var jabatan string
    var masaKerja, jumlahAnak int
    var gajiPokok, tunjanganAnak int

    fmt.Print("Masukkan jabatan (Staf/Manager/Direktur): ")
    fmt.Scan(&jabatan)
    fmt.Print("Masukkan masa kerja (tahun): ")
    fmt.Scan(&masaKerja)
    fmt.Print("Masukkan jumlah anak: ")
    fmt.Scan(&jumlahAnak)

    if jabatan == "Staf" {
        if masaKerja < 5 {
            gajiPokok = 4000
            tunjanganAnak = 0
        } else if masaKerja > 10 {
            gajiPokok = 5000
            tunjanganAnak = 100 * jumlahAnak
        } else {
            gajiPokok = 4000
            tunjanganAnak = 100 * jumlahAnak
        }
    } else if jabatan == "Manager" {
        if masaKerja > 10 {
            gajiPokok = 10000
            tunjanganAnak = 300 * jumlahAnak
        } else {
            gajiPokok = 8500
            tunjanganAnak = 300 * jumlahAnak
        }
    } else if jabatan == "Direktur" {
        gajiPokok = 20000
        tunjanganAnak = 500 * jumlahAnak
    } else {
        fmt.Println("Jabatan tidak valid")
        return
    }

    if jumlahAnak > 3 {
        tunjanganAnak = 3 * (tunjanganAnak / jumlahAnak)
    }

    totalGaji := gajiPokok + tunjanganAnak

    fmt.Printf("Total gaji: Rp. %d\n", totalGaji)
}
