package main 

import "fmt"


func main () {

    // constant type
    const (
        nasiGoreng = 120000
        mieGoreng = 130000
    )

    pesanan1 := 3 * nasiGoreng
    pesanan2 := 5 * mieGoreng

    fmt.Println(pesanan1)
    fmt.Println(pesanan2)

}