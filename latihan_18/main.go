package main

import (
	"fmt"
	"sync"
)

var counter int = 0

/*
Fungsi goroutine ini akan menyebabkan race condition

yang dimana nilai counter akan tidak deterministik

Setiap goroutine berebutan untuk menambahkan nilai counter yang dimana ketika
loop selesai maka nilai akhirnya menjadi tidak terduga

oleh karena itu menggunakan sync.Mutex untuk mengunci variabel counter ketika
ditambahkan dan melepaskan kuncinya ketika selesai
*/
func RaceCondIncrement(wg *sync.WaitGroup) {
	counter++
	wg.Done()
}

func safeIncrement() {

	counter++

}

func main() {
	//var wg sync.WaitGroup
	var mut sync.Mutex
	expectedCounter := 1000

	for i := 0; i < expectedCounter; i++ {
		//wg.Add(1)
		//go RaceCondIncrement(&wg)
		go func() {
			mut.Lock() // mengunci segala data yang sedang diproses goroutine ini
			defer mut.Unlock()
			safeIncrement()
		}()
	}

	//wg.Wait()
	fmt.Println("Expected Counter:", expectedCounter)
	fmt.Println("Actual Counter:", counter)
	// Check for race condition
	if expectedCounter != counter {
		fmt.Println("Race condition detected!")
	} else {
		fmt.Println("No race condition detected.")
	}
}
