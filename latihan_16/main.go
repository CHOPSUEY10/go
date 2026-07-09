package main

import (
	"antrean/customer"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func readString(token *bufio.Reader, prompt string) string {

	fmt.Print(prompt)
	res, _ := token.ReadString('\n')
	return strings.TrimSpace(res)

}

func catatOrderan(cust customer.Customer, orderQueue chan customer.Customer) {
	orderQueue <- cust
	fmt.Printf("Pelanggan %s masuk antrian\n", cust.Nama)
	time.Sleep(500 * time.Millisecond)
}

func prosesOrderan(orderQueue chan customer.Customer) customer.Customer {
	data := <-orderQueue
	fmt.Printf("Pesanan %s selesai diproses. Pesanannya adalah %s\n", data.Nama, data.Pesanan)
	fmt.Printf("\n")
	return data
}

func layaniPelanggan(pesanan []customer.Customer, orderQueue chan customer.Customer) []customer.Customer {
	hasil := make([]customer.Customer, 0, len(pesanan))

	for _, cust := range pesanan {
		catatOrderan(cust, orderQueue)
		hasil = append(hasil, prosesOrderan(orderQueue))
	}

	return hasil
}

func main() {
	cust := &customer.Customer{}
	pesanan := []customer.Customer{}
	reader := bufio.NewReader(os.Stdin)

	orderQueue := make(chan customer.Customer, 3)
	for {
		fmt.Printf("Menu : \n")
		fmt.Printf("a. Pelanggan Masuk \n")
		fmt.Printf("b. Layani Pelanggan\n")
		fmt.Printf("q. Keluar\n")

		option := strings.ToLower(readString(reader, "Pilihan : "))

		switch option {
		case "a":
			cust.Nama = readString(reader, "Nama Pelanggan : ")
			cust.Pesanan = readString(reader, "Pesanan : ")
			pesanan = append(pesanan, *cust)

		case "b":
			layaniPelanggan(pesanan, orderQueue)
			pesanan = []customer.Customer{}

		case "q":
			os.Exit(0)
		}
	}
}
