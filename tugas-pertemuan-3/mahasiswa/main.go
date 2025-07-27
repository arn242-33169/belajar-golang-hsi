package mahasiswa

var Versi string = "v1.0.0"

var maxNilai int = 100

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

func (m Mahasiswa) Info() string {
	return "Nama: " + m.Nama
}

func (m Mahasiswa) RataRata() float64 {
	return m.nilaiAvg
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}
