// Online Go compiler to run Golang program online
// Print "Start small. Ship something." message

package main

import (
	"context"
	"fmt"
)

func main() {

	//c := uuid.NewString()

	// session := InitSession("2")
	// session.WriteSession()
	// sessions, err := ReadSession()
	// if err != nil {
	// 	fmt.Printf("Gagal membaca sesi server %v", err)
	// }

	// for _, v := range sessions {
	// 	fmt.Printf("Id : %s\n Token : %s\n", v.Id, v.Token)
	// }

	//fmt.Println(c)\

	session := &Session{}
	session.Token = "a5e579e6-7df8-4e69-b432-730e8a5a0178"
	session.Id = "2"

	ctx := context.WithValue(context.Background(), session, session.Id)

	// var falseSession = &Session{
	// 	Id:    "5",
	// 	Token: uuid.NewString(),
	// }

	if err := checkCookies(ctx, session, session.Token); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Authorized")

	}

}
