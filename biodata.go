package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	Program := os.Args
	idSiswa, _ := strconv.Atoi(Program[1])
	getDataKelas(idSiswa)
	// fmt.Println(namaProgram)
}

type namaSiswa struct {
	absensi   int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var dataKelas []namaSiswa

func getDataKelas(noabsen int) {
	dataKelas = []namaSiswa{
		{absensi: 1, nama: "ADITYA RIZKI PRATAMA", alamat: "Bandung ", pekerjaan: "Mahasiswa", alasan: "Alasan Aditya"},
		{absensi: 2, nama: "YULYANO THOMAS DJAYA", alamat: "Jakarta", pekerjaan: "Mahasiswa", alasan: "Alasan Yulyano"},
		{absensi: 3, nama: "ARIFINAL", alamat: "Aceh", pekerjaan: "Mahasiswa", alasan: "Alasan Arifinal"},
	}
	for _, siswa := range dataKelas {

		if siswa.absensi == noabsen {
			fmt.Println(siswa.nama)
			fmt.Println(siswa.alamat)
			fmt.Println(siswa.pekerjaan)
			fmt.Println(siswa.alasan)

		}
	}
}
