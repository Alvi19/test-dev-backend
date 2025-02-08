package main

import (
	"esa_pets/handlers"
)

func main() {
	// Tampilkan semua hewan
	handlers.TampilkanHewan()

	// Tambah hewan peliharaan baru
	handlers.TambahHewanPeliharaan("Bobby", "Labrador", "Ramah dan pintar")

	// Tampilkan daftar setelah penambahan
	handlers.TampilkanHewan()

	// Tampilkan hewan kesayangan dalam urutan menaik
	handlers.HewanKesayangan(true)

	// Ganti nama kucing
	handlers.GantiKucing("Luna", "Mochi", "Scottish Fold")

	// Hitung jumlah hewan berdasarkan jenis
	handlers.JumlahHewanBerdasarkanJenis("Anjing")

	// Cek nama palindrome
	handlers.CekPalindrome()

	// Hitung jumlah bilangan genap
	nums := []int{2, 3, 4, 7, 10, 13, 16}
	handlers.JumlahBilanganGenap(nums)

	handlers.FormatJSON("assets/json/case.json", "assets/json/expectation.json")
}
