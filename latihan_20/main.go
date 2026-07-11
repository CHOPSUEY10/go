package main

import (
	"fmt"
	"time"
)

func fetchPictures(chPictures chan<- string) {
	time.Sleep(time.Millisecond * 80)
	chPictures <- "Pictures"
	close(chPictures)
}

func fetchMusics(chMusics chan<- string) {
	time.Sleep(time.Millisecond * 40)
	chMusics <- "Musics"
	close(chMusics)
}

// for select pattern ini sama dengan waitgroup

func fetchData(done chan<- bool, chPictures <-chan string, chMusics <-chan string) {
	received := 0

	for received < 2 {
		select {
		case pic, ok := <-chPictures:
			if !ok {
				chPictures = nil
				continue
			}
			fmt.Println(pic)
			received++

		case mus, ok := <-chMusics:
			if !ok {
				chMusics = nil
				continue
			}
			fmt.Println(mus)
			received++
		}
	}

	done <- true
}

func main() {
	done := make(chan bool)
	chPictures := make(chan string)
	chMusics := make(chan string)

	go fetchPictures(chPictures)
	go fetchMusics(chMusics)
	go fetchData(done, chPictures, chMusics)

	//menunggu sinyal done agar dapat blocking main goroutine
	//tanpa set timer
	<-done
	fmt.Println("Semua file berhasil didownload")
}
