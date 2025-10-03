package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&input)

	var binary [32]int
	i := 0

	for input > 0 {
		binary[i] = input % 2
		input = input / 2     
		i++
	}

	fmt.Print("Hasil konversi ke biner: ")
	for j := i - 1; j >= 0; j-- {
		fmt.Print(binary[j])
	}
	fmt.Println()
}