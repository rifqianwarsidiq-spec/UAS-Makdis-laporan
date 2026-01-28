package main

import "fmt"

// Struct untuk menyimpan data mahasiswa
type Mahasiswa struct {
	NIM  string
	Nama string
}

// Fungsi Sequential Search
func sequentialSearch(data []Mahasiswa, nim string) int {
	for i := 0; i < len(data); i++ {
		if data[i].NIM == nim {
			return i
		}
	}
	return -1
}

func main() {
	// Data mahasiswa (contoh)
	mahasiswa := []Mahasiswa{
		{"230411001", "Andi"},
		{"230411002", "Budi"},
		{"230411003", "Citra"},
	}

	var cariNIM string
	fmt.Print("Masukkan NIM yang dicari: ")
	fmt.Scanln(&cariNIM)

	// Proses pencarian
	index := sequentialSearch(mahasiswa, cariNIM)

	// Output hasil
	if index != -1 {
		fmt.Println("Data mahasiswa ditemukan")
		fmt.Println("NIM  :", mahasiswa[index].NIM)
		fmt.Println("Nama :", mahasiswa[index].Nama)
	} else {
		fmt.Println("Data mahasiswa tidak ditemukan")
	}
}
