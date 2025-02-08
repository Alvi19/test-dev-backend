package models

type Pet struct {
	Species         string
	Breed           string
	Name            string
	Characteristics string
}

var Pets = []Pet{
	{"Kucing", "Persia", "Luna", "Lembut, manja"},
	{"Anjing", "Golden Retriever", "Max", "Setia, cerdas"},
	{"Kelinci", "Angora", "Bunny", "Lincah, lucu"},
}
