package ganjil

// Mengimport package  yang dibutuhkan
import (
	"fmt"
	"log"
)

// Fungsi main, fungsi yang akan dijalankan ketika file di run
func Ganjil() {

	var batasAtas int // Deklarasi variable yang akan digunakan untuk batas atas sekaligus digunakan untuk pengulangan bertipe integer

	fmt.Print("Masukkan batas atas: ") // Menampilkan pesan ke layar untuk meminta input dari user

		_, err := fmt.Scanln(&batasAtas)   // Menyimpan inputan user ke variable angka sekaligus menangkap error jika terjadi
	if err != nil {
		log.Fatal(err) // Jika terjadi error, maka program akan berhenti dan menampilkan error
		return
	}
	for i := 1; i <= batasAtas; i++ { // Melakukan perulangan dari 1 sampai angka yang diinputkan user
		if i%2 != 0 { // Jika i modulus 2 tidak sama dengan 0, maka i adalah bilangan ganjil
			fmt.Print(i, " ") // Menampilkan bilangan ganjil ke layar
		}
	}

}
