package main

// getPolygonName devuelve el nombre técnico según el número de lados.
func getPolygonName(n int) string {
	names := map[int]string{
		3:  "Triángulo",
		4:  "Cuadrilátero",
		5:  "Pentágono",
		6:  "Hexágono",
		12: "Dodecágono",
	}
	if val, ok := names[n]; ok {
		return val
	}
	return "Polígono Complejo"
}
