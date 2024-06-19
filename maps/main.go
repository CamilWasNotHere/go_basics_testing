package main

import (
	"fmt"
)

func main() {

	map1 := make(map[int]string) // Declaración de un map
	map1[1] = "A"                // Ingresamos valores al map
	map1[2] = "B"
	map1[3] = "C"

	fmt.Println(map1)    // Imprimos todo el map
	fmt.Println(map1[1]) // Imprimimos un valor en especifico del map a través de su clave

	delete(map1, 2) // Eliminar un elemento de el map
	fmt.Println(map1)

	map1[7] = ""
	println(map1[7])   // Si imprimimos el valor de un key vacia nos saldra por defecto su valor dependiendo del tipo de dato
	println(map1[100]) // Si intentamos imprimir el valor de una key que no existe tambien nos saldra su valor por defecto dependiendo del tipo de dato que hayamos setiado anteriormente

	v1, ok1 := map1[7] // Para saber si el valor existe o no podemos obtener un valor adicional a traves de una segunda variable para saber si existe o no a traves de un boleano
	v2, ok2 := map1[9]

	fmt.Println("El valor es", v1, "Existe:", ok1)
	fmt.Println("El valor es", v2, "Existe:", ok2)

	map2 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(map2)
}
