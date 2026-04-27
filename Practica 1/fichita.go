package main

import "fmt"

func main() {

	var nombre string = "Shirley"
	var edad int = 20
	var carrera string = "Tecnologías de la Información"
	var semestre int = 6
	var promedio float64 = 8.50

	fmt.Printf("Soy %s, tengo %d, años\n", nombre, edad)
	fmt.Printf("Estudio %s, semestre %d\n", carrera, semestre)
	fmt.Printf("Mi promedio es %.2f\n", promedio)
}
