package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type responSukses struct {
	Status string `json:"status"`
}

type responError struct {
	Error string `json:"error"`
}

var (
	errEmailKosong = errors.New("email tidak boleh kosong")
	errUmurKecil   = errors.New("umur harus 18 tahun atau lebih")
)

func handlerValidasi(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	umurString := r.URL.Query().Get("age")

	log.WithFields(logrus.Fields{
		"email": email,
		"age":   umurString,
	}).Info("Menerima permintaan validasi")

	// Parsing umur
	umur, err := strconv.Atoi(umurString)
	if err != nil {
		kirimError(w, fmt.Errorf("gagal mengubah age jadi integer: %w", err))
		return
	}

	// Validasi email
	if email == "" {
		kirimError(w, fmt.Errorf("validasi gagal: %w", errEmailKosong))
		return
	}

	// Validasi umur
	if umur < 18 {
		kirimError(w, fmt.Errorf("validasi gagal: %w", errUmurKecil))
		return
	}

	// Jika valid
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responSukses{Status: "ok"})
}

func kirimError(w http.ResponseWriter, err error) {
	log.WithError(err).Warn("Terjadi kesalahan validasi")

	pesan := "Terjadi kesalahan"
	switch {
	case errors.Is(err, errEmailKosong):
		pesan = "Email wajib diisi"
	case errors.Is(err, errUmurKecil):
		pesan = "Umur minimal 18 tahun"
	default:
		pesan = err.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(responError{Error: pesan})
}

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	http.HandleFunc("/validate", handlerValidasi)

	log.Info("Menjalankan server di http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.WithError(err).Fatal("Gagal menjalankan server")
	}
}
