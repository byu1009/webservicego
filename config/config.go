package config

import (
	"log"
	"os"
	"strconv"
)

var EstimasiLayan int

func LoadConfig() {
	val := os.Getenv("ESTIMASI_LAYAN")
	if val == "" {
		log.Fatal("ENV ESTIMASI_LAYAN belum diset")
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("ENV ESTIMASI_LAYAN harus angka")
	}

	EstimasiLayan = n
	log.Println("ESTIMASI_LAYAN =", EstimasiLayan)
}
