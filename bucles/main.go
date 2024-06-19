package main

import (
	"fmt"
)

func main() {

	suma := 0

	for i := 0; i < 10; i++ { // Estructura del for

		suma++ // incremento de la variable suma
		fmt.Println(suma)
	}

	suma = 1

	for suma <= 10 {

		fmt.Println(suma)
		suma++

	}

	for { // bucle infinito

		if suma > 20 {

			break // Terminar el bucle infinito si se cumple la condición

		}
		println(suma)
		suma++
	}

	array := []int{1, 2, 3, 4, 5}

	for i := range array { // Estructura para recorrer un array

		fmt.Println("Indice: ", i, "- Valor: ", array[i])
	}

	for i, val := range array { // Podemos almacenar el valor del array diretamente agregando una segunda variable

		fmt.Println("Indice: ", i, "- Valor", val)
	}

	map3 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	for key, value := range map3 { // Estructura para recorrer un map

		fmt.Println("Clave: ", key, "- Valor:", value)
	}

	map4 := map[string][]int{

		"A": {1, 2, 3, 4, 5},
		"B": {6, 7, 8, 9, 10},
		"C": {11, 12, 13, 14, 15},
		"D": nil,
	}

	for key, value := range map4 { // Estructura para recorrer un map que contiene arrays como valor

		fmt.Println(key)
		for _, value := range value {

			fmt.Println("Valor: ", value)
		}
	}

}
