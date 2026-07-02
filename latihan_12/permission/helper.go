package permission

import (
	"errors"
	"os"
	"strconv"
)

var interpretation = []string{"read", "write", "execute"}
var permissions = map[string]int{"r": 4, "w": 2, "x": 1}
var addMsg = map[string]string{"r": "1", "w": "2", "x": "3"}
var rmMsg = map[string]string{"r": "4", "w": "5", "x": "6"}

type FileData struct {
	Nama string
	//permission value
	Pm   int
	Size DataSize
}

type DataSize struct {
	Value float64
	Unit  string
}

// Operasi OR digunakan untuk menambah nilai
func addPermission(p string, file *FileData) {

	state, err := HasPermission(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if !state {
		file.Pm |= permissions[p]
		SuccessMsg(addMsg[p])
	}
	SuccessMsg(addMsg[p])
}

// Operasi XOR digunakan untuk mengurangi nilai
func removePermission(p string, file *FileData) {

	state, err := HasPermission(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if !state {
		file.Pm ^= permissions[p]
		SuccessMsg(rmMsg[p])
	}
	SuccessMsg(rmMsg[p])
}

// Operasi AND untuk mengecek nilai
func HasPermission(p string, file *FileData) (bool, error) {

	if file.Pm >= 7 {
		return false, errors.New("1")
	}

	check := func(p string, v int) bool {
		return v&permissions[p] == permissions[p]

	}

	return check(p, file.Pm), nil

}

func CheckPermission(file *FileData) []string {

	binerString := strconv.FormatInt(int64(file.Pm), 2)
	var result = []string{}

	for i, v := range binerString {
		binary, err := strconv.ParseBool(strconv.QuoteRune(v))

		if err != nil {
			ErrorMsg(err)
		}

		if binary {
			result[i] = interpretation[i]
		} else {
			result[i] = "x"
		}
	}

	return result

}

func formatSize(size int64) (value float64, unit string) {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		value = float64(size) / GB
		unit = "GB"
		return value, unit

	case size >= MB:
		value = float64(size) / MB
		unit = "MB"
		return value, unit

	case size >= KB:
		value = float64(size) / KB
		unit = "KB"
		return value, unit

	default:
		value = float64(size)
		unit = "B"
		return value, unit
	}
}
func addFile(i string) error {

	file, err := os.Stat(i)

	if err != nil {
		ErrorMsg(err)
	}

	v, u := formatSize(file.Size())
	data := FileData{
		Nama: i,
		Pm:   0,
		Size: DataSize{
			Value: v,
			Unit:  u,
		},
	}

	PrintFileInfo(&data)

	return nil
}
