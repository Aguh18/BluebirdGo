package palindrome

// Mengimport package  yang dibutuhkan
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Palindrome() {
	reader := bufio.NewReader(os.Stdin) //Membuat object reader untuk membaca inputan dari user
	fmt.Print("Masukkan kalimat: ") // Menampilkan kalimat ke layar 
	kalimat, err := reader.ReadString('\n') // menampung hasil input ke variable kalimat sekaligus ke variable eror jika terjadi
	if err != nil {
		fmt.Println("Error:", err) // Menampilkan error sekaligus menghentikan progrram jika error terjadi
		return		
	}

	kalimat = strings.TrimSpace(kalimat) // Menghilangkan spasi agar cek palindrome bisa dilakukan dengan baik
	kalimat = strings.ToLower(kalimat) // Menjadikan semua huruf menjadi huruf kecil untuk menyamakan saat pengecekan nanti

	kalimatBalik := "" // Deklarasi variable untuk menampung kalimat yang sudah dibalik
	for i := len(kalimat) - 1; i >= 0; i-- {
		kalimatBalik += string(kalimat[i]) // merangkai kalimat secara terbalik merangkai satu persatu menggunakan looping 
	}

	if kalimat == kalimatBalik {
		fmt.Println("Kalimat adalah palindrome") // jika kalimat dan kalimat balik bernilai sama maka layar menamppilkan "kalimat adalah palindrome"
	} else {
		fmt.Println("Kalimat bukan palindrome") // jika kalimat dan kalimat bailik tidak bernilai sama maka layar akan menampilkan "kalimat bukan palindrome"
	}
}