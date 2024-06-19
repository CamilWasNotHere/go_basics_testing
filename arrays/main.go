package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//vector

	var my_array = [7]int{1, 2, 3, 4, 5, 6, 7} // Para definir un array tendremso que poner si o si el tamaño del array
	fmt.Printf("Tipo: %T, valor: %d, bytes: %d, bits %d\n", my_array, my_array, unsafe.Sizeof(my_array), unsafe.Sizeof(my_array)*8)

	//slice

	var my_slice []int                                                                                                               // Para definir un slice tendremos que colocar corchetes vacios
	fmt.Printf("Type: %T, valor: %d, bytes: %d, bits: %d\n", my_slice, my_slice, unsafe.Sizeof(my_slice), unsafe.Sizeof(my_slice)*8) // Imprimos el slice sin valores

	my_slice = append(my_slice, 2, 3, 4, 5)                                                                                          // agregamos valores al slice
	fmt.Printf("Type: %T, valor: %d, bytes: %d, bits: %d\n", my_slice, my_slice, unsafe.Sizeof(my_slice), unsafe.Sizeof(my_slice)*8) // imprimimos los valores con el array con los nuevos valroes

	fmt.Println(len(my_slice)) // Para obtener la longitud de un array usamos len

	my_slice2 := make([]int, len(my_slice)) // Para copiar el contenido de un array a otro primero tenemos que darle una longitud para que la capacidad de alamacenar los nuevos valores eso lo logramos con make
	copy(my_slice2, my_slice)               // utilizamos copy para capturar el contenido de un array para copiarlo en otro
	fmt.Println(my_slice2)                  // imprimimos el array

	fmt.Println(my_slice2[1:]) // para obtener poriciones de un array utilizamos ':'

	memoria_compartida := my_array[2:5]
	fmt.Println(my_array)
	fmt.Println(memoria_compartida)
	memoria_compartida[0] = 12
	memoria_compartida[1] = 54
	fmt.Println(memoria_compartida) // Si modificamos los valores de un array o slice que tomo una porción de otro array/slice el array/slice se vera afectando también ya que ambos comparten el mismo sepacio de memoria
	fmt.Println(my_array)           // Imprimos el array con sus nuevos valores

}
