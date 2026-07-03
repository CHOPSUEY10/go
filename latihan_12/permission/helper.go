package permission

import (
	"errors"
	"os"
	"strconv"
)

var interpretation = []string{"read", "write", "execute"}
var permissions = map[string]int64{"r": 4, "w": 2, "x": 1}
var addMsg = []string{"1", "2", "3"}
var rmMsg = []string{"4", "5", "6"}
var filelist = map[string]*FileData{}

type FileData struct {
	Nama string
	//permission value
	Pm   int64
	Size DataSize
}

type DataSize struct {
	Value float64
	Unit  string
}

// Operasi OR digunakan untuk menambah nilai
func addPermission(p int64, file *FileData) {

	state, err := hasPermission(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if !state {
		file.Pm |= p
		SuccessMsg(addMsg[p])
	}
	SuccessMsg(addMsg[p])
}

// Operasi XOR digunakan untuk mengurangi nilai
func removePermission(p int64, file *FileData) {

	state, err := hasPermission(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if !state {
		file.Pm ^= p
		SuccessMsg(rmMsg[p])
	}
	SuccessMsg(rmMsg[p])
}

// Operasi AND untuk mengecek nilai
func hasPermission(p int64, file *FileData) (bool, error) {

	if file.Pm >= 7 {
		return false, errors.New("1")
	}

	check := func(p int64, v int64) bool {
		return v&p == p

	}

	return check(p, file.Pm), nil

}

func parseToInt(i string) int64 {

	if len(i) != 3 {
		ErrorMsg(errors.New("4"))
	}
	parsed, err := strconv.ParseInt(i, 2, 64)

	if err != nil {
		ErrorMsg(errors.New("4"))
	}

	return parsed
}

func checkPermission(file *FileData) []string {

	binerString := strconv.FormatInt(int64(file.Pm), 2)
	var result = []string{}

	for i, v := range binerString {
		binary, err := strconv.ParseBool(strconv.QuoteRune(v))

		if err != nil {
			ErrorMsg(errors.New("3"))
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

func checkFile(i string) (os.FileInfo, error) {
	file, err := os.Stat(i)
	if err != nil {
		return nil, errors.New("3")
	}
	return file, nil
}

func addFile(i string) error {

	file, err := checkFile(i)
	if err != nil {
		return ErrorMsg(err)
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

	filelist[i] = &data
	printFileInfo(&data)

	return nil
}
