package main

import "fmt"

//array and slice
func main() {

	// Deklarasi array

	var umur = [4]int{18, 19, 20, 21}

	// bisa juga seperti ini

	//age := [4]int{21,25,19,22}

	for i := 0; i < len(umur); i++ {

		fmt.Println(umur[i])
	}

	// Slice dan for range
	name := []string{"Fadli", "Doni", "Putri", "Tiara"}
	jenis_kelamin := make([]string, len(name), len(name)*2)
	jenis_kelamin[0] = "laki-laki"
	jenis_kelamin[1] = "laki-laki"
	jenis_kelamin[2] = "perempuan"
	jenis_kelamin[3] = "perempuan"

	for k, v := range name {
		fmt.Printf("%d: %s", k, v)
		fmt.Printf(":%s\n", jenis_kelamin[k])

	}

}
