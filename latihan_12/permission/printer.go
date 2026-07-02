package permission

import "fmt"

var errorMessage = map[string]string{

	"1": "Invalid permissions",
	"2": "Can't find the find you've asked",
	"3": "Can't parse the permission of this file",
}

var successMessage = map[string]string{

	"1": "Berhasil memberikan izin membaca",
	"2": "Berhasil memberikan izin menulis",
	"3": "Berhasil memberikan izin mengeksekusi",
}

func ErrorMsg(msg error) {
	idx := msg.Error()
	fmt.Println(errorMessage[idx])
}

func SuccessMsg(msg string) {
	fmt.Println(successMessage[msg])
}

func PrintCheckPermission(data []string) {

	fmt.Printf("Izin akses data : %s \n", data)

}

func PrintFileInfo(file *FileData) {

	pm := CheckPermission(file)
	fmt.Printf("File berhasil ditambahkan.\n")
	fmt.Printf("\tNama file : %s\n\tIzin file : %s\n\t Ukuran : %f %s", file.Nama, pm, file.Size.Value, file.Size.Unit)

}
