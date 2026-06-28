package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func validateEmailFormat(e string) bool {

	match, err := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9._+-]+\\.[a-zA-Z]{2,}$", e)
	if err != nil {
		return false
	}
	return match
}

func validateNumberFormat(num string) bool {

	match, err := regexp.MatchString("^[0-9]{12}", num)

	if err != nil {
		return false
	}

	return match

}

func inputEmail(e string) (bool, string) {

	result := validateEmailFormat(e)

	if !result {

		return false, "incorrect email format"
	}

	return true, "correct email format"
}

func inputNumber(num string) (bool, string) {

	result := validateNumberFormat(num)

	if !result {

		return false, "incorrect number format"
	}

	return true, "correct number format"
}

func main() {

	state := true

	for state {
		email := bufio.NewScanner(os.Stdin)
		number := bufio.NewScanner(os.Stdin)
		fmt.Printf("Alamat Email : ")
		email.Scan()
		resEmail, msgEmailFormat := inputEmail(email.Text())

		fmt.Printf("No. Telp : ")
		number.Scan()
		resNumber, msgNumberFormat := inputNumber(number.Text())

		if resEmail && resNumber == true {
			break
		}

		fmt.Printf("%s,%s\n", msgEmailFormat, msgNumberFormat)
	}
}
