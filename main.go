package main

import (
	"fmt"
	"net/http"
	"time"
)

// GetHandler adalah fungsi yang mengembalikan http.HandlerFunc
// yang menampilkan tanggal dan waktu dengan format yang diinginkan
func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Ambil waktu saat ini
		now := time.Now()

		// Format waktu sesuai instruksi
		formattedTime := fmt.Sprintf("%s, %d %s %d", now.Weekday(), now.Day(), now.Month(), now.Year())

		// Tulis hasil format waktu ke ResponseWriter
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(formattedTime))
	}
}

func main() {
	// Jalankan server di localhost:8080
	http.ListenAndServe("localhost:8080", GetHandler())
}
