package main

import (
	"testing"
)

// CalculateTriangles devuelve n-2 lados.
func CalculateTriangles(n int) int {
	if n < 3 {
		return 0
	}
	return n - 2
}

func TestCalculateTriangles(t *testing.T) {
	// 1. Pruebas Unitarias
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Triángulo", 3, 1},
		{"Cuadrado", 4, 2},
		{"Dodecágono", 12, 10},
		{"Caso Error", 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CalculateTriangles(tt.input)
			if res != tt.expected {
				t.Errorf("Falló %s: esperado %d, obtuvimos %d", tt.name, tt.expected, res)
			}
		})
	}
}

// Prueba de Estrés / Performance
func BenchmarkCalculateTriangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTriangles(1000)
	}
}
