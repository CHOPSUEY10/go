package main
import "fmt"

func isDataLengthValid(data map[string]string) bool {

    if(len(data) < 2){
        return false
    }

    return true
}

func insertNewData(data map[int]string, name string) {
    
    data[len(data)+1] = name 
    
}

func deleteData(data map[int]string, id int) {

    delete(data,id)

}




func main(){ 

    // person := map[string]string{ 
    //     "nama" : "fadli",
    //     "jenis_kelamin" : "laki_laki",
    // }

    employee := map[int]string {
        1 : "Fadli Aidin",
        2 : "Muhammad Hafidz",
        3 : "Phillip Damian Samosir",
    }

    fmt.Println(len(employee))
    insertNewData(employee,"M. Ridho Faizal")
    //deleteData(employee,1)

    for i := 1; i <= len(employee) + 1; i++ {

        if employee[i] == "" {
            continue
        }
        fmt.Println("data karyawan baru : ", employee[i] )
    }

    //fmt.Println("validasi data : ",isDataLengthValid(person))
    
}
