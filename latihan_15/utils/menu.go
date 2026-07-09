package utils

import (
	"bufio"
	"crossroads/intersections"
	"crossroads/trafficlight"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readString(token *bufio.Reader, prompt string) string {

	fmt.Print(prompt)
	res, _ := token.ReadString('\n')
	return strings.TrimSpace(res)

}

func readInt(token *bufio.Reader, prompt string) int {

	for {
		text := readString(token, prompt)
		val, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Input harus berupa angka.")
		} else {
			return val
		}
	}

}
func Menu() {

	queue := &intersections.Queue{}
	intersect := &intersections.Intersections{}
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("TRAFFIC LIGHT SIMULATOR\n")
		fmt.Printf("Implementasi Round robin dan Goroutines\n")

		fmt.Printf("Menu : \n")
		fmt.Printf("a. Tambah lampu lalu lintas \n")
		fmt.Printf("b. Hapus lampu lalu lintas  \n")
		fmt.Printf("c. Mulai Simulasi \n")
		fmt.Printf("q. Keluar\n")

		option := strings.ToLower(readString(reader, "Pilihan : "))

		switch option {

		case "a":

			newTraffic := new(trafficlight.TrafficLight)
			newTraffic.RoadName = readString(reader, "Masukkan Pilihan Nama Jalan : ")
			newTraffic.Time = readInt(reader, "Masukkan Waktu Tunggu : ")

			intersect.AddTrafficLight(newTraffic)

		case "b":

			name := readString(reader, "Masukkan Nama Jalan Yang Ingin Dihapus : ")
			intersect.RemoveTrafficLight(intersect.Intersect[name])

		case "c":
			intersect.SwitchTrafficLight(queue)

		case "q":
			return
		}

	}

}
