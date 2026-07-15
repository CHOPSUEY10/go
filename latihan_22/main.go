package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {

	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {

	// 1. Validasi apakah koneksi nil
	if ws == nil {
		fmt.Println("Koneksi gagal: websocket.Conn bernilai nil")
		return
	}

	// 2. Validasi apakah nil pointer saat memanggil properti internal
	if ws.Config() == nil {
		fmt.Println("Koneksi masuk tanpa konfigurasi yang valid")
		return
	}

	fmt.Println("New incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true
	s.readLoop(ws)
}

func readText(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// Fungsi pembantu baru untuk membalas pesan secara asynchronous (non-blocking)
func (s *Server) handleReply(ws *websocket.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		res := readText(reader, "Balasan : ")
		if res != "" {
			_, err := ws.Write([]byte(res))
			if err != nil {
				fmt.Println("Gagal mengirim balasan:", err)
				break
			}
		}
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	// Jalankan pengirim balasan di latar belakang (Goroutine)
	// Jadi Anda bisa mengetik balasan kapan saja tanpa menghentikan proses baca pesan
	go s.handleReply(ws)

	buf := make([]byte, 2048)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client memutuskan koneksi.")
				break
			}
			fmt.Println("Read error:", err)
			break // Keluar loop jika ada error koneksi yang fatal
		}

		msg := buf[:n]
		// Cetak pesan dari client dengan format yang lebih rapi
		fmt.Printf("\nPesan dari client: %s\n", string(msg))
	}
}

func main() {

	fmt.Println("Menunggu koneksi...")
	server := NewServer()
	// Ganti bagian http.Handle sebelumnya dengan konfigurasi ini:
	wsServer := websocket.Server{
		Handler: websocket.Handler(server.handleWS),
		// Ini kuncinya! Fungsi ini akan mengabaikan error CORS/Origin
		Handshake: func(config *websocket.Config, req *http.Request) error {
			return nil
		},
	}

	http.Handle("/ws", wsServer)
	http.ListenAndServe(":3000", nil)
}
