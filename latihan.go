package main

import "fmt"

//----------------- KALKULASI START ------------------------//

func kali(angka1, angka2 int) int {

	perkalian := angka1 * angka2
	return perkalian

}
func tambah(angka1, angka2 int) int {

	penjumlahan := angka1 + angka2
	return penjumlahan
}
func kurang(angka1, angka2 int) int {
	pengurangan := angka1 - angka2
	return pengurangan
}
func bagi(angka1, angka2 int) int {
	pembagian := angka1 / angka2
	return pembagian
}

//----------------- KALKULASI END ------------------------//

//----------------- UCAPAN START ------------------------//

type orang struct {
	namanya string
	umurnya int
}

func (p *orang) ucapan() {
	fmt.Println("Hallo ", p.namanya+",", "apakah benar anda berusia", p.umurnya, "tahun?")
}

//----------------- UCAPAN END ------------------------//

func main() {
	//Inisialisasi variabel
	a := 20
	b := 20
	c := 40
	d := 30
	e := 50
	f := 15
	g := 80
	h := 4

	hasilKali := fmt.Sprintf("Hasil dari %d x %d adalah : ", a, b)
	hasilTambah := fmt.Sprintf("Hasil dari %d + %d adalah : ", c, d)
	hasilKurang := fmt.Sprintf("Hasil dari %d - %d adalah : ", e, f)
	hasilBagi := fmt.Sprintf("Hasil dari %d / %d adalah : ", g, h)

	fmt.Println(hasilKali, kali(20, 20))
	fmt.Println(hasilTambah, tambah(40, 30))
	fmt.Println(hasilKurang, kurang(50, 15))
	fmt.Println(hasilBagi, bagi(80, 4))

	// Membuat garis batas start //
	fmt.Println("-------------------------------------")
	// Membuat garis batas end//

	var seseorang orang = orang{namanya: "Arifin", umurnya: 27}

	seseorang.ucapan()

}
