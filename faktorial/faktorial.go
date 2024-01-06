package faktorial

import (
	"fmt"
)

func Faktorial() {

	var inputAngka int // Deklarasi variable untuk menyimpan inputan user

	fmt.Print("Masukkan angka faktorial: ") // Menampilkan keterangan untuk memasukan angka yang akan di faktorialkan
	_, err := fmt.Scanln(&inputAngka)       // Menyimpan inputan user ke variable angka sekaligus menangkap error jika terjadi

	if err != nil {
		fmt.Println("Error:", err) // Jika terjadi error, maka program akan berhenti dan menampilkan error
		return
	}

	hasil := 1
	for i := 1; i <= inputAngka; i++ { // Melakukan perulangan dari 1 sampai angka yang diinputkan user
		hasil *= i // Menghitung faktorial dengan mengalikan setiap angka dari 1 hingga inputAngka
	}

	fmt.Println("Hasil faktorial:", hasil) // Menampilkan hasil faktorial ke layar
}
