package mahasiswa

// hitungRataRata menghitung rata-rata dari sejumlah nilai
func hitungRataRata(nilai ...int) float64 {
	if len(nilai) == 0 {
		return 0
	}

	var total int
	for _, n := range nilai {
		total += n
	}
	return float64(total) / float64(len(nilai))
}

// BuatMahasiswa membuat objek Mahasiswa dengan nilai yang diberikan
func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	m := Mahasiswa{
		Nama:     nama,
		Nilai:    nilai,
		nilaiAvg: hitungRataRata(nilai...), // Hitung rata-rata dan simpan di struct
	}
	m.setUmur(umur) // Set umur menggunakan method private
	return m
}

// setUmur adalah method private untuk mengatur umur mahasiswa
func (m *Mahasiswa) setUmur(umur int) {
	m.umur = umur
}

// GetMaxNilai mengembalikan nilai tertinggi dari slice nilai
// Jika ada nilai yang melebihi batas maxNilai, maka dikembalikan maxNilai
func GetMaxNilai(nilai []int) int {
	if len(nilai) == 0 {
		return 0
	}
	max := nilai[0]
	for _, n := range nilai {
		if n > max {
			max = n
		}
	}
	if max > maxNilai {
		return maxNilai
	}
	return max
}
