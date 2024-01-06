package palindrome

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Palindrome() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	kalimat, _ := reader.ReadString('\n')

	kalimat = strings.TrimSpace(kalimat)
	kalimat = strings.ToLower(kalimat)

	kalimatBalik := ""
	for i := len(kalimat) - 1; i >= 0; i-- {
		kalimatBalik += string(kalimat[i])
	}

	if kalimat == kalimatBalik {
		fmt.Println("Kalimat adalah palindrome")
	} else {
		fmt.Println("Kalimat bukan palindrome")
	}
}