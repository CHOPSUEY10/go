package permission

import (
	"strings"
)

func AddPermission(p string, file *FileData) {
	p = strings.ToLower(p)
	addPermission(p, file)
}

func RemovePermission(p string, file *FileData) {

	p = strings.ToLower(p)
	removePermission(p, file)

}

func PermissionString() {

}
