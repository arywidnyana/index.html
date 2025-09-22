package main

import "fmt"

func main() {
	var n, lebar, max int

	fmt.Print("Masukkan jumlah daun: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Print("Masukkan lebar masing-masing daun: ")

		fmt.Scan(&max)

		for i := 1; i < n; i++ {
			fmt.Scan(&lebar)
			if lebar > max {
				max = lebar // Update nilai maksimum
			}
		}

		fmt.Println("Lebar daun yang paling lebar adalah:", max)
	} else {
		fmt.Println("Jumlah daun harus lebih dari 0.")
	}
}