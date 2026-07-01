package utils

import "fmt"

var errMessage = map[string]string{

	"1": "Invalid permissions",
}

var successMessage = map[string]string{

	"1": "Berhasil memberikan izin menulis",
	"2": "Berhasil memberikan izin membaca",
	"3": "Berhasil memberikan izin mengeksekusi",
}

func ErrorHandler(msg error) {
	idx := msg.Error()
	fmt.Println(errMessage[idx])
}

func MessageHandler(msg string) {
	fmt.Println(successMessage[msg])
}
