package permission

import (
	"fmt"
	"os"
)

var errorMessage = map[string]string{

	"1": "Invalid permissions",
	"2": "Can't find the find you've asked",
	"3": "Can't parse the permission of this file",
	"4": "Can't parse the permission argument",
	"5": "Please Input your file",
}

var successMessage = map[string]string{

	"1": "Berhasil memberikan izin membaca",
	"2": "Berhasil memberikan izin menulis",
	"3": "Berhasil memberikan izin mengeksekusi",
	"4": "Berhasil menghapus izin membaca",
	"5": "Berhasil menghapus izin menulis",
	"6": "Berhasil menghapus izin mengeksekusi",
}

func ErrorMsg(msg error) error {
	idx := msg.Error()
	message, ok := errorMessage[idx]
	if !ok {
		message = msg.Error()
	}

	fmt.Fprintln(os.Stderr, "Error:", message)
	os.Exit(1)
	return nil
}

func SuccessMsg(msg string) {
	fmt.Println(successMessage[msg])
}

func printCheckPermission(data []string) {

	fmt.Printf("Izin akses data : %s \n", data)

}

func printFileInfo(file *FileData) {

	pm := checkPermission(file)
	//fmt.Printf("File berhasil ditambahkan.\n")
	fmt.Printf("\tNama file : %s\n\tIzin file : %s\n\t Ukuran : %f %s\n", file.Nama, pm, file.Size.Value, file.Size.Unit)

}

func showManual() {

	fmt.Printf("AUTHMAN\n")
	fmt.Printf("----------------------\n")
	fmt.Printf("Manajer Otorisasi File\n\n")

	fmt.Printf("-f=\"[NAMA_FILE]\"\n")
	fmt.Printf("-p=\"[PERIZINAN]\"\n")
	fmt.Printf("\t Dalam bentuk binary. ex : 100 | 110 | 111  \n")
	fmt.Printf("-s=\"[FILE_INFO]\"\n")
	fmt.Printf("\t Informasi Izin file dan ukurannya \n")

}
