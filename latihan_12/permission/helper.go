package permission

import (
	"errors"
	"os"
	"strconv"
)

var interpretation = []string{"read", "write", "execute"}
var addMsg = map[int64]string{4: "1", 2: "2", 1: "3"}
var rmMsg = map[int64]string{4: "4", 2: "5", 1: "6"}
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

	result := make([]string, len(interpretation))
	masks := []int64{4, 2, 1}

	for i, mask := range masks {
		if file.Pm&mask == mask {
			result[i] = interpretation[i]
		} else {
			result[i] = "-"
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
		return nil, errors.New("2")
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
