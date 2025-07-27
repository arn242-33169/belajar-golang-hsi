package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {
	m1 := mahasiswa.BuatMahasiswa("Ayu", 20, 80, 85, 90)
	m2 := mahasiswa.BuatMahasiswa("Budi", 22, 70, 75, 65)
	m3 := mahasiswa.BuatMahasiswa("Citra", 21, 95, 88, 92)

	daftar := []mahasiswa.Mahasiswa{m1, m2, m3}

	for _, mhs := range daftar {
		fmt.Println(mhs.Info())
		fmt.Printf("Nilai rata-rata: %.2f\n", mhs.RataRata())
		fmt.Printf("Nilai maksimal: %d\n", mahasiswa.GetMaxNilai(mhs.Nilai))
		fmt.Println()
	}

	fmt.Println("Versi package:", mahasiswa.Versi)

	hitungTotalUmur := func(data []mahasiswa.Mahasiswa) int {
		total := 0
		for _, m := range data {
			total += m.GetUmur()
		}
		return total
	}

	fmt.Println("Total umur seluruh mahasiswa:", hitungTotalUmur(daftar))
}
