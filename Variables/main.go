package main // definimos la función princiapal

import ( // importamos los pquetes que vamos a usar
	"fmt" // fmt es un paquete para formatear y mostrar texto en consola
)

func main() {

	var variable_entera int = 12 // Si defenimos el tipo de variable como int el valor puede ser valores de - infinito hasta + infinito
	fmt.Println("El valor de la variable es: ", variable_entera)

	variable_entera2 := 15 // Una forma similar a la sintaxis de python para declarar variables, el tipo de dato lo define el valor de la variable
	fmt.Println("El valor de la variable es: ", variable_entera2)

	var variable_uint uint = 12 // Si definimos el tipo de variable como unit lo valores tienen que ser + infinitivo
	fmt.Println("el valor de la variable es: ", variable_uint)

	var variable_string string = "Illojuan" // tipo de dato string
	fmt.Println("El nombre de la persona es: ", variable_string)

	variable_string2 := "Golang :)" // otra forma de definir un tipo de dato string
	fmt.Println("El nombre de la persona es: ", variable_string2)

	var variable_boleana bool = true                                // tipo de dato boleano
	fmt.Printf("El valor de el boleano es: %t\n", variable_boleana) //formatenado los datos con Printf

	variable_boleana2 := false // otra forma de definir el tipo de dato boleano
	fmt.Printf("El valor de el boleano es: %t\n", variable_boleana2)

	fmt.Println("Ver la direccion de memoria de una variable", &variable_boleana2) // vemos la dirección de memoria utlizando ampersand & y la variable

}
