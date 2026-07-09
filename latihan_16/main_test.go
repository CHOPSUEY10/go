package main

import (
	"testing"

	"antrean/customer"
)

func TestLayaniPelangganMempertahankanUrutan(t *testing.T) {
	pesanan := []customer.Customer{
		{Nama: "Fadli", Pesanan: "Kopi"},
		{Nama: "Andi", Pesanan: "Teh"},
		{Nama: "Siti", Pesanan: "Matcha"},
	}

	result := layaniPelanggan(pesanan, make(chan customer.Customer, len(pesanan)))

	if len(result) != len(pesanan) {
		t.Fatalf("jumlah hasil tidak sesuai: got %d want %d", len(result), len(pesanan))
	}

	for i, got := range result {
		if got.Nama != pesanan[i].Nama || got.Pesanan != pesanan[i].Pesanan {
			t.Fatalf("urutan salah di indeks %d: got %+v want %+v", i, got, pesanan[i])
		}
	}
}
