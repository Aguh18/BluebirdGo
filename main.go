package main

import (
	"bluebird/ganjil"
	"bluebird/faktorial"
	"bluebird/suhu"
	"bluebird/palindrome"
	
	"fmt"
)

func main() {
	var pilihan int

	fmt.Println("1. Penghitung Bilangan Ganjiling")
	fmt.Println("2. Konversi Celcius ke Fahrenheit")
	fmt.Println("3. Cek Palindrome")
	fmt.Println("4. Faktorial")
	fmt.Print("Pilih fungsi\t:")

	_, err := fmt.Scanln(&pilihan)
	if err != nil {
		fmt.Println("Error:\t", err)
		return
		
	}
	if pilihan == 1 {
		ganjil.Ganjil()
		
	}else if pilihan == 2 {
		suhu.CelciusToFahrenheit()
	}else if pilihan == 3 {
		palindrome.Palindrome()
	}else if pilihan == 4 {
		faktorial.Faktorial()
	}else{
		fmt.Println("Pilihan tidak tersedia")
	}
	
}
