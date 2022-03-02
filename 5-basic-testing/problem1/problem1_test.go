package problem1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 1 + 2 = 3")
	}
}

// func TestSubstract(t *testing.T) {
// 	// Kondisi salah
// 	if Substract(2, 1) > 0 {
// 		t.Error("Salah, harusnya hasil 1 - 1= 0")
// 	}
// 	// if Substract(2, 1) == -1 {
// 	// 	t.Error("Salah, harusnya hasil 2 - 1 = 1")
// 	// }
// }

func TestGanjilGenap(t *testing.T) {
	t.Run("Cek Genap Sukses", func(t *testing.T) {
		if GanjilGenap(2) == "Genap" {
			t.Log("Hasil pengujian genap benar")
		}
	})
	// if GanjilGenap(3) == "Ganjil" {
	// 	t.Log("Hasil pengujian ganjil benar")
	// }
	if GanjilGenap(4) != "Ganjil" {
		t.Log("Hasil pengujian tidak sama dengan ganjil, sukses")
	}
	t.Run("Cek 0", func(t *testing.T) {
		if GanjilGenap(0) == "Genap" {
			t.Log("Hasil pengujian tidak sama dengan ganjil, sukses")
		}
	})
	t.Run("Cek Ganjil Sukses", func(t *testing.T) {
		assert.Equal(t, "Ganjil", GanjilGenap(5), "Hasil tidak sesuai")
		assert.NotEqual(t, "Genap", GanjilGenap(5), "Hasil tidak sesuai")
	})
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "Ganjil", GanjilGenap(5), "Hasil tidak sesuai")
	})
	t.Run("Case 2", func(t *testing.T) {
		assert.Equal(t, "Genap", GanjilGenap(6), "Hasil tidak sesuai")
	})
	t.Run("Case 3", func(t *testing.T) {
		assert.Equal(t, "Genap", GanjilGenap(10), "Hasil tidak sesuai")
	})
	t.Run("Case 4", func(t *testing.T) {
		assert.Equal(t, "Genap", GanjilGenap(2), "Hasil tidak sesuai")
	})
	t.Run("Case 5", func(t *testing.T) {
		assert.Equal(t, "Ganjil", GanjilGenap(1), "Hasil tidak sesuai")
	})
}
