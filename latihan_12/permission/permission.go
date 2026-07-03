package permission

import (
	"errors"
	"fmt"
	"os"
)

func Run(args ...any) {

	namefile, ok := args[0].(string)

	if !ok {
		ErrorMsg(errors.New("5"))
	}

	binary := args[1].(string)
	show, ok := args[2].(string)

	file := func(n string) bool {
		_, err := os.Stat(n)

		fmt.Println("Checking:", n)
		fmt.Println("Error:", err)

		if os.IsNotExist(err) {
			return false
		}
		return err == nil
	}

	if !file(namefile) {
		ErrorMsg(errors.New("3"))
	}

	pm := parseToInt(binary)
	f := filelist[namefile]

	if f.Pm >= pm {
		addPermission(pm, f)
	} else if f.Pm <= pm {
		removePermission(pm, f)
	}

	if show == "y" {
		printFileInfo(f)
	}

}
