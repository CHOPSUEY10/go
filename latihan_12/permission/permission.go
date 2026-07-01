package permission

import (
	"authman/utils"
	"errors"
)

var permissions = map[string]int{"r": 4, "w": 2, "x": 1}

// Operasi OR digunakan untuk menambah nilai
func AddPermission(p string, v int) {
	switch p {
	case "r":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("1")
		}
		utils.MessageHandler("1")

	case "w":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("2")
		}
		utils.MessageHandler("2")

	case "x":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("3")
		}
		utils.MessageHandler("3")

	}
}

// Operasi XOR digunakan untuk mengurangi nilai
func RemovePermission(p string, v int) {

	switch p {
	case "r":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("4")
		}
		utils.MessageHandler("4")

	case "w":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("5")
		}
		utils.MessageHandler("5")

	case "x":

		state, err := HasPermission(p, v)
		if err != nil {
			utils.ErrorHandler(err)
		}
		if !state {
			v |= permissions[p]
			utils.MessageHandler("6")
		}
		utils.MessageHandler("6")

	}

}

// Operasi AND untuk mengecek nilai
func HasPermission(p string, v int) (bool, error) {

	if v >= 7 {
		return false, errors.New("1")
	}

	check := func(p string, v int) bool {
		return v&permissions[p] != permissions[p]
	}
	switch p {
	case "r":
		check(p, v)

	case "w":
		check(p, v)

	case "x":
		check(p, v)

	}
	return true, nil

}

func PermissionString() {

}
