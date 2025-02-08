package handlers

import (
	"esa_pets/utils"
	"fmt"
	"sort"
)

type Hewan struct {
	Jenis         string
	Ras           string
	Nama          string
	Karakteristik string
}

var hewanPeliharaan = []Hewan{
	{"Anjing", "Golden Retriever", "Otto", "Energik dan senang bermain bola"},
	{"Anjing", "Siberian Husky", "Max", "Bulu lebat dan mata biru"},
	{"Anjing", "Beagle", "Bob", "Ceria dan aktif bermain"},
	{"Kucing", "Persia", "Luna", "Anggun dan manja"},
	{"Kucing", "British Short Hair", "Milo", "Cerdas dan aktif"},
	{"Ikan", "Koi", "Nana", "Indah"},
	{"Ikan", "Mas", "Goldie", "Cerah"},
}

func TampilkanHewan() {
	fmt.Println("Daftar Hewan Peliharaan:")
	for _, h := range hewanPeliharaan {
		fmt.Printf("Jenis: %s, Ras: %s, Nama: %s, Karakteristik: %s\n", h.Jenis, h.Ras, h.Nama, h.Karakteristik)
	}
}

func TambahHewanPeliharaan(nama, ras, karakteristik string) {
	hewanPeliharaan = append(hewanPeliharaan, Hewan{"Badak", "Jawa", nama, karakteristik})
	fmt.Printf("Hewan peliharaan baru ditambahkan: Nama: %s, Ras: %s, Karakteristik: %s\n", nama, ras, karakteristik)
}

func HewanKesayangan(ascending bool) {
	var kesayangan []Hewan
	for _, h := range hewanPeliharaan {
		if h.Nama == "Otto" || h.Nama == "Luna" || h.Nama == "Milo" || h.Nama == "Max" {
			kesayangan = append(kesayangan, h)
		}
	}

	if ascending {
		sort.SliceStable(kesayangan, func(i, j int) bool {
			return kesayangan[i].Nama < kesayangan[j].Nama
		})
	} else {
		sort.SliceStable(kesayangan, func(i, j int) bool {
			return kesayangan[i].Nama > kesayangan[j].Nama
		})
	}

	fmt.Println("Hewan Kesayangan:")
	for _, h := range kesayangan {
		fmt.Printf("Nama: %s, Ras: %s, Karakteristik: %s\n", h.Nama, h.Ras, h.Karakteristik)
	}
}

func GantiKucing(namaLama, namaBaru, rasBaru string) {
	for i := 0; i < len(hewanPeliharaan); i++ {
		if hewanPeliharaan[i].Nama == namaLama {
			hewanPeliharaan[i].Nama = namaBaru
			hewanPeliharaan[i].Ras = rasBaru
			fmt.Printf("Kucing dengan nama %s telah diganti menjadi %s dengan ras %s\n", namaLama, namaBaru, rasBaru)
		}
	}
}

func JumlahHewanBerdasarkanJenis(jenis string) {
	count := 0
	for _, h := range hewanPeliharaan {
		if h.Jenis == jenis {
			count++
		}
	}
	fmt.Printf("Jumlah hewan dengan jenis %s: %d\n", jenis, count)
}

func CekPalindrome() {
	fmt.Println("Hewan dengan nama palindrome:")
	for _, h := range hewanPeliharaan {
		if utils.IsPalindrome(h.Nama) {
			fmt.Printf("%s adalah palindrome, panjang nama: %d\n", h.Nama, len(h.Nama))
		}
	}
}

func JumlahBilanganGenap(nums []int) {
	total := 0
	for _, num := range nums {
		if num%2 == 0 {
			total++
		}
	}
	fmt.Printf("Jumlah bilangan genap: %d\n", total)
}

func FormatJSON(inputPath, outputPath string) {
	utils.FormatJSON(inputPath, outputPath)
}
