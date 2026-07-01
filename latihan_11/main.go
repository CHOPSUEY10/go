package main

import (
	"fmt"
	// Golden rule 3 : import berdasarkan nama package source code bukan nama file source code
	"kripto/gcd"
)

func main() {

	var a, b int = 40, 35

	var anon func(a, b int) int
	anon = func(a, b int) int {
		if b == 0 {
			return a
		}
		return anon(b, a%b)
	}

	fmt.Println(gcd.HitungGCD(a, b, anon))
}
