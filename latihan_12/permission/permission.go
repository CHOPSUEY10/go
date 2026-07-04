package permission

import (
	"errors"
	"os"
)

// Operasi OR digunakan untuk menambah nilai
func addPermission(p int64, file *FileData) {

	state, err := checkChanges(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if state == file.Pm {
		file.Pm |= p
		SuccessMsg(addMsg[p])
	}
}

// Operasi AND NOT digunakan untuk mengurangi nilai
func removePermission(p int64, file *FileData) {

	state, err := checkChanges(p, file)
	if err != nil {
		ErrorMsg(err)
	}
	if state == file.Pm {
		file.Pm &^= p
		SuccessMsg(rmMsg[p])
	}
}

// Operasi AND untuk mengecek nilai
func checkChanges(p int64, file *FileData) (int64, error) {

	if file.Pm > 7 {
		return 0, errors.New("1")
	}

	check := func(p int64, v int64) int64 {
		p = p & v
		return p

	}

	return check(p, file.Pm), nil

}

func Run(args ...any) {
	if len(args) < 3 {
		showManual()
		os.Exit(0)
	}

	namefile, ok := args[0].(string)

	if !ok || namefile == "" {
		showManual()
		os.Exit(0)
	}

	binary, ok := args[1].(string)
	if !ok || binary == "" {
		showManual()
		os.Exit(0)
	}

	show, ok := args[2].(string)
	if !ok {
		show = ""
	}

	pm := parseToInt(binary)

	if _, ok := addMsg[pm]; !ok {
		defer showManual()
		ErrorMsg(errors.New("1"))
	}

	if _, ok := filelist[namefile]; !ok {
		if err := addFile(namefile); err != nil {
			return
		}
	}

	f := filelist[namefile]

	if state, err := checkChanges(pm, f); err != nil {
		ErrorMsg(err)
	} else if state < pm {
		addPermission(pm, f)
	} else {
		removePermission(pm, f)
	}

	if show == "y" {
		printFileInfo(f)
	}

}
