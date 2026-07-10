package main

import (
	"fmt"
	"time"
)

//  Channel
//  Channel adalah tipe data khusus di Go-lang yang
//  memungkinkan dua goroutines bertukar data

// Write & Close only channel parameter; Artinya channel ini hanya bisa digunakan
// Dalam goroutine untuk menulis data dan menutup channel jika sudah selesai digunakan
func tambah10(n int, buff chan<- int) {

	n += 10
	buff <- n
	buff <- n

	close(buff)
}

// Read only channel parameter; Artinya channel ini hanya bisa digunakan untuk membaca data
// Channel yang telah ditutup, nilainya masih bisa dibaca oleh goroutine / variabel lainnya
func kali2(buff <-chan int) (n int) {

	n = <-buff
	return n * 2

}




func main() {

}
