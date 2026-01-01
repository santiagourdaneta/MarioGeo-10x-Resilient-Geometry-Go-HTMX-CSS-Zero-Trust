package main

import "testing"

func TestGetPolygonName(t *testing.T) {
	tests := []struct {
		sides    int
		expected string
	}{
		{3, "Triángulo"},
		{4, "Cuadrilátero"},
		{5, "Pentágono"},
		{10, "Polígono Complejo"},
	}

	for _, tt := range tests {
		result := getPolygonName(tt.sides)
		if result != tt.expected {
			t.Errorf("Para %d lados esperábamos %s, obtuvimos %s", tt.sides, tt.expected, result)
		}
	}
}

func TestGeometryLogic(t *testing.T) {
	sides := 5
	triangles := sides - 2
	if triangles != 3 {
		t.Errorf("Error en lógica geométrica: esperábamos 3 triángulos, obtuvimos %d", triangles)
	}
}
