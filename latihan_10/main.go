package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	//"unicode"
)

var ciphertable = make(map[string]string)
var plaintable = make(map[string]string)

func shuffleAlphabet(table map[string]string) {

	var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	shuffled := make([]string, len(alphabet))
	copy(shuffled, alphabet)

	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for i := range alphabet {

		table[alphabet[i]] = shuffled[i]
		plaintable[shuffled[i]] = alphabet[i]
	}

}

func encrypt(text string) string {

	var hasil string = ""

	uppercase := strings.ToUpper(text)
	for c := 0; c < len(text); c++ {

		hasil += ciphertable[string(uppercase[c])]
	}

	return hasil

}

func decrypt(text string) string {

	uppercase := strings.ToUpper(text)
	var hasil strings.Builder
	for i := 0; i < len(text); i++ {

		hasil.WriteString(plaintable[string(uppercase[i])])
	}

	return hasil.String()
}

func mainMenu(shuffle func(map[string]string)) {
	shuffle(ciphertable)
	var option int

	state := true

	for state {
		fmt.Println("MONOALPHABETIC CIPHER")
		fmt.Printf("Masukkan pilihan menu (")
		fmt.Printf("1. Enkripsi 2. Dekripsi 3.Keluar) :")
		fmt.Scanln(&option)
		switch option {

		case 1:

			plaintext := bufio.NewScanner(os.Stdin)
			fmt.Println("\nMasukkan kata yang ingin di enkripsi : ")
			plaintext.Scan()

			e := encrypt(plaintext.Text())

			fmt.Printf("Hasil enkripsi adalah %s\n\n", e)

		case 2:

			ciphertext := bufio.NewScanner(os.Stdin)
			fmt.Println("\nMasukkan kata yang ingin di dekripsi : ")
			ciphertext.Scan()

			d := decrypt(ciphertext.Text())

			fmt.Printf("Hasil dekripsi adalah %s\n\n", d)

		case 3:
			fmt.Printf("terimakasih\n")
			state = false

		}
	}

}

func main() {
	mainMenu(shuffleAlphabet)
}
