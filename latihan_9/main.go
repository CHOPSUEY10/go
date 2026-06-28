package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generateRandomNumber(input string) (result int, e error) {

	seed, err := strconv.Atoi(input)

	if err != nil {
		return result, e
	}

	r := rand.New(rand.NewSource(int64(seed)))

	result = r.Int() % 50

	return result, nil
}

func genSeriesRandNum(input ...int) (result []int, e error) {

	for _, v := range input {

		r := rand.New(rand.NewSource(int64(v)))
		result = append(result, r.Int()%100)

	}
	return result, e

}
func main() {

	// Membuat generator random

	var input string

	state := true
	for state {

		var menu string

		fmt.Printf("G : Generate angka\n")
		fmt.Printf("S : Generate rangkaian angka\n")
		fmt.Printf("Q : Keluar\n")
		fmt.Printf("\tMasukkan pilihan menu : ")
		fmt.Scanln(&menu)

		switch menu {
		case "G":
			fmt.Printf("Masukkan angka : ")
			fmt.Scanln(&input)
			result, _ := generateRandomNumber(input)
			fmt.Printf("Angka random : %d\n\n", result)
		case "Q":
			state = false
		case "S":
			seriesInput := bufio.NewScanner(os.Stdin)
			fmt.Printf("Masukkan rangkaian angka : ")
			seriesInput.Scan()

			fields := strings.Fields(seriesInput.Text())
			var numbers []int

			// mengubah seluruh rangkaian string angka menjadi integer dan append ke slice numbers
			for _, v := range fields {
				num, err := strconv.Atoi(v)

				if err != nil {
					fmt.Printf("Angka tidak valid")
					continue
				}

				numbers = append(numbers, num)

			}

			// Mengubah slice numbers menjadi variadic arguments
			result, _ := genSeriesRandNum(numbers...)
			var printed string
			for _, v := range result {

				char := strconv.Itoa(v)

				printed += char + ","

			}
			fmt.Printf("\nAngka random : %s\n\n", printed)
		default:

			fmt.Printf("Pilihan tidak valid\n\n")

		}

	}

}
