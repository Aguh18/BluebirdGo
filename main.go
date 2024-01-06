package main

// Mengimport package  yang dibutuhkan
import (
	"bluebird/faktorial"
	"bluebird/ganjil"
	"bluebird/palindrome"
	"bluebird/suhu"

	"fmt"
)

func main() {
	var pilihan int // deklarasi variabel yang akan digunakan untuk memilih menu

	// menampilkan Menu
	fmt.Println("1. Penghitung Bilangan Ganjil")
	fmt.Println("2. Konversi Celcius ke Fahrenheit")
	fmt.Println("3. Cek Palindrome")
	fmt.Println("4. Faktorial")
	fmt.Print("Pilih fungsi\t:")

	_, err := fmt.Scanln(&pilihan) // memasukan nilai ke variabel pilihan sekaligus error jika terjadi
	if err != nil {
		fmt.Println("Error:\t", err) // menampilkan dan menghentikan program jika error terjadi
		return

	}
	if pilihan == 1 {
		ganjil.Ganjil() // jika pilihan 1 maka akan memanggil fungsi ganjil

	} else if pilihan == 2 {
		suhu.CelciusToFahrenheit() // jika pilihan 2 maka akan memanggil fungsi Ce;ciusToFahrenheit
	} else if pilihan == 3 {
		palindrome.Palindrome() //jika pilihan 3 maka akan memanggil fungsi palindrome
	} else if pilihan == 4 {
		faktorial.Faktorial() // jika pilihan 4 maka akan memanggil fungsi faktorial
	} else {
		fmt.Println("Pilihan tidak tersedia") // jika pilihan tidak tersedia maka akan menampilkan pesan ini
	}

}
