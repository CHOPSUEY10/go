package main

import (
	"bufio"
	"errors"
	"fmt"
	"latihan_17/auth"
	"os"
	"strings"
	"sync"
	_ "time"

	_ "github.com/golang-jwt/jwt/v5"
)

type Account struct {
	UserID       string
	Password     string
	Salt         []byte
	LoginAttempt int
	Mut          sync.Mutex
}

// hashed, salt, err := auth.HashPassword(password)
func login(username, password string, account map[string]*Account) (error, string) {

	userAccount := account[username]
	if username != userAccount.UserID {

		return errors.New("User tidak ditemukan"), ""
	}

	if userAccount.LoginAttempt > 5 {

		return errors.New("Percobaan login telah lebih dari 5 kali"), ""
	}

	verified := auth.VerifyPassword(password, userAccount.Salt, []byte(userAccount.Password))

	if !verified {
		userAccount.Mut.Lock()
		defer userAccount.Mut.Unlock()
		userAccount.LoginAttempt += 1
		return errors.New("Your password is wrong"), ""
	}

	userAccount.Mut.Lock()
	defer userAccount.Mut.Unlock()
	userAccount.LoginAttempt = 0
	return nil, "Successfully login"

}

func register(username, password string, account map[string]*Account) (error, string) {

	hashed, salt, err := auth.HashPassword(password)

	if err != nil {
		return errors.New("Cannot generate hash or salt"), ""
	}

	account[username] = &Account{
		UserID:       username,
		Password:     string(hashed),
		Salt:         salt,
		LoginAttempt: 0,
	}
	return nil, "Successfully registered"

}

func readText(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {

	accountRepo := map[string]*Account{}
	reader := bufio.NewReader(os.Stdin)
	msg := make(chan string, 2)

	for {

		fmt.Printf("Autentikasi\n\n")
		fmt.Printf("a. Login\n")
		fmt.Printf("b. Register\n")
		choice := strings.ToLower(readText(reader, "Pilihan : "))

		switch choice {
		case "a":
			usrname := readText(reader, "Masukkan Username :")
			pass := readText(reader, "Masukkan Password : ")
			go func() {
				err, message := login(usrname, pass, accountRepo)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				msg <- message

			}()
			fmt.Println(<-msg)
		case "b":
			for {
				usrname := readText(reader, "Masukkan Username : ")
				pass := readText(reader, "Masukkan Password : ")
				verifPass := readText(reader, "Ulangi Password : ")

				IsSame := func(n, m string) bool {

					return n == m
				}(pass, verifPass)

				if IsSame {
					go func() {
						err, message := register(usrname, pass, accountRepo)
						if err != nil {
							fmt.Printf("%v", err)
						}
						msg <- message
					}()

					break
				}
				fmt.Println(<-msg)
			}

		case "q":
			close(msg)
			return

		}

	}
}
