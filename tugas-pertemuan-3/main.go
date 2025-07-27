package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {
	// Membuat tiga objek Mahasiswa dengan data nama, umur, dan nilai
	m1 := mahasiswa.BuatMahasiswa("Ayu", 20, 80, 85, 90)
	m2 := mahasiswa.BuatMahasiswa("Budi", 22, 70, 75, 65)
	m3 := mahasiswa.BuatMahasiswa("Citra", 21, 95, 88, 92)

	// Menggabungkan mahasiswa ke dalam slice
	daftar := []mahasiswa.Mahasiswa{m1, m2, m3}

	// Menampilkan informasi masing-masing mahasiswa
	for _, mhs := range daftar {
		fmt.Println(mhs.Info())
		fmt.Printf("Nilai rata-rata: %.2f\n", mhs.RataRata())
		fmt.Printf("Nilai maksimal: %d\n", mahasiswa.GetMaxNilai(mhs.Nilai))
		fmt.Println()
	}

	// Menampilkan versi dari package mahasiswa
	fmt.Println("Versi package:", mahasiswa.Versi)

	// Fungsi lokal untuk menghitung total umur seluruh mahasiswa
	hitungTotalUmur := func(data []mahasiswa.Mahasiswa) int {
		total := 0
		for _, m := range data {
			total += m.GetUmur()
		}
		return total
	}

	fmt.Println("Total umur seluruh mahasiswa:", hitungTotalUmur(daftar))
}
