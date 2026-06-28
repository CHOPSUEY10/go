package main 
import "fmt"
import "math"

func main() {


    // typecasting
    // Hati - hati, jika melebih melebihi atau kurang  nilai casting
    // maka akan terjadi overflow / underflow 

    //contoh overflow
    var kapasitasGalonBesar int32 = 32768
    var kapasitasGalonKecil int16 = int16(kapasitasGalonBesar)
    fmt.Println(kapasitasGalonKecil)
    
    var batasMinimalX int32 = -32769 
    var batasMinimalY int16 = int16(batasMinimalX)
    fmt.Println(batasMinimalY)
    
    // contoh underflow
    var nilaiDesimal64 float64 = math.SmallestNonzeroFloat32 * 2
    var nilaiDesimal32Maks float32 = float32(nilaiDesimal64)
    fmt.Printf("%f",nilaiDesimal32Maks)

    // string typecasting 
    var nama = "fadli aidin"
    // karakter pada string merupakan byte jadi untuk menampilkannya perlu typecasting
    var charByte = nama[3]
    var char = string(charByte)

    fmt.Printf("\n%s",char)

}