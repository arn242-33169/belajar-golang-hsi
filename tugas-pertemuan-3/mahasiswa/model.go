package mahasiswa

var Versi string = "v1.0.0"

// maxNilai digunakan sebagai batas maksimum nilai
var maxNilai int = 100

// Mahasiswa merepresentasikan data mahasiswa
type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

// Info mengembalikan informasi nama mahasiswa
func (m Mahasiswa) Info() string {
	return "Nama: " + m.Nama
}

// RataRata mengembalikan nilai rata-rata yang sudah dihitung
func (m Mahasiswa) RataRata() float64 {
	return m.nilaiAvg
}

// GetUmur mengembalikan umur mahasiswa
func (m Mahasiswa) GetUmur() int {
	return m.umur
}
