package main

import (
	"fmt"
)

func main() {
	var input int
	var totalPengunjung, jumlahHari, kenaikanHari int
	var pengunjungSebelumnya, pengunjungValid int
	data := []int{}

	for {
		fmt.Print("Masukkan jumlah pengunjung (masukkan angka negatif untuk berhenti): ")
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	if len(data) == 0 {
		fmt.Println("0 0")
		return
	}

	for i, pengunjung := range data {
		if pengunjung > 200 {
			fmt.Printf("%d tidak valid, tidak dihitung.\n", pengunjung)
			continue
		}

		totalPengunjung += pengunjung
		jumlahHari++

		if i > 0 && pengunjung > pengunjungSebelumnya {
			kenaikanHari++
		}
		pengunjungSebelumnya = pengunjung
		pengunjungValid++
	}

	rataRata := 0.0
	if pengunjungValid > 0 {
		rataRata = float64(totalPengunjung) / float64(jumlahHari)
	}

	fmt.Printf("%d %.2f\n", kenaikanHari, rataRata)
}