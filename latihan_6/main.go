package main 
import "fmt"



func main(){


	// Switch Case 
	// Merupakan jenis percabangan yang dimana setiap cabangnya bernilai setara 

	var input int 

	fmt.Printf("Masukkan angka : ")
	fmt.Scanf("%d",&input)
	var buah = [3]string{"nanas","mangga","semangka"} 


	switch input {
	case 1:
		fmt.Printf("ini buah %s", buah[input-1])
	case 2 : 
		fmt.Printf("ini buah %s", buah[input-1])
	case 3 : 		
		fmt.Printf("ini buah %s", buah[input-1])
	
	}
}