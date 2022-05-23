package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	var result string //inisialiasasi hasil yang dikeluarkan dalama bentuk

	//case kondisi
	if nilai >= 80 && nilai <= 100 {
		result = "A"
	} else if nilai >= 65 && nilai <= 79 {
		result = "B+"
	} else if nilai >= 50 && nilai <= 64 {
		result = "B"
	} else if nilai >= 35 && nilai <= 49 {
		result = "C"
	} else if nilai >= 0 && nilai <= 34 {
		result = "D"
	} else {
		result = "Nilai invalid"
	}
	return result //menampilkan variable penampung

}

func main() {
	var nilai int = 0

	fmt.Println(KonversiNilai(nilai))
}
