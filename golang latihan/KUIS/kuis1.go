package main

import (
	"fmt"
)

func main() {
	var batasBawah, batasAtas int

	// Header tampilan
	fmt.Println("====================================")
	fmt.Println("    PROGRAM PENGECEKAN GENAP/GANJIL")
	fmt.Println("====================================")

	// Meminta input batas bawah dan atas
	fmt.Print("\nMasukkan batas bawah: ")
	fmt.Scanln(&batasBawah)
	fmt.Print("Masukkan batas atas: ")
	fmt.Scanln(&batasAtas)

	// Divider
	fmt.Println("\nHasil Pengecekan:")
	fmt.Println("====================================")

	// Menampilkan hasil pengecekan genap atau ganjil
	for angka := batasBawah; angka <= batasAtas; angka++ {
		kategori := "Ganjil [FALSE]"
		if angka%2 == 0 {
			kategori = "Genap [TRUE]"
		}
		fmt.Printf("%d -> %s\n", angka, kategori)
	}

	