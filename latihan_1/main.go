package main

import "fmt"

func main() {

    // Define Variables 
    var nama string = "fadli aidin"
    var umur int = 23 
    var tinggi_badan float32 = 172.4
    var pelajar bool = true

    // Another way define variables 

    jenis_hewan := "mamalia"
    umur_hewan := 19
    tinggi_hewan := 23.12


    fmt.Println(nama,umur,tinggi_badan,pelajar)
    fmt.Printf("jenis hewan : %s\numur_hewan : %d\ntinggi_hewan %f", jenis_hewan, umur_hewan, tinggi_hewan)

}