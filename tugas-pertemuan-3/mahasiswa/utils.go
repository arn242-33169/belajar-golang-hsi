package mahasiswa

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

func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	m := Mahasiswa{
		Nama:     nama,
		Nilai:    nilai,
		nilaiAvg: hitungRataRata(nilai...),
	}
	m.setUmur(umur)
	return m
}

func (m *Mahasiswa) setUmur(umur int) {
	m.umur = umur
}

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
