package suhu

// Mengimport package  yang dibutuhkan
import (
	"fmt"
	
)

func CelciusToFahrenheit() {
	var nilaiCelsius float32 // Deklarasi variabel untuk input suhu dalam format celcius
	fmt.Print("Masukkan nilai celcius:\t\t") // menampilkan ke layar perintah untuk memasukan nilai celcius
	_, err := fmt.Scanln(&nilaiCelsius)    // Menyimpan inputan user ke variable angka sekaligus menangkap error jika terjadi

	if err != nil {
		fmt.Println("Error:\t", err) // Jika terjadi error, maka program akan berhenti dan menampilkan error
		return
		
	}
	nilaiFahrenheit := (nilaiCelsius*9/5)+32 // mendeklarasikan sekaligus mengisi nilaiFahrenheit menggunakan rumus (c*9/5)+32

	fmt.Println("Nilai Fahrenheit:\t\t", nilaiFahrenheit) //Menampilkan nilai fahrenheit ke layar


}