package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Input Nama
	fmt.Print("Nama: ")
	namaInput, err := reader.ReadString('\n')
	if err != nil {
		logrus.WithError(err).Error("gagal membaca input nama")
		return
	}
	nama := strings.TrimSpace(namaInput)

	// Input Umur
	fmt.Print("Umur: ")
	umurInput, err := reader.ReadString('\n')
	if err != nil {
		logrus.WithError(err).Error("gagal membaca input umur")
		return
	}
	umurInput = strings.TrimSpace(umurInput)

	umur, err := strconv.Atoi(umurInput)
	if err != nil {
		err := errors.New("umur harus berupa angka")
		fmt.Printf(">> Error: %s\n", err.Error())
		return
	}

	// Validasi Umur
	if umur < 18 {
		err := errors.New("umur tidak valid (minimal 18 tahun)")
		fmt.Printf(">> Error: %s\n", err.Error())
		return
	}

	// Jika Valid
	fmt.Printf(">> Selamat datang, %s!\n", nama)
}
